// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package metrics

import (
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/rpc/code"
)

func (me *MetricsEngine) recordDtxMetric(m *Metric, op, dest string, statusCode code.Code, start time.Duration) {
	if !me.enabled {
		return
	}

	if !me.validateMetricEngine() {
		me.logger.Panic("unable to emit metrics due to misconfiguration")
	}

	mHandle := me.Core.Havester.MetricAggregator()
	mHandle.Summary(m.MetricName, map[string]interface{}{
		"service.source":      m.ServiceName,
		"service.operation":   op,
		"service.destination": dest,
		"metric.help":         m.Help,
		"namespace":           m.Namespace,
		"subsystem":           m.Subsystem,
		"duration":            start,
		"status.code":         statusCode,
	})
}

func (me *MetricsEngine) recordLatencyMetric(m *Metric, op, dest string, start time.Duration) {
	if !me.enabled {
		return
	}

	if !me.validateMetricEngine() {
		me.logger.Panic("unable to emit metrics due to misconfiguration")
	}

	me.logger.Info(fmt.Sprintf("MetricName: %s. Emitting Metric ========>", m.MetricName))

	mHandle := me.Core.Havester.MetricAggregator()
	mHandle.Summary(m.MetricName, map[string]interface{}{
		"service.source":      m.ServiceName,
		"service.operation":   op,
		"service.destination": dest,
		"metric.help":         m.Help,
		"namespace":           m.Namespace,
		"subsystem":           m.Subsystem,
		"duration":            start,
	})
}

func (me *MetricsEngine) recordCounterMetric(m *Metric, op string, statusCode code.Code) {
	if !me.enabled {
		return
	}

	if !me.validateMetricEngine() {
		me.logger.Panic("unable to emit metrics due to misconfiguration")
	}

	mHandle := me.Core.Havester.MetricAggregator()
	mHandle.Count(m.MetricName, map[string]interface{}{
		"service.source":    m.ServiceName,
		"service.operation": op,
		"metric.help":       m.Help,
		"namespace":         m.Namespace,
		"subsystem":         m.Subsystem,
		"status.code":       statusCode,
	})
}

func (me *MetricsEngine) recordOpCounterMetric(m *Metric, op string) {
	if !me.enabled {
		return
	}

	if !me.validateMetricEngine() {
		me.logger.Panic("unable to emit metrics due to misconfiguration")
	}

	mHandle := me.Core.Havester.MetricAggregator()
	mHandle.Count(m.MetricName, map[string]interface{}{
		"service.source":    m.ServiceName,
		"service.operation": op,
		"metric.help":       m.Help,
		"metric.namespace":  m.Namespace,
		"metric.subsystem":  m.Subsystem,
	})
}

func (me *MetricsEngine) recordErrorMetric(m *Metric, op, msg string, timeOfOccurence time.Time) {
	if !me.enabled {
		return
	}

	if !me.validateMetricEngine() {
		me.logger.Panic("unable to emit metrics due to misconfiguration")
	}

	mHandle := me.Core.Havester.MetricAggregator()
	mHandle.Summary(m.MetricName, map[string]interface{}{
		"service.source":    m.ServiceName,
		"service.operation": op,
		"error.message":     msg,
		"metric.help":       m.Help,
		"metric.namespace":  m.Namespace,
		"metric.subsystem":  m.Subsystem,
		"metric.occurence":  timeOfOccurence.String(),
	})
}

func (me *MetricsEngine) recordSummaryMetric(m *Metric, op, dest string) {
	if !me.enabled {
		return
	}

	if !me.validateMetricEngine() {
		me.logger.Panic("unable to emit metrics due to misconfiguration")
	}

	mHandle := me.Core.Havester.MetricAggregator()
	mHandle.Summary(m.MetricName, map[string]interface{}{
		"service.source":      m.ServiceName,
		"service.operation":   op,
		"service.destination": dest,
		"metric.help":         m.Help,
		"namespace":           m.Namespace,
		"subsystem":           m.Subsystem,
	})
}

func (me *MetricsEngine) recordGaugeMetric(m *Metric, op string) {
	if !me.enabled {
		return
	}

	if !me.validateMetricEngine() {
		me.logger.Panic("unable to emit metrics due to misconfiguration")
	}

	mHandle := me.Core.Havester.MetricAggregator()
	mHandle.Gauge(m.MetricName, map[string]interface{}{
		"service.source":    m.ServiceName,
		"service.operation": op,
		"service.help":      m.Help,
		"namespace":         m.Namespace,
		"subsystem":         m.Subsystem,
	})
}

// This is a validation function that checks if the metrics engine is enabled and if the core is not
// nil.
func (me *MetricsEngine) validateMetricEngine() bool {
	if me == nil || (me.Core == nil && !me.enabled) || me.Core.Havester == nil {
		return false
	}

	return true
}
