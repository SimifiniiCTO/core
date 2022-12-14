package core_message_consumer

import (
	"context"
	"fmt"
	"sync"
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

// ConcurrentConsumer creates a limited parallel queue, and continues to poll AWS until all the limit is reached.
// This is performed by implementing a token bucket” using a buffered channel hence this approach is only limited by aws throughput
//
// Some scenarios will require a different set of resources consumed, depending on the message type (Lets say you want your handler to be able to process from 1 to N emails in 1 message).
//  To maintain our limitations, we could introduce the timely based token bucket algorithm , which will ensure we don’t process more than N emails over a period of time (like 1 minute),
//  by grabbing the exact amount of “worker tokens” from the pool, depending on emails count in message. Also, if your code can be timed out, there is a good approach to impose timeout and cancellation,
//  based on golang context.WithCancel function. Check out the golang semaphore library to build the nuclear-resistant solution. (the mechanics are the same as in our example, abstracted to library,
// so instead of using channel for limiting our operation we will call semaphore.Acquire, which will also block our execution until “worker tokens” will be refilled).
//
//LINK - Ref: https://docs.microsoft.com/en-us/azure/architecture/microservices/model/domain-analysis
//LINK - Ref: https://docs.microsoft.com/en-us/azure/architecture/microservices/design/interservice-communication
func (c *Consumer) ConcurrentConsumer(f MessageProcessorFunc) {
	var (
		messages []*core_message_queue.Message
		err      error
	)
	sync := createFullBufferedChannel(c.ConcurrencyFactor)
	for {
		ctx := context.Background()
		messages, err = c.Client.Receive(ctx, *c.QueueUrl)
		if err != nil {
			c.reportErr("receive_message", err)
			continue
		}

		if len(messages) == 0 {
			time.Sleep(c.QueuePollingDuration)
		} else {
			for _, message := range messages {
				// request the exact amount of "workers" from pool.
				// Again, empty buffer will block this operation
				<-sync

				go func() {
					f(ctx, message)
					// return "worker" to the "pool"
					sync <- true
				}()
			}
		}
	}
}

func createFullBufferedChannel(capacity int) chan bool {
	sync := make(chan bool, capacity)

	for i := 0; i < capacity; i++ {
		sync <- true
	}
	return sync
}

// With an SQS message subscriber we will be receiving messages in small batches out of the box,
// In order for our message consumer to achieve a high throughput, we will process the messages in parallel,
// and in order this to be robust, we should impose a limit on how many messages we should process simultaneously.

// As standard aws sqs receive call gives us maximum of 10 messages, the naive approach will be to process them
//  in parallel, then call the next batch.
//
// With approach like this we will be limited to the
// 1 minute / slowest message processing in batch * 10, for example having the slowest message being processed in 50ms
// it will give us (1000 ms / 50ms) * 10 = 200 messages per second of processing time minus network latency, that can
// eat up most of the projected capacity.
func (c *Consumer) NaiveConsumer(f MessageProcessorFunc) {
	for {
		ctx := context.Background()
		messages, err := c.Client.Receive(ctx, *c.QueueUrl)
		if err != nil {
			c.reportErr("receive_message", err)
			continue
		}

		wg := sync.WaitGroup{}
		for _, message := range messages {
			wg.Add(1)
			go func() {
				defer wg.Done()
				f(ctx, message)
			}()
		}

		if len(messages) == 0 { // add aditional sleep if queue is empty
			time.Sleep(c.QueuePollingDuration)
			continue
		}
	}
}
