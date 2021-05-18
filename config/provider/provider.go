// Package provider provides config provider implementations that can be used in the adapters, as well as the Options type containing options for various aspects of an adapter.
package provider

const (
	// ViperKey corresponds to viper configuration
	ViperKey = "viper"
	// InMemKey corresponds to in memory configuration
	InMemKey = "in-memory"
)

// Options contains config options for various aspects of an adapter.
type Options struct {
	FilePath string
	FileType string
	FileName string
}
