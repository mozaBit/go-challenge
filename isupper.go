package piscine

func IsUpper(s string) bool {
	for i := range s {
		if !(rune(s[i]) > 64 && rune(s[i]) < 91) {
			return false
		}
	}
	return true
}
