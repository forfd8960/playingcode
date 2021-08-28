package main

import (
	"fmt"
	"time"

	"github.com/forfd8960/playingcode/memcache"
)

func main() {
	cache := memcache.NewCache(128, 1*time.Second)
	val, ok := cache.Get("test-key1")
	fmt.Printf("val: %v, ok: %v\n", val, ok)

	key := "test-key1"
	cache.Set(key, 100)
	val, _ = cache.Get(key)
	fmt.Printf("key: %v, value: %v\n", key, val)

	key2 := "test-key2"
	cache.SetWithTTL(key2, 1000, 2*time.Second)
	val, _ = cache.Get(key2)
	fmt.Printf("key: %v, value: %v\n", key2, val)

	fmt.Println("wait for 2 second...")
	time.Sleep(time.Second * 2)
	val, ok = cache.Get(key2)
	fmt.Printf("key: %v, value: %v, ok: %v\n", key2, val, ok)

	fmt.Println("wait another 2 second...")
	time.Sleep(time.Second * 2)
	val, ok = cache.Get(key2)
	fmt.Printf("key: %v, value: %v, ok: %v\n", key2, val, ok)
}
