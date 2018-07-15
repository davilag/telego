package metrics

import (
	"fmt"
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
)

func SetupMetrics() {
	prometheus.MustRegister(messageSent)
	prometheus.MustRegister(messageReceive)
	prometheus.MustRegister(session)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MessageReceived() {
	messageReceive.Inc()
}

func MessageSent() {
	messageSent.Inc()
}

func SessionStarted() {
	fmt.Println("Starting session")
	session.Inc()
}

func SessionFinished() {
	fmt.Println("Finishing session")
	session.Dec()
}
