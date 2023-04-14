package piscine

func Any(f func(string) bool, a []string) bool {
	const passed = true
	const didntpassed = false
	z := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		z[i] = f(a[i])
	}
	for i := 0; i < len(z); i++ {
		if z[i] {
			return passed
		}
	}
	return didntpassed
}
