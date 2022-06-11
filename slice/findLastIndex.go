package slice

func FindLastIndex[T any](s []T, callback func(v T, i int) bool) int {
	for i := len(s); i > -1; i-- {
		exist := callback(s[i], i)
		if exist {
			return i
		}
	}
	return -1
}
