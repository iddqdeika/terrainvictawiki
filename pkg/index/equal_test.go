package index

import (
	"testing"
)

type SampleStruct struct {
	No int
}

func TestSomeShit(t *testing.T) {
	var ind Indexator[int, SampleStruct]
	ind = intIndexator(func(v *SampleStruct) int {
		return v.No
	})

	index := newEqualIndex(ind)

	index.Put(&SampleStruct{No: 666})
	res, err := index.Search("= 666")
	if err != nil || res.No != 666 {
		t.Errorf("wrong put-search result")
	}
}
