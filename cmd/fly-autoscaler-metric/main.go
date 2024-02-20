package main

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func main() {
	go walk()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// This function walks back and forth between a range of values.
func walk() {
	const min, max = 0, 100
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		for i := int64(min); i <= max; i++ {
			<-ticker.C
			queueDepth.Store(i)
		}

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
