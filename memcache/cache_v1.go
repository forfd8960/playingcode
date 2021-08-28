package memcache

import (
	"sync"
	"time"
)

type cacheV1 struct {
	data          sync.Map
	cleanInterval time.Duration
}

func NewCacheV1(cleanInternal time.Duration) *cacheV1 {
	return &cacheV1{
		cleanInterval: cleanInternal,
	}
}

func (c *cacheV1) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

func (c *cacheV1) Get(key string) (interface{}, bool) {
	return c.data.Load(key)
}
