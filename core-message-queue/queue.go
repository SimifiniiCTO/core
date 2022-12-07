// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package core_message_queue

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type QueueUrlSet struct {
	AlgoliaSearchUrl *string
	SendgridUrl      *string
	GetStreamUrl     *string
}

type ClientParams struct {
	Region                      *string
	Endpoint                    *string
	AwsProfile                  *string
	AwsId                       *string
	AwsSecret                   *string
	Urls                        *QueueUrlSet
	WriteOperationTimeout       *time.Duration
	MaxNumberOfMessagesToIngest *int
	ReadOperationTimeout        *time.Duration
	Attributes                  *[]string
}

type HandleConfig struct {
	MaxNumberOfMessages *int
	MaxWaitTimeSeconds  *time.Duration
	Attributes          *[]string
}

type SqsQueueHandle struct {
	Client    sqsiface.SQSAPI
	QueueUrls *QueueUrlSet
	timeout   time.Duration
	Config    *HandleConfig
}

var _ MessageClientInterface = SqsQueueHandle{}

// NewSQS returns a new sns client for the passed in region
func NewClient(params *ClientParams) (SqsQueueHandle, error) {
	if params == nil {
		return SqsQueueHandle{}, errors.New("invalid input arguments. params cannot be nil")
	}

	if params.Region == nil || params.Endpoint == nil || params.Urls == nil {
		return SqsQueueHandle{}, fmt.Errorf("invalid input arguments. region: %v, endpoint: %v, queueUrl: %v", params.Region, params.Endpoint, params.Urls)
	}

	sess, err := New(Config{
		Address: *params.Endpoint,
		Region:  *params.Region,
		Profile: *params.AwsProfile,
		ID:      *params.AwsId,
		Secret:  *params.AwsSecret,
	})

	if err != nil {
		return SqsQueueHandle{}, err
	}

	return SqsQueueHandle{
		Client:    sqs.New(sess),
		QueueUrls: params.Urls,
		Config: &HandleConfig{
			MaxNumberOfMessages: params.MaxNumberOfMessagesToIngest,
			MaxWaitTimeSeconds:  params.ReadOperationTimeout,
			Attributes:          params.Attributes,
		},
	}, nil
}

func (h SqsQueueHandle) SendMessage(ctx context.Context, msg *sqs.SendMessageInput) (*string, error) {
	//ctx, cancel := context.WithTimeout(ctx, h.timeout)
	//defer cancel()

	res, err := h.Client.SendMessageWithContext(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("could not send message to queue %v: %v", msg.QueueUrl, err)
	}

	return res.MessageId, nil
}

func (h SqsQueueHandle) Receive(ctx context.Context, queueURL string) (*Message, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	res, err := h.Client.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:              aws.String(queueURL),
		MaxNumberOfMessages:   aws.Int64(1),
		WaitTimeSeconds:       aws.Int64(20),
		MessageAttributeNames: aws.StringSlice([]string{"All"}),
	})
	if err != nil {
		return nil, fmt.Errorf("receive: %w", err)
	}

	if len(res.Messages) == 0 {
		return nil, nil
	}

	attrs := make(map[string]string)
	for key, attr := range res.Messages[0].MessageAttributes {
		attrs[key] = *attr.StringValue
	}

	return &Message{
		ID:            *res.Messages[0].MessageId,
		ReceiptHandle: *res.Messages[0].ReceiptHandle,
		Body:          *res.Messages[0].Body,
		Attributes:    attrs,
	}, nil
}

func (h SqsQueueHandle) Delete(ctx context.Context, queueURL, rcvHandle string) error {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	if _, err := h.Client.DeleteMessageWithContext(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(rcvHandle),
	}); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func (h SqsQueueHandle) Send(ctx context.Context, req *SendRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	attrs := make(map[string]*sqs.MessageAttributeValue, len(req.Attributes))
	for _, attr := range req.Attributes {
		attrs[attr.Key] = &sqs.MessageAttributeValue{
			StringValue: aws.String(attr.Value),
			DataType:    aws.String(attr.Type),
		}
	}

	res, err := h.Client.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		MessageAttributes: attrs,
		MessageBody:       aws.String(req.Body),
		QueueUrl:          aws.String(req.QueueURL),
	})
	if err != nil {
		return "", fmt.Errorf("send: %w", err)
	}

	return *res.MessageId, nil
}
