// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

/*
Copyright 2019 The Simfiny Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"fmt"
	"os"

	"github.com/newrelic/newrelic-telemetry-sdk-go/telemetry"
)

type ctxKeyType struct{}

var (
	// TransactionContextKey is the key used for newrelic.FromContext and
	// newrelic.NewContext.
	TransactionContextKey = ctxKeyType(struct{}{})
)

// MetricsCore encapsulates the registration functionality as well as the facility to emit metrics to
// the new relic platform for observance
type MetricsCore struct {
	Havester *telemetry.Harvester
}

// ServiceMetadata outlines important pieces of information pertaining to the service
// the data points making up this object should further aid the on-call engineer to
// properly root cause any ambiguities tied to a ny metrics
type ServiceMetadata struct {
	// Name is the service name
	Name string
	// The version of the service actively deployed
	Version string
	// The service P.O.
	PointOfContact string
	// A link to documentation around the service's functionality and uses
	DocumentationLink string
	// The environment in which the service is actively running and deployed in
	Environment string
}

// newMetricsCore returns a new instance of the service metrics engine  that can be used to record metrics
func newMetricsCore(licenseKey *string, serviceMetadata *ServiceMetadata) (*MetricsCore, error) {
	if serviceMetadata == nil {
		return nil, fmt.Errorf("invalid input argument. service name cannot be nil")
	}

	if licenseKey == nil {
		return nil, fmt.Errorf("invalid input argument. license key cannot be nil")
	}

	metricCfg := telemetry.ConfigAPIKey(*licenseKey)
	svcAttr := telemetry.ConfigCommonAttributes(map[string]interface{}{
		"app.name":             serviceMetadata.Name,
		"app.version":          serviceMetadata.Version,
		"app.point-of-contact": serviceMetadata.PointOfContact,
		"app.docs":             serviceMetadata.DocumentationLink,
		"app.env":              serviceMetadata.Environment,
	})
	errLogAttr := telemetry.ConfigBasicErrorLogger(os.Stderr)
	debugLogAttr := telemetry.ConfigBasicDebugLogger(os.Stdout)

	h, err := telemetry.NewHarvester(metricCfg, svcAttr, errLogAttr, debugLogAttr)
	if err != nil {
		return nil, err
	}

	return &MetricsCore{
		Havester: h,
	}, nil
}
