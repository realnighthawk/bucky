package provider

import (
	"encoding/json"
	"sync"

	"github.com/kumarabd/gokit/config"
)

// Type InMem implements the config interface Handler for an in-memory configuration registry.
type InMem struct {
	store map[string]string
	mutex sync.Mutex
}

// NewInMem returns a new instance of an in-memory configuration provider using the provided Options opts.
func NewInMem(opts Options) (config.Handler, error) {
	return &InMem{
		store: make(map[string]string),
	}, nil
}

// -------------------------------------------Application config methods----------------------------------------------------------------

// SetKey sets a key value in local store
func (l *InMem) SetKey(key string, value string) {
	l.mutex.Lock()
	l.store[key] = value
	l.mutex.Unlock()
}

// GetKey gets a key value from local store
func (l *InMem) GetKey(key string) string {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.store[key]
}

// GetObject gets an object value for the key
func (l *InMem) GetObject(key string, result interface{}) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	err := json.Unmarshal([]byte(l.store[key]), result)
	if err != nil {
		return ErrGetObject(err)
	}
	return nil
}

// SetObject sets an object value for the key
func (l *InMem) SetObject(key string, value interface{}) error {
	l.mutex.Lock()
	val, err := json.Marshal(value)
	defer l.mutex.Unlock()
	if err != nil {
		return ErrSetObject(config.ErrInMem(err))
	}
	l.store[key] = string(val)
	return nil
}
