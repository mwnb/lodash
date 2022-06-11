package base

func Include[T BaseType](s []T, v T) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}
