// Package config provides the interface Handler and errors related to the configuration of adapters.
package config

// Handler is the config interface to fascilitate multiple providers
type Handler interface {
	// SetKey is used to set a string value for a given key.
	SetKey(key string, value string)
	// GetKey is used to retrieve a string value for a given key.
	GetKey(key string) string

	// GetObject is used to retrieve an object for a given key and a given interface representing that object in result.
	GetObject(key string, result interface{}) error
	// SetObject is used to set an object for a given key and a given interface representing that object in result.
	SetObject(key string, value interface{}) error

	// GetAll is used to retrieve all objects.
	GetAll(result interface{}) error
}
