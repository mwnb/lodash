package slice

import "lodash/base"

func Difference[T base.BaseType](s ...[]T) []T {
	if len(s) == 1 {
		return s[0]
	}
	rtS := make([]T, 0)
	comparison := Concat(s[1:]...)
	for _, v := range s[0] {
		if !base.Include(comparison, v) {
			rtS = append(rtS, v)
		}
	}
	return rtS
}
