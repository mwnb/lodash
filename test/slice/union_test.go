package slice

import (
	"fmt"
	"testing"

	"github.com/mwnb/lodash/slice"
)

func TestUnion(t *testing.T) {
	s := slice.Union(
		[]int{1, 2, 3, 4, 5},
		[]int{5, 6, 7},
		[]int{1, 3, 5, 6, 77},
	)
	fmt.Println("slice.Union:", s)
	if len(s) <= 1 {

	}
}
