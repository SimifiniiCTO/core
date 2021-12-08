/*
Copyright 2019 The FeelGuuds Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package core_metrics

import (
	"math"
	"net/url"
	"time"
)

const (
	RequestLatency     = "request_duration_seconds"
	RateLimiterLatency = "ratelimiter_duration_seconds"
	RequestResult      = "request_total"
	CertTTL            = "exec_cert_ttl_seconds"
	CertRotation       = "exec_cert_rotation_age"
)

var (
	// requestLatency is a Prometheus Summary metric type partitioned by
	// "verb" and "url" labels. It is used for the rest client latency metrics.
	requestLatency = NewHistogramVec(
		&HistogramOpts{
			Name:    RequestLatency,
			Help:    "Request latency in seconds. Broken down by verb and URL.",
			Buckets: ExponentialBuckets(0.001, 2, 10),
		},
		[]string{"verb", "url"},
	)

	rateLimiterLatency = NewHistogramVec(
		&HistogramOpts{
			Name:    RateLimiterLatency,
			Help:    "Client side rate limiter latency in seconds. Broken down by verb and URL.",
			Buckets: ExponentialBuckets(0.001, 2, 10),
		},
		[]string{"verb", "url"},
	)

	requestResult = NewCounterVec(
		&CounterOpts{
			Name: RequestResult,
			Help: "Number of HTTP requests, partitioned by status code, method, and host.",
		},
		[]string{"code", "method", "host"},
	)

	execPluginCertTTLAdapter = &expiryToTTLAdapter{}

	execPluginCertTTL = NewGaugeFunc(
		GaugeOpts{
			Name: CertTTL,
			Help: "Gauge of the shortest TTL (time-to-live) of the client " +
				"certificate(s) managed by the auth exec plugin. The value " +
				"is in seconds until certificate expiry (negative if " +
				"already expired). If auth exec plugins are unused or manage no " +
				"TLS certificates, the value will be +INF.",
			StabilityLevel: ALPHA,
		},
		func() float64 {
			if execPluginCertTTLAdapter.e == nil {
				return math.Inf(1)
			}
			return execPluginCertTTLAdapter.e.Sub(time.Now()).Seconds()
		},
	)

	execPluginCertRotation = NewHistogram(
		&HistogramOpts{
			Name: CertRotation,
			Help: "Histogram of the number of seconds the last auth exec " +
				"plugin client certificate lived before being rotated. " +
				"If auth exec plugin client certificates are unused, " +
				"histogram will contain no data.",
			// There are three sets of ranges these buckets intend to capture:
			//   - 10-60 minutes: captures a rotation cadence which is
			//     happening too quickly.
			//   - 4 hours - 1 month: captures an ideal rotation cadence.
			//   - 3 months - 4 years: captures a rotation cadence which is
			//     is probably too slow or much too slow.
			Buckets: []float64{
				600,       // 10 minutes
				1800,      // 30 minutes
				3600,      // 1  hour
				14400,     // 4  hours
				86400,     // 1  day
				604800,    // 1  week
				2592000,   // 1  month
				7776000,   // 3  months
				15552000,  // 6  months
				31104000,  // 1  year
				124416000, // 4  years
			},
		},
	)

	rest_metrics = []Registerable{
		requestLatency, requestResult, rateLimiterLatency,
	}
)

type latencyAdapter struct {
	m *HistogramVec
}

func (l *latencyAdapter) Observe(verb string, u url.URL, latency time.Duration) {
	l.m.WithLabelValues(verb, u.String()).Observe(latency.Seconds())
}

type resultAdapter struct {
	m *CounterVec
}

func (r *resultAdapter) Increment(code, method, host string) {
	r.m.WithLabelValues(code, method, host).Inc()
}

type expiryToTTLAdapter struct {
	e *time.Time
}

func (e *expiryToTTLAdapter) Set(expiry *time.Time) {
	e.e = expiry
}

type rotationAdapter struct {
	m *Histogram
}

func (r *rotationAdapter) Observe(d time.Duration) {
	r.m.Observe(d.Seconds())
}
