// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package core_message_queue

type SendRequest struct {
	QueueURL   string
	Body       string
	Attributes []Attribute
}

type Attribute struct {
	Key   string
	Value string
	Type  string
}

type Message struct {
	ID            string
	ReceiptHandle string
	Body          string
	Attributes    map[string]string
}
