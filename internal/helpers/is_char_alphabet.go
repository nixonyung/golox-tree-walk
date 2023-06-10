package helpers

func IsCharAlphabet(ch rune) bool {
	return ('A' <= ch && ch <= 'Z') ||
		('a' <= ch && ch <= 'z')
}
