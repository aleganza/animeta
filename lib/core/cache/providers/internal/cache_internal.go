package cache_internal

import (
	"fmt"
	"sync"
)

var cache Cache
var mu sync.RWMutex

func Get(key string) (CacheEntryData, error) {
	mu.RLock()
	defer mu.RUnlock()

	entryIndex, ok := findEntry(key)
	if !ok {
		return CacheEntry{}, fmt.Errorf(`cache with key "%s" does not exist`, key)
	}

	return cache[entryIndex].data, nil
}

func Set(key string, data any) error {
	mu.Lock()
	defer mu.Unlock()

	i, ok := findEntry(key)

	if !ok {
		return fmt.Errorf(`ache with key "%s" does not exist`, key)
	}

	cache[i].data = data

	return nil
}
