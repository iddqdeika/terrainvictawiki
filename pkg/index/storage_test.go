package index

import "testing"

type IndexItemStub struct {
	id string
}

func (i *IndexItemStub) ID() string {
	return i.id
}

func TestIndexPutGet(t *testing.T) {
	s := NewStorage[string, *IndexItemStub]()

	s.Put(&IndexItemStub{id: "some"})
	v, ok := s.Get("some")
	if !ok || v.id != "some" {
		t.Errorf("incorrect result after put-get")
	}
	if len(s.GetAll()) != 1 {
		t.Errorf("incorrect put-getall result")
	}
}
