// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package core_database

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/giantswarm/retry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConn struct {
	Engine                    *gorm.DB
	QueryTimeout              *time.Duration
	MaxConnectionRetries      *int
	MaxConnectionRetryTimeout *time.Duration
	RetrySleep                *time.Duration
	ConnectionString          *string
}

type Parameters struct {
	QueryTimeout              *time.Duration
	MaxConnectionRetries      *int
	MaxConnectionRetryTimeout *time.Duration
	RetrySleep                *time.Duration
	ConnectionString          *string
}

// NewDatabaseConn obtains a reference to a database connection object
func NewDatabaseConn(params *Parameters) *DatabaseConn {
	if err := validateParams(params); err != nil {
		os.Exit(1)
	}

	conn, err := Connect(params)
	if err != nil {
		panic("failed to connect to database")
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := conn.DB()
	if err != nil {
		panic("failed to obtain generic db connection")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// TODO: Make this a configurable value
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	// TODO: Make this a configurable value
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &DatabaseConn{
		Engine:                    conn,
		QueryTimeout:              params.QueryTimeout,
		MaxConnectionRetries:      params.MaxConnectionRetries,
		MaxConnectionRetryTimeout: params.MaxConnectionRetryTimeout,
		RetrySleep:                params.RetrySleep,
		ConnectionString:          params.ConnectionString,
	}
}

// Connect attempts to connect to the database using retries
func Connect(params *Parameters) (*gorm.DB, error) {
	var (
		connectionString     = *params.ConnectionString
		maxConnectionRetries = *params.MaxConnectionRetries
		maxRetryTimeout      = *params.MaxConnectionRetryTimeout
		retrySleepInterval   = *params.RetrySleep
	)
	var connection = make(chan *gorm.DB, 1)

	err := retry.Do(
		func(conn chan<- *gorm.DB) func() error {
			return func() error {
				newConn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
				conn <- newConn
				return err
			}
		}(connection),
		retry.MaxTries(maxConnectionRetries),
		retry.Timeout(maxRetryTimeout),
		retry.Sleep(retrySleepInterval),
	)

	if err != nil {
		return nil, errors.New("exceeded retries")
	}

	return <-connection, nil
}

func validateParams(params *Parameters) error {
	if params == nil {
		return fmt.Errorf("invalid input argument, param cannot be nil")
	}

	if params.ConnectionString == nil || params.MaxConnectionRetries == nil || params.MaxConnectionRetryTimeout == nil || params.QueryTimeout == nil {
		return fmt.Errorf("invalid input argument")
	}

	return nil
}
