package helpers

func IsCharNumeric(ch rune) bool {
	return ('0' <= ch && ch <= '9')
}
