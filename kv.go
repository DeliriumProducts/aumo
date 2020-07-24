package aumo

import "time"

// KeyValue defines a key value store
type KeyValue interface {
	Set(key, value string, expiry time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}
