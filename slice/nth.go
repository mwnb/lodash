package slice

import "math"

func Nth[T any](s []T, i int, v T) T {
	l := len(s)
	if i > 0 && i >= l || i < 0 && math.Abs(float64(i)) >= float64(l) {
		return v
	} else if i < 0 {
		i = len(s) - int(math.Abs(float64(i)))
	}
	return s[i]
}
