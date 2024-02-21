package main

import (
	"log"
	"log/slog"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const addr = ":8080"

func main() {
	go walk()

	slog.Info("serving metrics", slog.String("addr", addr))
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

// This function walks back and forth between a range of values.
func walk() {
	const min, max = 0, 100
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		slog.Info("queue depth increasing")
		for i := int64(min); i <= max; i++ {
			<-ticker.C
			queueDepth.Store(i)
		}

		slog.Info("queue depth decreasing")
		for i := int64(max); i >= min; i-- {
			<-ticker.C
			queueDepth.Store(i)
		}
	}
}

var queueDepth atomic.Int64

var _ = promauto.NewGaugeFunc(
	prometheus.GaugeOpts{
		Name: "queue_depth",
		Help: "Generated value representing a queue depth.",
	},
	func() float64 { return float64(queueDepth.Load()) },
)
