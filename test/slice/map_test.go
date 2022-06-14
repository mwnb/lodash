package slice

import (
	"fmt"
	"testing"

	"github.com/mwnb/lodash/slice"
)

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s = slice.Map(s, func(v int, i int) int {
		return v * 10
	})
	fmt.Println("test-map -> s:", s)
}
