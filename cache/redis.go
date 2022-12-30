package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vietnguyen87/pkg-golang/xerrors"
	"github.com/vietnguyen87/pkg-golang/xredis"
)

const (
	defaultExpiration = 24 * time.Hour
)

type redisCache struct {
	redisCmd   xredis.Cmdable
	expiration time.Duration
	keyFn      KeyFunc
	serializer Serializer
}

// New instance of Cache with default expiration for cached object is 12 hours
func New(redisCmd xredis.Cmdable, opts ...Option) RedisCache {
	options := &cacheOptions{
		expiration: defaultExpiration,
		keyFunc:    DefaultKeyFunc,
		serializer: &JSONSerializer{},
	}

	for _, opt := range opts {
		opt.apply(options)
	}

	return &redisCache{
		redisCmd,
		options.expiration,
		options.keyFunc,
		options.serializer,
	}
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	cacheKey := r.keyFn(key)
	jsonData, err := r.redisCmd.Get(ctx, cacheKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", xerrors.NotFound.Newf("key [%s] not found", cacheKey)
		}

		return "", xerrors.CannotGetFromCache.Wrap(err, "failed to get data from cache")
	}
	return jsonData, nil
}

func (r *redisCache) HGet(ctx context.Context, key, field string, data interface{}) error {
	cacheKey := r.keyFn(key)
	jsonData, err := r.redisCmd.HGet(ctx, cacheKey, field).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return xerrors.NotFound.Newf("key [%s] not found", cacheKey)
		}

		return xerrors.CannotGetFromCache.Wrap(err, "failed to get data from cache")
	}

	err = r.serializer.Deserialize(jsonData, data)
	if err != nil {
		return xerrors.DeserializingError.Wrap(err, "failed to deserialize data from cache")
	}

	return nil
}

func (r *redisCache) HSet(ctx context.Context, key, field string, data interface{}) error {
	jsonData, err := r.serializer.Serialize(data)
	if err != nil {
		return xerrors.SerializingError.Wrap(err, "failed to serialize data")
	}

	cacheKey := r.keyFn(key)
	_, err = r.redisCmd.HSet(ctx, cacheKey, field, jsonData).Result()

	if err != nil {
		return xerrors.CannotSaveToCache.Wrap(err, "failed to save data to cache")
	}

	return nil
}

func (r *redisCache) HDel(ctx context.Context, key string, fields ...string) error {
	cacheKey := r.keyFn(key)
	_, err := r.redisCmd.HDel(ctx, cacheKey, fields...).Result()

	if err != nil {
		return xerrors.CannotDeleteFromCache.Wrap(err, "failed to delete from cache")
	}

	return nil
}

// TODO: pass in slice pointer for result destination like GORM
func (r *redisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	cacheKey := r.keyFn(key)
	m, err := r.redisCmd.HGetAll(ctx, cacheKey).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, xerrors.NotFound.Newf("key [%s] not found", cacheKey)
		}

		return nil, xerrors.CannotGetFromCache.Wrap(err, "failed to get data from cache")
	}
	return m, nil
}

func (r *redisCache) Set(ctx context.Context, key string, data interface{}) error {
	return r.SetWithExpiration(ctx, key, data, r.expiration)
}

func (r *redisCache) SetWithExpiration(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	if expiration <= 0 {
		expiration = r.expiration
	}

	cacheKey := r.keyFn(key)
	_, err := r.redisCmd.Set(ctx, cacheKey, data, expiration).Result()

	if err != nil {
		return xerrors.CannotSaveToCache.Wrap(err, "failed to save data to cache")
	}

	return nil
}

func (r *redisCache) Delete(ctx context.Context, key string) error {
	cacheKey := r.keyFn(key)
	_, err := r.redisCmd.Del(ctx, cacheKey).Result()
	if err != nil {
		return xerrors.CannotDeleteFromCache.Wrapf(err, "cannot delete key %s from cache", cacheKey)
	}

	return nil
}

func (r *redisCache) AllKeys(ctx context.Context, pattern string) ([]string, error) {
	return r.redisCmd.Keys(ctx, pattern).Result()
}

func (r *redisCache) Incr(ctx context.Context, key string) (result int64, err error) {
	cacheKey := r.keyFn(key)
	result, err = r.redisCmd.Incr(ctx, cacheKey).Result()
	if err != nil {
		err = xerrors.CannotIncrInCache.Wrapf(err, "cannot increment key %s in cache", cacheKey)
		return
	}

	return
}

func (r *redisCache) Decr(ctx context.Context, key string) (result int64, err error) {
	cacheKey := r.keyFn(key)
	result, err = r.redisCmd.Decr(ctx, cacheKey).Result()
	if err != nil {
		err = xerrors.CannotDecrInCache.Wrapf(err, "cannot decrement key %s in cache", cacheKey)
		return
	}

	return
}

func (r *redisCache) SetNX(ctx context.Context, key string, data interface{}) (result bool, err error) {
	cacheKey := r.keyFn(key)
	result, err = r.redisCmd.SetNX(ctx, cacheKey, data, r.expiration).Result()
	return
}
