package slice

func Drop[T any](s []T, i int) []T {
	if i > len(s) {
		i = len(s)
	}
	return s[i:]
}
