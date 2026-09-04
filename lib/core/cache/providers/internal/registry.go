package cache_internal

import "fmt"

// must acquire mutex before calling
func isEntryKeyAvailable(key string) error {
	for _, entry := range cache {
		if entry.meta.key == key {
			return fmt.Errorf(`key "%s" is already occupied`, key)
		}
	}

	return nil
}

func RegisterEntry(key string) error {
	mu.Lock()
	defer mu.Unlock()
	
	if err := isEntryKeyAvailable(key); err != nil {
		return err
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