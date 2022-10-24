package index

import "sync"

type Indexable[K comparable] interface {
	ID() K
}

func NewStorage[K comparable, V Indexable[K]]() *IndexedStorage[K, V] {
	return &IndexedStorage[K, V]{
		index: make(map[K]int),
	}
}

type IndexedStorage[K comparable, V Indexable[K]] struct {
	list []V
	// stores index in list by key (ID)
	index map[K]int
	m     sync.RWMutex
}

func (i *IndexedStorage[K, V]) Put(t V) {
	key := t.ID()
	i.m.Lock()
	defer i.m.Unlock()
	// try to find item in index
	index, ok := i.index[key]
	// in case t already exists in index - rewrite it
	if ok {
		i.list[index] = t
		return
	}
	// create new index item
	i.addItem(t, key)
}

func (i *IndexedStorage[K, V]) Get(key K) (V, bool) {
	i.m.RLock()
	defer i.m.RUnlock()
	index, ok := i.index[key]
	if ok {
		return i.list[index], true
	}
	return *new(V), false
}

func (i *IndexedStorage[K, V]) GetAll() []V {
	return i.list
}

func (i *IndexedStorage[K, V]) addItem(t V, key K) {
	index := len(i.list)
	i.list = append(i.list, t)
	i.index[key] = index
}

func findIndex[K comparable, V Indexable[K]](arr []V, key K) int {
	for i, t := range arr {
		if t.ID() == key {
			return i
		}
	}
	return -1
}
