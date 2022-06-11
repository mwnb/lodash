package slice

import (
	"testing"

	"github.com/mwnb/lodash/slice"
)

func TestNth(t *testing.T) {
	s := []int{100, 200, 300, 400, 500}
	v := slice.Nth(s, -1, 110)
	if v != 500 {
		t.Error("slice.Nth tes fail! v != 500")
	}
	v = slice.Nth(s, 0, 110)
	if v != 100 {
		t.Error("slice.Nth tes fail! v != 100")
	}
	v = slice.Nth(s, -3, 110)
	if v != 300 {
		t.Error("slice.Nth tes fail! v != 300")
	}
	v = slice.Nth(s, 4, 110)
	if v != 500 {
		t.Error("slice.Nth tes fail! v != 500")
	}
}
