package slice

func First[T any](s []T, d T) T {
	if len(s) == 0 {
		return d
	}
	return s[0]
}
