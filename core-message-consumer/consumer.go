package core_message_consumer

import (
	"context"
	"fmt"
	"time"

	core_message_queue "github.com/SimifiniiCTO/core/core-message-queue"
	telemetry "github.com/SimifiniiCTO/core/core-telemetry"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

// MessageProcessorFunc serves as the logic used to process each incoming message from a msg queue
type MessageProcessorFunc = func(ctx context.Context, message *core_message_queue.Message) error

type Consumer struct {
	Logger                *zap.Logger
	NewRelicClient        *newrelic.Application
	MetricEngine          *telemetry.MetricsEngine
	ServiceMetrics        *telemetry.ServiceMetrics
	Client                *core_message_queue.SqsQueueHandle
	QueueUrl              *string
	ConcurrencyFactor     int
	QueuePollingDuration  time.Duration
	MessageProcessTimeout time.Duration
}

type IConsumer interface {
	ConcurrentConsumer(f MessageProcessorFunc)
	NaiveConsumer(f MessageProcessorFunc)
}

var _ IConsumer = (*Consumer)(nil)

type ConsumerConfigs struct {
	ConcurrencyFactor     int
	MessageProcessTimeout time.Duration
	QueuePollingDuration  time.Duration
}

type ConsumerParams struct {
	QueueURl       *string
	Logger         *zap.Logger
	NrClient       *newrelic.Application
	MetricsEngine  *telemetry.MetricsEngine
	ServiceMetrics *telemetry.ServiceMetrics
	AwsClient      *core_message_queue.SqsQueueHandle
	Config         *ConsumerConfigs
}

// NewConsumer instantiates a new instance of the aws consumer object
func NewConsumer(params *ConsumerParams) (*Consumer, error) {
	if params == nil {
		return nil, fmt.Errorf("invalid input argument. params: %v", params)
	}

	if err := validateConsumerParams(params); err != nil {
		return nil, err
	}

	return &Consumer{
		Logger:                params.Logger,
		NewRelicClient:        params.NrClient,
		MetricEngine:          params.MetricsEngine,
		ServiceMetrics:        params.ServiceMetrics,
		Client:                params.AwsClient,
		QueueUrl:              params.QueueURl,
		ConcurrencyFactor:     params.Config.ConcurrencyFactor,
		MessageProcessTimeout: params.Config.MessageProcessTimeout,
		QueuePollingDuration:  params.Config.QueuePollingDuration,
	}, nil
}

// reportErr emits an error metric and logs erros obtained while processing message asynchronously
func (c *Consumer) reportErr(op string, err error) {
	c.Logger.Error(err.Error())
	c.MetricEngine.RecordErrorCountMetric(op)
}

func (c *Consumer) reportProcessedMessage(op string) {
	c.MetricEngine.RecordStandardMetrics(op, true)
}

func validateConsumerParams(params *ConsumerParams) error {
	if params.AwsClient == nil ||
		params.Config == nil ||
		params.Logger == nil ||
		params.MetricsEngine == nil ||
		params.NrClient == nil ||
		params.QueueURl == nil ||
		params.ServiceMetrics == nil {
		return fmt.Errorf("invalid params object. params: %v", params)
	}

	return nil
}
