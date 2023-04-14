package piscine

func Index(s string, d string) int {
	n := len(d)
	if n == 0 {
		return -1
	}
	if n > len(s) {
		return -1
	}
	if n == len(s) {
		if d == s {
			return 0
		}
	}
	if n < len(s) {
		for i := range s {
			if d[0] == s[i] {
				return i
			}
		}
	}

	return -1
}
