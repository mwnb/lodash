package slice

import (
	"github.com/mwnb/lodash/base"
)

func Union[T base.BaseType](s ...[]T) []T {
	rtS := make([]T, 0)
	for _, ss := range s {
		for _, v := range ss {
			i := FindIndex(rtS, func(val T, idx int) bool {
				return val == v
			})
			if i == -1 {
				rtS = append(rtS, v)
			}
		}
	}
	return rtS
}
