package metrics

import (
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

type Telemetry struct {
	Engine  *MetricsEngine
	Metrics *ServiceMetrics
}

type TelemetryParams struct {
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
	// logger instance used by the newrelic client
	Logger *zap.Logger
	// NewRelic telemetry client
	Client *newrelic.Application
	// Enable metric reporting
	Enabled bool
}

func New(params *TelemetryParams) (*Telemetry, error) {
	if params == nil {
		return nil, fmt.Errorf("invalid input argument. params cannot be nil")
	}

	engine, err := newMetricsEngine(params)
	if err != nil {
		return nil, err
	}

	return &Telemetry{
		Metrics: engine.Metrics,
		Engine:  engine,
	}, nil

}
