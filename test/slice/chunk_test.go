package slice

import (
	"lodash/slice"
	"testing"
)

func TestFirst(t *testing.T) {
	s := []int{100, 200, 300}
	v := slice.First(s, 110)
	if v != 100 {
		t.Error("slice.First test fail!")
	}

}
