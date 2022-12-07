// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package core_utilities

import (
	"math"
	"math/rand"
	"time"
)

// Sleep generates a normally distributed random delay with given mean and stdDev
// and blocks for that duration.
func Sleep(mean time.Duration, stdDev time.Duration) {
	fMean := float64(mean)
	fStdDev := float64(stdDev)
	delay := time.Duration(math.Max(1, rand.NormFloat64()*fStdDev+fMean))
	time.Sleep(delay)
}
