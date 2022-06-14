package slice

func ForEach[T any](s []T, callback func(v T, i int)) {
	for i, v := range s {
		callback(v, i)
	}
}
