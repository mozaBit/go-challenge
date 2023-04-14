package piscine

func Compare(a, b string) int {
	/*	var c int
		zA := []rune(a)
		zB := []rune(b)
		for i := 0; i == len(zA)-1 && i == len(zB)-1; i++ {
			if zA[i] == zB[i] {
				c := 0
				return c
			} else {
				if zA[i] > zB[i] {
					c := 1
					return c
				} else {
					if zA[i] < zB[i] {
						c := -1
						return c
					}
				}
			}
		}

		return c */
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
