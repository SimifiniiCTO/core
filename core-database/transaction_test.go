// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package core_database_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	core_database "github.com/SimifiniiCTO/core/core-database"
)

var (
	dbConn    *core_database.DatabaseConn
	host      = "localhost"
	port      = 5433
	user      = "postgres"
	password  = "postgres"
	dbname    = "postgres"
	globalCtx = context.TODO()
)

func init() {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port)

	queryTimeout := 10 * time.Second
	maxConnectionRetries := 5
	maxConnectionRetryTimeout := 10 * time.Second
	retrySleep := 1 * time.Second

	dbConn = core_database.NewDatabaseConn(&core_database.Parameters{
		QueryTimeout:              &queryTimeout,
		MaxConnectionRetries:      &maxConnectionRetries,
		MaxConnectionRetryTimeout: &maxConnectionRetryTimeout,
		ConnectionString:          &connectionString,
		RetrySleep:                &retrySleep,
	})
}

// TestTransaction Tests the result of a transaction
func TestTransaction(t *testing.T) {
	// success scenarios
	t.Run("TestName:Passed_TransactionPassed", TransactionPassed)
	// failure scenarios
	t.Run("TestName:Failed_Transaction", TransactionFailed)
}

func TestComplexTransaction(t *testing.T) {
	// success scenarios
	t.Run("TestName:Passed_ComplexTransaction", ComplexTransactionPassed)
	// failure scenarios
	t.Run("TestName:Failed_ComplexTransaction", ComplexTransactionFailed)
}

func TransactionFailed(t *testing.T) {
	f := func(ctx context.Context, tx *gorm.DB) error {
		return errors.New("failed transactions")
	}

	err := dbConn.PerformTransaction(globalCtx, f)
	assert.NotEmpty(t, err)
}

func ComplexTransactionFailed(t *testing.T) {
	f := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		return nil, errors.New("failed transactions")
	}

	output, err := dbConn.PerformComplexTransaction(globalCtx, f)
	assert.NotEmpty(t, err)
	assert.Empty(t, output)
}

func TransactionPassed(t *testing.T) {
	f := func(ctx context.Context, tx *gorm.DB) error {
		return nil
	}

	err := dbConn.PerformTransaction(globalCtx, f)
	assert.Empty(t, err)
}

func ComplexTransactionPassed(t *testing.T) {
	f := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		return true, nil
	}

	output, err := dbConn.PerformComplexTransaction(globalCtx, f)
	assert.Empty(t, err)
	assert.NotEmpty(t, output)
}
