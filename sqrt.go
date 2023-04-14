package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for i := 2; i <= nb; i++ {
		if IMterativePower(i, 2) == nb {
			return i
		}
	}
	return 0
}

func IMterativePower(nb int, power int) int {
	result := 1
	if power < 0 || power*nb > 2*63 {
		return 0
	} else {
		for i := 1; i < power+1; i++ {
			result = result * nb
		}
	}
	return result
}
