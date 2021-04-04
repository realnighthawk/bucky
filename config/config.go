// Package config provides the interface Handler and errors related to the configuration of adapters.
package config

// Provided implementations can be found in the package config/provider.
type Handler interface {
	// SetKey is used to set a string value for a given key.
	SetKey(key string, value string)

	// GetKey is used to retrieve a string value for a given key.
	GetKey(key string) string

	// GetObject is used to retrieve an object for a given key and a given interface representing that object in result.
	// An example of such an object is map[string]string. These objects can e.g. be set in the factory function for a specific
	// config provider implementation.
	GetObject(key string, result interface{}) error

	SetObject(key string, value interface{}) error
}
