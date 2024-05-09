package ristretto

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

type RistrettoCache interface {
	Get(key string) (value any, found bool)
	Set(key string, value any, ttl time.Duration) (setAble bool)
}

type Config struct {
	NumCounters int64
	MaxCost     int64
	BufferItems int64
}

type instance struct {
	cache *ristretto.Cache
}

func New(config Config) RistrettoCache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: config.NumCounters, // number of keys to track frequency of (1M).
		MaxCost:     config.MaxCost,     // maximum cost of cache (16MB).
		BufferItems: config.BufferItems, // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
	ins := &instance{
		cache,
	}
	return ins
}

func (ist *instance) Get(key string) (value any, found bool) {
	return ist.cache.Get(key)
}

func (ist *instance) Set(key string, value any, ttl time.Duration) (setAble bool) {
	setAble = ist.cache.SetWithTTL(key, value, 1, ttl)
	ist.cache.Wait()

	return setAble
}
