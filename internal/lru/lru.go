package lru

type LruNode struct {
	Value string
	Prev  *LruNode
	Next  *LruNode
}

type Lru struct {
	Capacity int
	Entries  map[int]*LruNode
	Head     *LruNode
	Tail     *LruNode
}

func NewLru(capacity int) *Lru {
	return &Lru{
		Capacity: capacity,
		Entries:  make(map[int]*LruNode, 0),
		Head:     nil,
		Tail:     nil,
	}
}

func (l *Lru) Put(key int, val string) {
	noOfEntries := len(l.Entries)
	freeCapacity := true
	if noOfEntries >= l.Capacity {
		freeCapacity = false
	}

	currentHead := l.Head

	if !freeCapacity {
		// remove last element
		currentTail := l.Tail
		currentTail.Prev.Next = nil
		l.Entries[key] = nil
	}

	// add new entry then
	newEntry := &LruNode{
		Value: val,
		Prev:  nil,
		Next:  currentHead,
	}
	l.Entries[key] = newEntry
	currentHead.Prev = newEntry
}

func (l *Lru) Get(key int) *string {
	val := l.Entries[key]

	if val == nil {
		return nil
	}

	currentHead := l.Head

	// move accessed element to the top
	val.Prev = nil
	val.Next = currentHead
	currentHead.Prev = val

	return &val.Value
}
