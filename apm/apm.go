package apm

var (
	Prometheus MetricsType = "prometheus"
)

type MetricOptions struct {
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type Options struct {
	Prometheus MetricOptions `json:"prometheus,omitempty" yaml:"prometheus,omitempty"`
}

type MetricsType string
