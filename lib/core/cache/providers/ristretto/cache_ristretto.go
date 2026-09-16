package cache_ristretto

import (
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto/v2"
)

type Cache[K comparable, V any] struct {
	cache *ristretto.Cache[string, V]
}

func New[K comparable, V any]() *Cache[K, V] {
	cache, err := ristretto.NewCache(&ristretto.Config[string, V]{
		NumCounters: 1e4,
		MaxCost:     100 << 20,
		BufferItems: 64,
	})
	if err != nil {
		panic(fmt.Errorf("ristretto cache could not be initialized: %w", err))
	}

	return &Cache[K, V]{
		cache: cache,
	}
}

func (c *Cache[K, V]) Set(key string, value V, cost int64, ttl time.Duration) {
	c.cache.SetWithTTL(key, value, cost, ttl)
}

func (c *Cache[K, V]) Get(key string) (V, bool) {
	return c.cache.Get(key)
}