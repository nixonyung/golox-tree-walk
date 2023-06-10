package helpers

func IsValueTruthy(val any) bool {
	if val == nil {
		return false
	} else if val, ok := val.(bool); ok {
		return val
	} else {
		return true // strings and numbers are always true
	}
}
