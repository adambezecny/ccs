package lru

import "testing"

func TestLru(t *testing.T) {
	lru := NewLru(3)
	lru.Put(1, "one")
	lru.Put(2, "two")
	lru.Put(3, "three")

	lru.printLru()

	lru.Get(1)
	lru.Put(4, "four")
	lru.printLru()

	lru.Put(5, "five")
	lru.printLru()

	lru.Put(6, "six")
	lru.printLru()
}
