package slice

import (
	"errors"
	"math"
)

func Chunk[T any](s []T, l int) ([][]T, error) {
	if l < 1 {
		return nil, errors.New("[range err:] l must > 0")
	}
	sz := int(math.Ceil(float64(len(s)) / float64(l)))
	rtS := make([][]T, sz)
	idx := 0
	for i := 0; i < len(s); i += l {
		end := i + l
		if end > len(s) {
			end = len(s)
		}
		rtS[idx] = s[i:end]
		idx++
	}
	return rtS, nil
}
