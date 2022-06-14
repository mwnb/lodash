package slice

func Map[T any](s []T, callback func(v T, i int) T) []T {
	rtS := make([]T, len(s))
	for i, v := range s {
		rtS[i] = callback(v, i)
	}
	return rtS
}
