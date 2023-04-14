package piscine

func ToUpper(s string) string {
	zS := []rune(s)
	for i := range s {
		if !(rune(s[i]) > 96 && rune(s[i]) < 123) {
		} else {
			zS[i] = (rune(s[i]) - 32)
		}
	}
	return string(zS)
}
