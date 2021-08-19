package memcache

import (
	"sync"
	"time"

	"github.com/cespare/xxhash/v2"
)

const (
	bucketCount          = 256
	defaultKeyCount      = 8
	defaultCleanInterval = 10 * time.Second
)

type Value struct {
	Val      interface{}
	ExpireAt *time.Time // zero time value means, this key never expire
}

type kvstore struct {
	mu sync.Mutex
	m  map[string]*Value
}

func (kv *kvstore) Get(key string) (*Value, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	v, ok := kv.m[key]
	return v, ok
}

func (kv *kvstore) Set(key string, value *Value) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.m[key] = value
}

func (kv *kvstore) Del(key string) (*Value, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	v, ok := kv.m[key]
	delete(kv.m, key)
	return v, ok
}

func (kv *kvstore) RemoveExpireKeys(now time.Time) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	keys := []string{}
	for k, v := range kv.m {
		if v.ExpireAt.Unix()-now.Unix() <= 0 {
			keys = append(keys, k)
		}
	}

	for _, key := range keys {
		delete(kv.m, key)
	}
}

type cache struct {
	data          [bucketCount]*kvstore
	cleanInterval time.Duration
}

func NewCache(initKeyCount uint32, cleanInterval time.Duration) *cache {
	if initKeyCount <= 0 {
		initKeyCount = defaultKeyCount
	}
	c := &cache{
		cleanInterval: cleanInterval,
	}

	if c.cleanInterval <= 0 {
		c.cleanInterval = defaultCleanInterval
	}

	for i := 0; i < bucketCount; i++ {
		c.data[i] = &kvstore{m: make(map[string]*Value, initKeyCount)}
	}

	go c.checkAndRemoveExpireKeys()
	return c
}

func (c *cache) Get(key string) (interface{}, bool) {
	i := xxhash.Sum64String(key) % bucketCount
	bucket := c.data[i]
	v, ok := bucket.Get(key)
	if !ok {
		return nil, false
	}

	if v.ExpireAt.IsZero() {
		return v.Val, true
	}

	now := time.Now()
	// this value expired
	if v.ExpireAt.Unix()-now.Unix() <= 0 {
		return nil, false
	}

	return v.Val, true
}

// Set set key value that never expire
func (c *cache) Set(key string, value interface{}) {
	c.setKeyWithExpire(key, value, &time.Time{})
}

// SetWithTTL set a key that will expire after ttl
func (c *cache) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	expireAt := time.Now().Add(ttl)
	c.setKeyWithExpire(key, value, &expireAt)
}

func (c *cache) Delete(key string) (interface{}, bool) {
	i := xxhash.Sum64String(key) % bucketCount
	bucket := c.data[i]
	v, ok := bucket.Del(key)
	if !ok {
		return nil, false
	}

	return v.Val, ok
}

func (c *cache) setKeyWithExpire(key string, value interface{}, expireAt *time.Time) {
	i := xxhash.Sum64String(key) % bucketCount
	bucket := c.data[i]

	v := &Value{Val: value, ExpireAt: expireAt}
	bucket.Set(key, v)
}

func (c *cache) checkAndRemoveExpireKeys() {
	t := time.NewTicker(c.cleanInterval)
	defer t.Stop()

	for now := range t.C {
		c.removeExpireKeys(now)
	}
}

func (c *cache) removeExpireKeys(now time.Time) {
	for i := 0; i < len(c.data); i++ {
		data := c.data[i]
		go data.RemoveExpireKeys(now)
	}
}
