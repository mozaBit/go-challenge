package piscine

func IsAlpha(s string) bool {
	for i := range s {
		if !((rune(s[i]) > 64 && rune(s[i]) < 91) || (rune(s[i]) > 96 && rune(s[i]) < 123) || (rune(s[i]) > 47 && rune(s[i]) < 58) || (rune(s[i]) == 32)) {
			return false
		}
	}
	return true
}
