package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "wuh_http_request_total",
		Help: "Total number of HTTP requests processed by the server",
	}, []string{
		"method",
		"path",
		"status",
	})
	HttpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "wuh_http_request_duration_seconds_total",
		Help:    "Total duration of HTTP requests processed by the server",
		Buckets: prometheus.DefBuckets,
	}, []string{
		"method",
		"path",
		"status",
	})
)

func init() {
	prometheus.MustRegister(HttpRequestTotal, HttpRequestDuration)
}
