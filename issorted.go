package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	const sorted = true
	const notsorted = false
	p := 0
	d := 0
	z := make([]int, len(a)-1)
	for i := 0; i < len(a)-1; i++ {
		z[i] = f(a[i+1], a[i])
	}
	for i := 0; i < len(z); i++ {
		if z[i] > 0 {
			p += 1
			d += 1
		} else {
			if z[i] == 0 {
				d += 0
			}
		}
	}
	if len(z) == p || d == 0 {
		return sorted
	} else {
		return notsorted
	}
}
