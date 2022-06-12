package slice_test

import (
	"testing"

	"github.com/mwnb/lodash/slice"
)

func TestFirst(t *testing.T) {
	s := []int{100, 200, 300}
	v := slice.First(s, 110)
	if v != 100 {
		t.Error("slice.First test fail!")
	}

}
