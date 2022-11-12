package metrics

import (
	"errors"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

type MetricType string

var (
	OperationLatencyMetric MetricType = "service.operation.latency"
	OperationStatusMetric  MetricType = "service.operation.status"
	RpcStatusMetric        MetricType = "service.rpc.status"
)

// MetricsEngine enables this service to emit metrics to new relic
type MetricsEngine struct {
	// Metrics encompasses all metrics defined for the various operations this service is part of
	Metrics *ServiceMetrics
	// ServiceName encompasses the name of the service
	ServiceName *string
	// Core is the utility by which metrics are emitted to new-relic
	Client  *newrelic.Application
	Core    *MetricsCore
	enabled bool
	logger  *zap.Logger
}

type ServiceDetails struct {
	ServiceName string
	// The version of the service actively deployed
	Version string
	// The service P.O.
	PointOfContact string
	// A link to documentation around the service's functionality and uses
	DocumentationLink string
	// The environment in which the service is actively running and deployed in
	Environment string
	// license key for interactions with the new-relic platform
	NewRelicLicenseKey string
}

// NewTelemetry returns a instance of the metrics engine object in which all defined service metrics are present
func newMetricsEngine(params *TelemetryParams) (*MetricsEngine, error) {
	if params == nil {
		return nil, errors.New("invalid input argument. params must be provided")
	}

	metadata := &ServiceMetadata{
		Name:              params.ServiceName,
		Version:           params.Version,
		PointOfContact:    params.PointOfContact,
		DocumentationLink: params.DocumentationLink,
		Environment:       params.Environment,
	}

	metrics, err := newServiceMetrics(&params.ServiceName)
	if err != nil {
		return nil, err
	}

	if !params.Enabled {
		return &MetricsEngine{
			Core:        nil,
			ServiceName: &params.ServiceName,
			Metrics:     metrics,
			enabled:     params.Enabled,
			logger:      params.Logger,
			Client:      params.Client,
		}, err
	}

	engine, err := newMetricsCore(&params.NewRelicLicenseKey, metadata)
	if err != nil {
		return nil, err
	}

	return &MetricsEngine{
		Core:        engine,
		ServiceName: &params.ServiceName,
		Metrics:     metrics,
		enabled:     params.Enabled,
		logger:      params.Logger,
	}, nil
}
