package memcache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheGet(t *testing.T) {
	c := NewCache(128)
	v, ok := c.Get("keyNotExists")
	assert.False(t, ok)
	assert.Nil(t, v)
}

func TestCacheSet(t *testing.T) {
	c := NewCache(128)
	v, ok := c.Get("keyNotExists")
	assert.False(t, ok)
	assert.Nil(t, v)

	c.Set("keyNotExists", "value")

	v, ok = c.Get("keyNotExists")
	assert.True(t, ok)
	assert.Equal(t, v, "value")
}
