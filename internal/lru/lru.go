package lru

import (
	"container/list"
	"fmt"
)

/*
https://dev.to/johnscode/implement-an-lru-cache-in-go-1hbc
https://github.com/johnscode/gocodingchallenges
*/

type LruItem struct {
	Value           string
	LastUsedElement *list.Element
}

type Lru struct {
	capacity int
	entries  map[int]LruItem
	lastUsed *list.List
}

func NewLru(capacity int) *Lru {
	return &Lru{
		capacity: capacity,
		entries:  make(map[int]LruItem, 0),
		lastUsed: list.New(),
	}
}
func (l *Lru) Put(key int, val string) {
	pushedElement := l.lastUsed.PushFront(key)

	l.entries[key] = LruItem{
		Value:           val,
		LastUsedElement: pushedElement,
	}

	if l.lastUsed.Len() > l.capacity {
		oldestElement := l.lastUsed.Back()
		oldestElementVal, _ := oldestElement.Value.(int)
		delete(l.entries, oldestElementVal)
		l.lastUsed.Remove(oldestElement)
	}
}

func (l *Lru) Get(key int) *string {
	if val, ok := l.entries[key]; !ok {
		return nil
	} else {
		elem := val.LastUsedElement
		l.lastUsed.MoveToFront(elem)
		return &val.Value
	}
}

func (l *Lru) printLru() {
	fmt.Println("-----")
	for key, val := range l.entries {
		fmt.Printf("key: %d, val: %s\n", key, val.Value)
	}
	fmt.Println("-----")
}
