package cache

import "time"

type Handler interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, exp ...time.Duration) error
}
