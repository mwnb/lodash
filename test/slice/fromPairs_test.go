package slice_test

import (
	"fmt"
	"testing"

	"github.com/mwnb/lodash/slice"
)

func TestFromPairs(t *testing.T) {
	s := [][2]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	m := slice.FromPairs(s)
	fmt.Println(m)
}
