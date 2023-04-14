package piscine

func LastRune(s string) rune {
	z := []rune(s)
	i := len(z)
	return z[i-1]
}
