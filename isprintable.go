package piscine

func IsPrintable(s string) bool {
	for i := range s {
		if !(rune(s[i]) > 31 && rune(s[i]) < 128) {
			return false
		}
	}
	return true
}
