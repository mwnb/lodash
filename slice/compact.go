package slice

func Compact(s []interface{}) []interface{} {
	rtS := make([]interface{}, 0)
	for _, v := range s {
		switch v.(type) {
		case int:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case int8:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case int16:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case int32:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case int64:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case uint:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case uint8:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case uint16:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case uint32:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case uint64:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case float32:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case float64:
			if v != 0 {
				rtS = append(rtS, v)
			}
		case bool:
			if v != false {
				rtS = append(rtS, v)
			}
		case string:
			if v != "" {
				rtS = append(rtS, v)
			}
		default:
			if v != nil {
				rtS = append(rtS, v)
			}
		}

	}
	return rtS
}
