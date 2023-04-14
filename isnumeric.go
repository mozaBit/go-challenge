package piscine

func IsNumeric(s string) bool {
	for i := range s {
		if !(rune(s[i]) > 47 && rune(s[i]) < 58) {
			return false
		}
	}
	return true
}
