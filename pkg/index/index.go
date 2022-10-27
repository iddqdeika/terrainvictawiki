package index

import (
	"errors"
)

type Indexator[K comparable, V any] interface {
	Key(v *V) K
	ParseKey(k string) K
}

type Index[V any] interface {
	Put(*V)
	Search(condition string) (*V, error)
}

var (
	ErrNotFound = errors.New("not found")
)

type genericIndexator[K comparable, V any] struct {
	KeyGen    func(v *V) K
	KeyParser func(s string) K
}

func (i *genericIndexator[K, V]) Key(v *V) K {
	return i.KeyGen(v)
}

func (i *genericIndexator[K, V]) ParseKey(s string) K {
	return i.KeyParser(s)
}
