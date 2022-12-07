// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package core_message_queue

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type MessageClientInterface interface {
	// Sends a message to a queue
	Send(ctx context.Context, req *SendRequest) (string, error)
	// Sends a message to a queue
	SendMessage(ctx context.Context, msg *sqs.SendMessageInput) (*string, error)
	// Long polls given amount of messages from a queue.
	Receive(ctx context.Context, queueURL string) (*Message, error)
	// Deletes a message from a queue.
	Delete(ctx context.Context, queueURL, rcvHandle string) error
}
