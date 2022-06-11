package slice

func Concat[T any](s ...[]T) []T {
	rtS := make([]T, 0)
	for _, v := range s {
		rtS = append(rtS, v...)
	}
	return rtS
}
