package memcache

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type kvpair struct {
	key   string
	value interface{}
}

func TestCacheGet(t *testing.T) {
	c := NewCache(128, 2*time.Second)
	v, ok := c.Get("keyNotExists")
	assert.False(t, ok)
	assert.Nil(t, v)
}

func TestCacheSet(t *testing.T) {
	c := NewCache(128, 2*time.Second)
	v, ok := c.Get("keyNotExists")
	assert.False(t, ok)
	assert.Nil(t, v)

	c.Set("keyNotExists", "value")

	v, ok = c.Get("keyNotExists")
	assert.True(t, ok)
	assert.Equal(t, v, "value")
}

func BenchmarkCacheSet(b *testing.B) {
	kvpairs := prepareKvpairs(b.N)
	cache := NewCache(128, 2*time.Second)

	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			cache.Set(kvpairs[i].key, kvpairs[i].value)
			i++
		}
	})
}

func BenchmarkCacheGet(b *testing.B) {
	kvpairs := prepareKvpairs(b.N)
	cache := NewCache(128, 2*time.Second)
	for _, kv := range kvpairs {
		cache.Set(kv.key, kv.value)
	}

	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			cache.Get(kvpairs[i].key)
			i++
		}
	})

}

func BenchmarkCacheSetV1(b *testing.B) {
	kvpairs := prepareKvpairs(b.N)
	cache := NewCacheV1(2 * time.Second)

	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			cache.Set(kvpairs[i].key, kvpairs[i].value)
			i++
		}
	})
}

func BenchmarkCacheGetV1(b *testing.B) {
	kvpairs := prepareKvpairs(b.N)
	cache := NewCacheV1(2 * time.Second)
	for _, kv := range kvpairs {
		cache.Set(kv.key, kv.value)
	}

	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			cache.Get(kvpairs[i].key)
			i++
		}
	})
}

func prepareKvpairs(n int) (result []*kvpair) {
	const keyLength = 10
	for i := 0; i < n; i++ {
		result = append(result, &kvpair{
			key:   generateRandomKey(keyLength),
			value: i,
		})
	}
	return result
}

func generateRandomKey(n int) string {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bs)
}
