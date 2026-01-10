package go_kit

func Includes(arr []any, value any) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
