package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
	// Types of prometheus metrics
	Counters   []prometheus.CounterOpts   `json:"counters,omitempty" yaml:"counters,omitempty"`
	Gauges     []prometheus.GaugeOpts     `json:"gauges,omitempty" yaml:"gauges,omitempty"`
	Histograms []prometheus.HistogramOpts `json:"histograms,omitempty" yaml:"histograms,omitempty"`
	Summaries  []prometheus.SummaryOpts   `json:"summaries,omitempty" yaml:"summaries,omitempty"`
}

func GetHTTPHandler() http.Handler {
	return promhttp.Handler()
}
