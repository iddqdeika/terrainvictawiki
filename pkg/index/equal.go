package index

import (
	"strconv"
	"strings"
)

func intIndexator[V any](f func(v *V) int) *genericIndexator[int, V] {
	return &genericIndexator[int, V]{
		KeyGen:    f,
		KeyParser: intParser,
	}
}

func intParser(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func newEqualIndex[K string | int, V any](ind Indexator[K, V]) *equalIndex[K, V] {
	return &equalIndex[K, V]{
		f:    ind,
		data: make(map[K]*V),
	}
}

type equalIndex[K string | int, V any] struct {
	f    Indexator[K, V]
	data map[K]*V
}

func (e *equalIndex[K, V]) Put(v *V) {
	e.data[e.f.Key(v)] = v
}

func (e *equalIndex[K, V]) Search(condition string) (*V, error) {
	c := strings.TrimLeft(condition, "= ")
	v, ok := e.data[e.f.ParseKey(c)]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}
