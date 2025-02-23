package lru

import (
	"container/list"
	"fmt"
	"sort"
)

/*
https://dev.to/johnscode/implement-an-lru-cache-in-go-1hbc
https://github.com/johnscode/gocodingchallenges
*/

type LruCacheItem struct {
	Key   int
	Value string
}

type ByKey []LruCacheItem

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

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
	if val_, ok := l.entries[key]; !ok /* pushing new element */ {
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
	} else /* overwriting existing element, new value for same key! */ {
		// update value for given key
		l.entries[key] = LruItem{
			Value:           val,
			LastUsedElement: val_.LastUsedElement,
		}

		// make it LRU
		l.lastUsed.MoveToFront(val_.LastUsedElement)
	}
}

func (l *Lru) Get(key int) *string {
	if val, ok := l.entries[key]; !ok {
		return nil
	} else {
		l.lastUsed.MoveToFront(val.LastUsedElement) // make this key LRU
		return &val.Value                           // return corresponding value
	}
}

func (l *Lru) ToList() []LruCacheItem {
	result := make([]LruCacheItem, 0)
	for key, val := range l.entries {
		result = append(result, LruCacheItem{
			Key:   key,
			Value: val.Value,
		})
	}

	sort.Sort(ByKey(result))
	return result
}

func (l *Lru) printLru() {
	fmt.Println("-----")
	for key, val := range l.entries {
		fmt.Printf("key: %d, val: %s\n", key, val.Value)
	}
	fmt.Println("-----")
}
