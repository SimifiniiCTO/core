package core_message_consumer

import (
	"context"
	"time"

	core_message_queue "github.com/SimifiniiCTO/core/core-message-queue"
)

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
		message []*core_message_queue.Message
		err     error
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
