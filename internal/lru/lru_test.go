package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLru1(t *testing.T) {
	lru := NewLru(3)
	lru.Put(1, "one")
	lru.Put(2, "two")
	lru.Put(3, "three")

	assert.Equal(t, []LruCacheItem{
		{Key: 1, Value: "one"},
		{Key: 2, Value: "two"},
		{Key: 3, Value: "three"},
	}, lru.ToList())

	// one is removed
	lru.Put(4, "four")
	assert.Equal(t, []LruCacheItem{
		{Key: 2, Value: "two"},
		{Key: 3, Value: "three"},
		{Key: 4, Value: "four"},
	}, lru.ToList())

	// two is removed
	lru.Put(5, "five")
	assert.Equal(t, []LruCacheItem{
		{Key: 3, Value: "three"},
		{Key: 4, Value: "four"},
		{Key: 5, Value: "five"},
	}, lru.ToList())

	// three is removed
	lru.Put(6, "six")
	assert.Equal(t, []LruCacheItem{
		{Key: 4, Value: "four"},
		{Key: 5, Value: "five"},
		{Key: 6, Value: "six"},
	}, lru.ToList())
}

func TestLru2(t *testing.T) {
	lru := NewLru(3)
	lru.Put(1, "one")
	lru.Put(2, "two")
	lru.Put(3, "three")

	assert.Equal(t, []LruCacheItem{
		{Key: 1, Value: "one"},
		{Key: 2, Value: "two"},
		{Key: 3, Value: "three"},
	}, lru.ToList())

	_ = lru.Get(1)

	// two is removed, one was just accessed making it LRU!
	lru.Put(4, "four")
	assert.Equal(t, []LruCacheItem{
		{Key: 1, Value: "one"},
		{Key: 3, Value: "three"},
		{Key: 4, Value: "four"},
	}, lru.ToList())
}

func TestLru3(t *testing.T) {
	lru := NewLru(3)
	lru.Put(1, "one")
	lru.Put(2, "two")
	lru.Put(3, "three")

	assert.Equal(t, []LruCacheItem{
		{Key: 1, Value: "one"},
		{Key: 2, Value: "two"},
		{Key: 3, Value: "three"},
	}, lru.ToList())

	lru.Put(1, "ONE")

	// two is removed, one was just accessed making it LRU!
	lru.Put(4, "four")
	assert.Equal(t, []LruCacheItem{
		{Key: 1, Value: "ONE"},
		{Key: 3, Value: "three"},
		{Key: 4, Value: "four"},
	}, lru.ToList())
}
