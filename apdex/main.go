package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	apdexT := 0.1
	apdexTarget := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_request_apdex_target_seconds",
			Help: "A histogram of latencies for requests.",
		})
	apdexTarget.Set(apdexT)

	duration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "A histogram of latencies for requests.",
			Buckets: []float64{apdexT / 2, apdexT, apdexT * 2, 4 * apdexT, 8 * apdexT},
		},
		[]string{"code", "handler"},
	)

	prometheus.MustRegister(apdexTarget, duration)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rand.Float64() > 0.8 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("boom!"))
			return
		}

		time.Sleep(time.Duration(float64(time.Second) * (0.1 + 0.1*rand.NormFloat64())))
		w.Write([]byte("OK"))
	})

	healthz := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	handlerChain := promhttp.InstrumentHandlerDuration(
		duration.MustCurryWith(prometheus.Labels{"handler": "/q"}),
		handler)

	healthzChain := promhttp.InstrumentHandlerDuration(
		duration.MustCurryWith(prometheus.Labels{"handler": "/healthz"}),
		healthz)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/healthz", healthzChain)
	http.Handle("/q", handlerChain)
	http.ListenAndServe(":8080", nil)
}
