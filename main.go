package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	qps = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "http_qps",
		Help: "The number of HTTP requests on / served in the last second",
	})
)

func main() {
	resetQPS()

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func resetQPS() {
	go func() {
		for {
			qps.Set(0)
			time.Sleep(1 * time.Second)
		}
	}()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go!")
	qps.Inc()
}
