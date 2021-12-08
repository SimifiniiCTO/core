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
	"runtime"
)

const (
	RuntimeSubsystem = "Runtime"
	ActiveGoRoutines = "active_go_routines"
)

var (
	activeGoRoutines GaugeFunc
)

func initializeCoreRuntimeCounters(namespace string) {
	activeGoRoutines = NewGaugeFunc(
		GaugeOpts{
			Namespace:      namespace,
			Subsystem:      RuntimeSubsystem,
			Name:           ActiveGoRoutines,
			Help:           "number of goroutines that currently exist",
			StabilityLevel: ALPHA,
		},
		func() float64 { return float64(runtime.NumGoroutine()) },
	)
}
