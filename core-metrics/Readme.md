## Blackspace Core Metrics Library
---
```bash
go get github.com/Lens-Platform/Platform/src/libraries/core/core-metrics
```

Requires:

* Go >= 1.12

This document outlines how to effectively make use of this library.

```go
// define a core engine registry object to which the version info would be tied to
CoreEngine := NewCoreMetricsEngineInstance("blackspace_platform",nil)

// Define a
LoginRequestCounter = NewGaugeVec(&GaugeOpts{
        Namespace: "blackspace_platform"
		Subsystem:  "authenticationservice",
		Name:      "login_request_counter",
		Help:      "Number of log in requests",
	}, []string{"Request"})

CoreEngine.RegisterMetric(LoginRequestCounter)

LoginRequestCounter.WithLabelValues("Request").Observe(1.0)
```
