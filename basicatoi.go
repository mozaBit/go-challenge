package piscine

func BasicAtoi(s string) int {
	arrayRune := []rune(s)
	var sf int
	for i, p := range arrayRune {
		a := 1
		for c := 1; c < len(s)-i; c++ {
			a = a * 10
		}
		sf += int(p-48) * a
	}
	return sf
}
