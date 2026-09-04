package cache_internal

type CacheEntry struct {
	meta struct {
		key string
	}
	data CacheEntryData
}

type CacheEntryData any

type Cache []CacheEntry
