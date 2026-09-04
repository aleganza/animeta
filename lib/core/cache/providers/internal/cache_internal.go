package cache_internal

import "sync"

var cache Cache
var mu sync.RWMutex
