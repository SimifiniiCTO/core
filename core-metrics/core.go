/*
Copyright 2019 The FeelGuuds Authors.

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
package core_metrics

import (
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus"
	apimachineryversion "k8s.io/apimachinery/pkg/version"
	"k8s.io/component-base/version"
)

var (
	MAJOR_VERSION = version.Get().Major
	MINOR_VERSION = version.Get().Minor
	GIT_VERSION   = version.Get().GitVersion
)

var metricVersion = parseVersion(apimachineryversion.Info{
	Major:      MAJOR_VERSION,
	Minor:      MINOR_VERSION,
	GitVersion: GIT_VERSION,
})

// CoreMetricsEngine encapsulates metric registration functionality as well as metric emitting func.
type CoreMetricsEngine struct {
	Registry *platformRegistry
}

// NewCoreMetricsEngineInstance returns an instance to the metrics engine through which metrics can be emitted and registered
func NewCoreMetricsEngineInstance(namespace string, /* more specific to the service which aims to subscribe to the metrics engine */
	db *gorm.DB) *CoreMetricsEngine {
	coreMetricsEngine := &CoreMetricsEngine{Registry: NewPlatformRegistry()}

	if db != nil {
		initializeCoreDatabaseCounters(namespace, db)
		coreMetricsEngine.RegisterCustomMetric(database_metrics...)
	}

	coreMetricsEngine.RegisterMetric(queue_metrics...)
	coreMetricsEngine.RegisterMetric(rest_metrics...)
	return coreMetricsEngine
}

// RegisterMetric registers a metrics to the metrics engine
func (engine *CoreMetricsEngine) RegisterMetric(metrics ...Registerable) {
	engine.Registry.MustRegister(metrics...)
}

// RegisterCustomMetric registers a custom metric to the metrics engine
func (engine *CoreMetricsEngine) RegisterCustomMetric(metrics ...prometheus.Collector) {
	engine.Registry.RawMustRegister(metrics...)
}
