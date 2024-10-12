package pokecache

import "time"

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCacheEntry(val []byte) CacheEntry {
	return CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
