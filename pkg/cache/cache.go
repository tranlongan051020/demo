package cache

import (
	"time"

	"demo/pkg/cache/ristretto"
)

// cache interface
type Cache interface {
	Set(key string, value any, ttl time.Duration) (setAble bool)
	Get(key string) (value any, found bool)
}

type CacheProvider string

// list cache providers
var (
	// configs: NumCounters, MaxCost, BufferItems
	RistrettoProvider CacheProvider = "Ristretto"
)

type Config struct {
	Provider    CacheProvider
	NumCounters int64
	MaxCost     int64
	BufferItems int64
}

// factory
func New(config Config) Cache {
	switch config.Provider {
	case RistrettoProvider:
		return ristretto.New(ristretto.Config{
			NumCounters: config.NumCounters,
			MaxCost:     config.MaxCost,
			BufferItems: config.BufferItems,
		})
	default:
		return ristretto.New(ristretto.Config{
			NumCounters: config.NumCounters,
			MaxCost:     config.MaxCost,
			BufferItems: config.BufferItems,
		})
	}
}
