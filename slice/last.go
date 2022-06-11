package slice

func Last[T any](s []T, d T) T {
	if len(s) == 0 {
		return d
	}
	return s[len(s)-1]
}
