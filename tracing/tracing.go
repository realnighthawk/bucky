package tracing

var (
	OTel TracingType = "otel"
)

type TracingOptions struct {
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type Options struct {
	OTel TracingOptions `json:"otel,omitempty" yaml:"otel,omitempty"`
}

type TracingType string
