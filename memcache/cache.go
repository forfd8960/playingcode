package memcache

import (
	"sync"

	"github.com/cespare/xxhash/v2"
)

const (
	bucketCount = 256
)

type kvstore struct {
	mu sync.Mutex
	m  map[string]interface{}
}

func (kv *kvstore) Get(key string) (interface{}, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	v, ok := kv.m[key]
	return v, ok
}

func (kv *kvstore) Set(key string, value interface{}) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.m[key] = value
}

type cache struct {
	data [bucketCount]*kvstore
}

func NewCache(initKeyCount uint32) *cache {
	c := &cache{}
	for i := 0; i < bucketCount; i++ {
		c.data[i] = &kvstore{m: make(map[string]interface{}, initKeyCount)}
	}
	return c
}

func (c *cache) Get(key string) (interface{}, bool) {
	i := xxhash.Sum64String(key) % bucketCount
	bucket := c.data[i]
	return bucket.Get(key)
}

func (c *cache) Set(key string, value interface{}) {
	i := xxhash.Sum64String(key) % bucketCount
	bucket := c.data[i]
	bucket.Set(key, value)
}
