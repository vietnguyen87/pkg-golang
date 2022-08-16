package cache

import (
	"context"
	"encoding/json"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, data interface{}) error
	SetWithExpiration(ctx context.Context, key string, data interface{}, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	AllKeys(ctx context.Context, pattern string) ([]string, error)
}

type RedisCache interface {
	Cache
	HGet(ctx context.Context, key, field string, data interface{}) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HSet(ctx context.Context, key, field string, data interface{}) error
	HDel(ctx context.Context, key string, fields ...string) error
	Incr(ctx context.Context, key string) (result int64, err error)
	Decr(ctx context.Context, key string) (result int64, err error)
	SetNX(ctx context.Context, key string, data interface{}) (result bool, err error)
}

// KeyFunc defines a transformer for cache keys
type KeyFunc func(s string) string

// DefaultKeyFunc is the default implementation of cache keys
// All it does is to return the key sent in by client code
func DefaultKeyFunc(s string) string {
	return s
}

type Serializer interface {
	Serialize(data interface{}) (string, error)
	Deserialize(data string, target interface{}) error
}

type JSONSerializer struct{}

func (s JSONSerializer) Serialize(data interface{}) (string, error) {
	return Serialize(data)
}

func (s JSONSerializer) Deserialize(data string, target interface{}) error {
	return Deserialize(data, target)
}

func Serialize(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func Deserialize(data string, target interface{}) error {
	b := []byte(data)
	return json.Unmarshal(b, target)
}
