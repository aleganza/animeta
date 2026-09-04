package cache_internal

import "fmt"

func findEntry(key string) (int, bool) {
	for i, entry := range cache {
		if entry.meta.key == key {
			return i, true
		}
	}

	return -1, false
}

func doesEntryKeyExist(key string) bool {
	_, ok := findEntry(key)

	return ok
}

func RegisterEntry(key string) error {
	mu.Lock()
	defer mu.Unlock()

	if keyExists := doesEntryKeyExist(key); keyExists {
		return fmt.Errorf(`cache with key "%s" is already occupied`, key)
	}

	var entry CacheEntry
	entry.meta.key = key

	cache = append(cache, entry)

	return nil
}

func PrintEntries() {
	mu.RLock()
	defer mu.RUnlock()

	for _, entry := range cache {
		fmt.Println(entry.meta.key)
	}
}
