package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	messageSent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "telego_message_sent",
		Help: "Telego sending a message",
	})
	messageReceive = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "telego_message_receive",
		Help: "Telego receiving a message",
	})

	session = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "telego_session_counter",
		Help: "Sessions that telego is keeping waiting for messages",
	})

	abstract = map[string]prometheus.Collector{}
)

// ExposeMetrics exposes the metrics so they can be queury by a prometheus server
func ExposeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// AddGauge adds and registers a Gauge metric to the prometheus client
func AddGauge(name, help string) {
	abstract[name] = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: name,
			Help: help,
		},
	)
	prometheus.MustRegister(abstract[name])
}

// AddCounter adds and registers a Counter metric to the prometheus client
func AddCounter(name, help string) {
	abstract[name] = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
	)
	prometheus.MustRegister(abstract[name])
}

// AddHistogram adds and registers an Histogram metric to the prometheus client
func AddHistogram(name, help string) {
	abstract[name] = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: name,
			Help: help,
		},
	)
	prometheus.MustRegister(abstract[name])
}

// AddSummary adds and registers an Summary metric to the prometheus client
func AddSummary(name, help string) {
	abstract[name] = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: name,
			Help: help,
		},
	)
	prometheus.MustRegister(abstract[name])
}

// Unregister removes the metric from the exposed metrics
func Unregister(name string) {
	v, ok := abstract[name]
	if ok {
		prometheus.Unregister(v)
	}
}

// GetGauge gets a gauge metric given the name. It returns two values,
// the metric and a boolean indicating if the metric exists
func GetGauge(name string) (prometheus.Gauge, bool) {
	v, ok := abstract[name]
	return v.(prometheus.Gauge), ok
}

// GetCounter gets a counter metric given the name. It returns two values,
// the metric and a boolean indicating if the metric exists
func GetCounter(name string) (prometheus.Counter, bool) {
	v, ok := abstract[name]
	return v.(prometheus.Counter), ok
}

// GetHistogram gets a historic metric given the name. It returns two values,
// the metric and a boolean indicating if the metric exists
func GetHistogram(name string) (prometheus.Histogram, bool) {
	v, ok := abstract[name]
	return v.(prometheus.Histogram), ok
}

// GetSummary gets a summary metric given the name. It returns two values,
// the metric and a boolean indicating if the metric exists
func GetSummary(name string) (prometheus.Summary, bool) {
	v, ok := abstract[name]
	return v.(prometheus.Summary), ok
}
