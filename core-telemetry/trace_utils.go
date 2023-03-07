package metrics

import (
	"context"
	"fmt"
	"strings"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type operationType string

const (
	DtxOperationType             operationType = "dtx"
	RpcOperationType             operationType = "rpc"
	RestOperationType            operationType = "rest"
	SagaOperationType            operationType = "saga"
	TwoPhasedCommitOperationType operationType = "2phasedcommit"
	DbOperationType              operationType = "db"
)

type TxMetadata struct {
	OperationName string
	OperationType operationType
	ServiceName   string
}

// Format the metadata to ensure that the values are in the correct format
func (m TxMetadata) Format() {
	m.OperationName = strings.TrimSpace(strings.ToLower(m.OperationName))
	m.ServiceName = strings.TrimSpace(strings.ToLower(m.ServiceName))
}

// TxName returns a formatted transaction name
func (t Telemetry) TxName(meta *TxMetadata) string {
	meta.Format()
	return fmt.Sprintf("tx.%s.type.%s.service.%s", meta.OperationName, meta.OperationType, meta.ServiceName)
}

// SpanName returns a formatted span name
func (t Telemetry) SpanName(meta *TxMetadata) string {
	meta.Format()
	return fmt.Sprintf("segment.%s.type.%s.service.%s", meta.OperationName, meta.OperationType, meta.ServiceName)
}

// TxFromContext returns a transaction from the context if it exists, otherwise it creates a new transaction
func (t Telemetry) TxFromContext(ctx context.Context, meta *TxMetadata, sdk *newrelic.Application) *newrelic.Transaction {
	if !t.Engine.enabled {
		return &newrelic.Transaction{}
	}

	tx := newrelic.FromContext(ctx)
	if tx == nil {
		return sdk.StartTransaction(t.TxName(meta))
	}

	return tx
}

// NewTransaction returns a new transaction based on the provided metadata
func (t Telemetry) NewTransaction(ctx context.Context, meta *TxMetadata) *newrelic.Transaction {
	if !t.Engine.enabled {
		return &newrelic.Transaction{}
	}

	return t.Engine.Client.StartTransaction(t.TxName(meta))
}

// NewSegment returns a new segment based on the provided metadata
func (t Telemetry) NewSegment(txn *newrelic.Transaction, meta *TxMetadata) *newrelic.Segment {
	if !t.Engine.enabled {
		return &newrelic.Segment{}
	}

	return newrelic.StartSegment(txn, t.SpanName(meta))
}
