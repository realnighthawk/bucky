// Package provider provides config provider implementations that can be used in the adapters, as well as the Options type containing options for various aspects of an adapter.
package provider

const (
	// Provider keys
	ViperKey = "viper"
	InMemKey = "in-memory"
)

// Type Options contains config options for various aspects of an adapter.
type Options struct {
	FilePath string
	FileType string
	FileName string
}
