package cache_internal

type CacheEntry struct {
	meta struct {
		key string
	}
	data any
}
type Cache []CacheEntry
