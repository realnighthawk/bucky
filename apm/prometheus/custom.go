package prometheus

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func InitMetrics(config Config) {
	if len(config.Counters) > 0 {
		for _, opt := range config.Counters {
			go startCounters(opt, 2*time.Second)
		}
	}

	if len(config.Gauges) > 0 {
		for _, opt := range config.Gauges {
			prometheus.MustRegister(prometheus.NewGauge(opt))
		}
	}
}

func startCounters(opts prometheus.CounterOpts, delay time.Duration) {
	for {
		promauto.NewCounter(opts).Inc()
		time.Sleep(delay)
	}
}
