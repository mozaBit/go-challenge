package piscine

func BasicAtoi2(s string) int {
	arrayRune := []rune(s)
	var sf int
	for i, p := range arrayRune {
		a := 1
		for c := 1; c < len(s)-i; c++ {
			a = a * 10
		}
		if !(int(p) > 47 && int(p) < 58) {
			return 0
		}
		sf += int(p-48) * a
	}
	return sf
}
