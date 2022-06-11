package slice

import (
	"fmt"
	"lodash/slice"
	"testing"
)

func TestFindIndex(t *testing.T) {
	s := []User{
		User{"xiaoma"},
		User{"mwnb"},
		User{"mw"},
		User{"mawei"},
	}

	i := slice.FindIndex(s, func(v User, i int) bool {
		return v.Name == "mw"
	})
	fmt.Println("i", i)
	if i != 2 {
		t.Error("slice.FindIndex test fail! i!=1")
	}

}

type User struct {
	Name string
}
