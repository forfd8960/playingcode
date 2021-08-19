package memcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
