package piscine

func NRune(s string, n int) rune {
	z := []rune(s)
	i := len(z)
	if n > i || n < 0 {
		return rune(0)
	} else {
		if n == 0 {
			return 0
		} else {
			if n >= 1 {
				return z[n-1]
			}
		}
	}
	return rune(0)
}
