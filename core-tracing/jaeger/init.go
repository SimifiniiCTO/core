// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jaeger

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

type TracingEngine struct {
	Tracer opentracing.Tracer
}

// New returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func New(service string, collectorEndpoint string) (*TracingEngine, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			CollectorEndpoint:   collectorEndpoint,
		},
	}

	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger))

	if err != nil {
		panic("Could not initialize jaeger tracer: " + err.Error())
	}

	return &TracingEngine{Tracer: tracer}, closer
}

// TraceFunction wraps funtion with opentracing span adding tags for the function name and caller details
func (engine *TracingEngine) TraceFunction(ctx context.Context, fn interface{}, params ...interface{}) (result []reflect.Value) {
	// Get function name
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	// Create child span
	parentSpan := opentracing.SpanFromContext(ctx)
	sp := opentracing.StartSpan(
		"Function - "+name,
		opentracing.ChildOf(parentSpan.Context()))
	defer sp.Finish()

	sp.SetTag("function", name)

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)

	// Check params and call function
	f := reflect.ValueOf(fn)
	if f.Type().NumIn() != len(params) {
		e := fmt.Sprintf("Incorrect number of parameters calling wrapped function %s", name)
		panic(e)
	}
	inputs := make([]reflect.Value, len(params))
	for k, in := range params {
		inputs[k] = reflect.ValueOf(in)
	}
	return f.Call(inputs)
}

// CreateChildSpan creates a new opentracing span adding tags for the span name and caller details. Returns a Span.
// User must call `defer sp.Finish()`
func (engine *TracingEngine) CreateChildSpan(ctx context.Context, name string) opentracing.Span {
	parentSpan := opentracing.SpanFromContext(ctx)
	sp := opentracing.StartSpan(
		name,
		opentracing.ChildOf(parentSpan.Context()))
	sp.SetTag("name", name)

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)

	return sp
}

// NewTracedRequest generates a new traced HTTP request with opentracing headers injected into it
func (engine *TracingEngine) NewTracedRequest(method string, url string, span opentracing.Span) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, method)
	span.Tracer().Inject(span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	return req, err
}

func (engine *TracingEngine) Inject(span opentracing.Span, request *http.Request) error {
	return span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header))
}

func (engine *TracingEngine) Extract(tracer opentracing.Tracer, r *http.Request) (opentracing.SpanContext, error) {
	return tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
}

func (engine *TracingEngine) OpenTracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wireCtx, _ := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))

		serverSpan := opentracing.StartSpan(r.URL.Path,
			ext.RPCServerOption(wireCtx))
		defer serverSpan.Finish()

		r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
		next.ServeHTTP(w, r)
	})
}

type jaegerLoggerAdapter struct {
	logger zap.SugaredLogger
}

func (l jaegerLoggerAdapter) Error(msg string) {
	l.logger.Error(errors.New(msg), msg)
}

func (l jaegerLoggerAdapter) Infof(msg string, args ...interface{}) {
	l.logger.Info(msg, zap.Any("details", args))
}
