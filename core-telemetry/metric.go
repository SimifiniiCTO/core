// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package metrics

// Metric encompasses the relative metadata associated with a service level metric being emitted
type Metric struct {
	MetricName       string
	Name             string
	Help             string
	Subsystem        Subsystem
	Namespace        Namespace
	ServiceName      string
	MetricPartitions map[string]string
}

// Counter a metric type that only linearly increases
type Counter struct {
	Metric
}

// Summary summarizes numerous facets of a metric
type Summary struct {
	Metric
}

// Gauge a metric type that can increase and decrease
type Gauge struct {
	Metric
}

type Namespace string
type Subsystem string

const (
	RequestNamespace                   Namespace = "request.namespace"
	DistributedTxNamespace             Namespace = "dtx.namespace"
	DistributedSagaTxNamespace         Namespace = "dtx.saga.namespace"
	DistributedTx2PhaseCommitNamespace Namespace = "dtx.2phase_commit.namespace"
	DistributedTxTccNamespace          Namespace = "dtx.tcc.namespace"
	ServiceNamespace                   Namespace = "service.namespace"
	DatabaseNamespace                  Namespace = "database.namespace"
)

const (
	GrpcSubSystem       Subsystem = "grpc.subsystem"
	ThirdPartySubSystem Subsystem = "thirdparty.subsystem"
	ErrorSubSystem      Subsystem = "error.subsystem"
	DbSubSystem         Subsystem = "database.subsystem"
)
