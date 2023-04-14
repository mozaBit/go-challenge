package piscine

func Map(f func(int) bool, a []int) []bool {
	z := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		z[i] = f(a[i])
	}
	return z
}
