package jaeger

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
)

// HTTPClient wraps an http.Client with tracing instrumentation.
type HTTPClient struct {
	TracerEngine *TracingEngine
	Client       *http.Client
}

func NewHTTPClient(tracer *TracingEngine) *HTTPClient {
	return &HTTPClient{
		TracerEngine: tracer,
		Client:       http.DefaultClient,
	}
}

// GetJSON executes HTTP GET against specified url and tried to parse
// the response into out object.
func (c *HTTPClient) GetJSON(ctx context.Context, endpoint string, url string, out interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req, ht := nethttp.TraceRequest(c.TracerEngine.Tracer, req, nethttp.OperationName("HTTP GET: "+endpoint))
	defer ht.Finish()

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}
