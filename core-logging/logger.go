// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package core_logging

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerInstance struct {
	Logger             *zap.Logger
	RedirectLoggerFunc func()
}

// NewProductionLogger creates a production instance of a logger object
func NewProductionLogger() *LoggerInstance {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	stdLog := zap.RedirectStdLog(logger)

	return &LoggerInstance{
		Logger:             logger,
		RedirectLoggerFunc: stdLog,
	}
}

// New creates a logger instance of a logger object
func New(logLevel string) *LoggerInstance {
	var logger *zap.Logger

	logger, _ = initZap(logLevel)
	stdLog := zap.RedirectStdLog(logger)

	return &LoggerInstance{
		Logger:             logger,
		RedirectLoggerFunc: stdLog,
	}
}

// ConfigureLogger configures a logger object
func (l *LoggerInstance) ConfigureLogger() {
	defer l.Logger.Sync()
	defer l.RedirectLoggerFunc()
}

// initZap initializes a zap logging object
func initZap(logLevel string) (*zap.Logger, error) {
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	switch logLevel {
	case "debug":
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	}

	zapEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConfig := zap.Config{
		Level:       level,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zapEncoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return zapConfig.Build()
}
