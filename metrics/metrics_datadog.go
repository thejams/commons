package metrics

import "github.com/DataDog/datadog-go/statsd"

// NewDDClient returns a new datadog client
func NewDDClient(host string) (Metrics, error) {
	return statsd.New(host)
}
