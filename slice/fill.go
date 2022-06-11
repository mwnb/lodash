package slice

func Fill[T any](s []T, start, end int, v T) []T {
	if start > end || start > len(s) {
		return s
	} else if end > len(s) {
		end = len(s)
	}
	for i := start; i < end; i++ {
		s[i] = v
	}
	return s
}
