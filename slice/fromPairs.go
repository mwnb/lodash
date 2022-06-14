package slice

import "github.com/mwnb/lodash/base"

func FromPairs[T base.BaseType](s [][2]T) map[T]T {
	rtM := make(map[T]T)
	for _, arr := range s {
		rtM[arr[0]] = arr[1]
	}
	return rtM
}
