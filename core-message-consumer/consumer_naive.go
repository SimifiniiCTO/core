package core_message_consumer

import (
	"context"
	"sync"
	"time"
)

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
