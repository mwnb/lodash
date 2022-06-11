package slice

func FindIndex[T any](s []T, callback func(v T, i int) bool) int {
	for i := 0; i < len(s); i++ {
		exist := callback(s[i], i)
		if exist {
			return i
		}
	}
	return -1
}
