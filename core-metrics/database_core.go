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
)

const (
	DatabaseSubsystem          = "Database"
	OpenDbConnections          = "open_database_connections"
	IdleDbConnections          = "idle_database_connections"
	ConnectionsInUse           = "database_connections_in_use"
	ConnectionWaitDuration     = "database_connection_wait_duration"
	ConnectionOperationLatency = "database_connection_operation_latency"
)

var (
	openDatabaseConnections             GaugeFunc
	idleDatabaseConnections             GaugeFunc
	DatabaseConnectionsInUse            GaugeFunc
	DatabaseConnectionsWaitDuration     GaugeFunc
	DatabaseConnectionsOperationLatency *HistogramVec
	database_metrics                    []prometheus.Collector
)

func initializeCoreDatabaseCounters(namespace string, db *gorm.DB) {
	openDatabaseConnections = NewGaugeFunc(
		GaugeOpts{
			Namespace:      namespace,
			Subsystem:      DatabaseSubsystem,
			Name:           OpenDbConnections,
			Help:           "time blocked waiting for a new connection to the database",
			StabilityLevel: ALPHA,
		},
		func(db *gorm.DB) func() float64 {
			return func() float64 {
				return float64(db.DB().Stats().MaxOpenConnections)
			}
		}(db),
	)

	idleDatabaseConnections = NewGaugeFunc(
		GaugeOpts{
			Namespace:      namespace,
			Subsystem:      DatabaseSubsystem,
			Name:           IdleDbConnections,
			Help:           "number of idle database connections opened",
			StabilityLevel: ALPHA,
		},
		func(db *gorm.DB) func() float64 {
			return func() float64 {
				return float64(db.DB().Stats().Idle)
			}
		}(db),
	)

	DatabaseConnectionsInUse = NewGaugeFunc(
		GaugeOpts{
			Namespace:      namespace,
			Subsystem:      DatabaseSubsystem,
			Name:           ConnectionsInUse,
			Help:           "number of database connections in use",
			StabilityLevel: ALPHA,
		},
		func(db *gorm.DB) func() float64 {
			return func() float64 {
				return float64(db.DB().Stats().InUse)
			}
		}(db),
	)

	DatabaseConnectionsWaitDuration = NewGaugeFunc(
		GaugeOpts{
			Namespace:      namespace,
			Subsystem:      DatabaseSubsystem,
			Name:           ConnectionWaitDuration,
			Help:           "time blocked waiting for a new connection to the database",
			StabilityLevel: ALPHA,
		},
		func(db *gorm.DB) func() float64 {
			return func() float64 {
				return float64(db.DB().Stats().WaitDuration)
			}
		}(db),
	)

	DatabaseConnectionsOperationLatency = NewHistogramVec(
		&HistogramOpts{
			Namespace: namespace,
			Subsystem: DatabaseSubsystem,
			Name:      ConnectionOperationLatency,
			Help:      "Request latency in seconds. Broken down by verb and URL.",
			Buckets:   ExponentialBuckets(0.001, 2, 10),
		},
		[]string{"operation"},
	)

	database_metrics = []prometheus.Collector{
		DatabaseConnectionsInUse, DatabaseConnectionsOperationLatency, DatabaseConnectionsWaitDuration, idleDatabaseConnections,
		openDatabaseConnections,
	}
}
