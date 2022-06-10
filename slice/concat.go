package slice

func Concat(s ...[]interface{}) []interface{} {
	rtS := make([]interface{}, 0)
	for _, v := range s {
		rtS = append(rtS, v...)
	}
	return rtS
}
