package piscine

func IsLower(s string) bool {
	for i := range s {
		if !(rune(s[i]) > 96 && rune(s[i]) < 123) {
			return false
		}
	}
	return true
}
