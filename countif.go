package piscine

func CountIf(f func(string) bool, tab []string) int {
	p := 0
	z := make([]bool, len(tab))
	for i := 0; i < len(tab); i++ {
		z[i] = f(tab[i])
	}
	for i := 0; i < len(z); i++ {
		if z[i] {
			p += 1
		}
	}
	return p
}
