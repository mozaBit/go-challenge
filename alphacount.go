package piscine

func AlphaCount(s string) int {
	count := 0
	for i := range s {
		if rune(s[i]) > 64 && rune(s[i]) < 91 || rune(s[i]) > 96 && rune(s[i]) < 123 {
			count++
		}
	}
	return count
}
