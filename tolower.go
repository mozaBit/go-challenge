package piscine

func ToLower(s string) string {
	zS := []rune(s)
	for i := range s {
		if !(rune(s[i]) > 64 && rune(s[i]) < 91) {
		} else {
			zS[i] = (rune(s[i]) + 32)
		}
	}
	return string(zS)
}
