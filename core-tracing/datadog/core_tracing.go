// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package datadog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type OperationType string

const (
	GET_REQUEST    OperationType = "GET"
	PUT_REQUEST    OperationType = "PUT"
	DELETE_REQUEST OperationType = "DELETE"
	POST_REQUEST   OperationType = "POST"
)

type HttpOperationArguments struct {
	UrlEndpoint    string
	DecodedRequest interface{}
	Operation      OperationType
	RequestBody    io.Reader
}

type TracingConfigurations struct {
	// ServiceName specifies the name of the service from which the traces will be tagged
	ServiceName string

	// ServiceVersion specifies the version of the service running
	ServiceVersion string

	// OpentracingAgentAddr specifies the connection url of the opentracing agent
	OpentracingAgentAddr string

	// DataDogAgentAddr specifies the connection url of the datadog agent
	DataDogAgentAddr string

	// TraceWithAnalytics defines wether to provide analytics along with trace data
	TraceWithAnalytics bool

	// TraceWithRunTimeMetrics specifies wether to provide runtime metrics along with traces
	TraceWithRuntimeMetrics bool

	// TraceInDebugMode specifies wether to enable verbose logging while tracing
	TraceInDebugMode bool

	// HttpClient specifies the http client requests used to trace
	HttpClient *http.Client
}

// Start sets up an opentracing instance using Datadog's tracers
//
// Usage:
//  t := NewOpenTracingTracer(configs)
//  defer t.Stop() // useful for data integrity and flushes any leftovers
//
//  // Use it with the Opentracing API. The (already started) Datadog tracer
//	// may be used in parallel with the Opentracing API if desired.
//	opentracing.SetGlobalTracer(t)
func Start(configs *TracingConfigurations) {
	if configs == nil {
		panic("invalid tracing configurations")
	}

	options := []tracer.StartOption{
		tracer.WithService(configs.ServiceName),
		tracer.WithAgentAddr(configs.OpentracingAgentAddr),
		tracer.WithAnalytics(configs.TraceWithAnalytics),
		tracer.WithDogstatsdAddress(configs.DataDogAgentAddr),
		tracer.WithRuntimeMetrics(),
		tracer.WithServiceVersion(configs.ServiceVersion),
		tracer.WithHTTPClient(configs.HttpClient),
	}
	// start datadog tracer
	tracer.Start(options...)
}

// Close stops the started tracer. Subsequent calls are valid but become no-op.
func Close() {
	tracer.Stop()
}

// TraceFunction wraps function with opentracing span adding tags for the function name and caller details
func TraceFunction(ctx context.Context, fn interface{}, params ...interface{}) (result []reflect.Value) {
	// Get function name
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	// Create child span
	parentSpan, _ := tracer.SpanFromContext(ctx)
	sp := tracer.StartSpan(
		fmt.Sprintf("Function %s", name),
		tracer.ChildOf(parentSpan.Context()))
	defer sp.Finish()

	sp.SetTag("function", name)

	// Get caller function name, file and line
	getCallerFunctionMetadata(sp)

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
func CreateChildSpan(ctx context.Context, name string) tracer.Span {
	parentSpan, _ := tracer.SpanFromContext(ctx)
	sp := tracer.StartSpan(
		name,
		tracer.ChildOf(parentSpan.Context()))
	sp.SetTag("name", name)

	// Get caller function name, file and line
	getCallerFunctionMetadata(sp)
	return sp
}

// getCallerFunctionMetadata Gets a calling function metadata including file, function, and execution line
func getCallerFunctionMetadata(sp tracer.Span) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)
}

// TraceHttpRequest properly traces an http request
func TraceHttpRequest(args *HttpOperationArguments, r *http.Request) error {
	if args == nil {
		return errors.New("invalid input arguments")
	}

	span, ctx := tracer.StartSpanFromContext(r.Context(), fmt.Sprintf("%s %s", string(args.Operation), args.UrlEndpoint))
	defer span.Finish()

	req, err := http.NewRequest(string(args.Operation), args.UrlEndpoint, args.RequestBody)
	req = req.WithContext(ctx)
	// Inject the span Context in the Request headers
	err = tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(req.Header))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode >= 400 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(args.DecodedRequest)
}
