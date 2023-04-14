package piscine

func Atoi(s string) int {
	arrayRune := []rune(s)
	var arrayRuneN []rune
	var sf int
	if len(s) == 0 || !((rune(s[0]) == '+' && len(s) > 1) || (rune(s[0]) == '-' && len(s) > 1) || (rune(s[0]) > 47 && rune(s[0]) < 58)) {
		return 0
	} else {
		for in := 1; in < len(s); in++ {
			if !(rune(s[in]) > 47 && rune(s[in]) < 58) {
				return 0
			}
		}
	}
	if arrayRune[0] == '+' && (arrayRune[1] == '0' || arrayRune[1] == '1' || arrayRune[1] == '2' || arrayRune[1] == '3' || arrayRune[1] == '4' || arrayRune[1] == '5' || arrayRune[1] == '6' || arrayRune[1] == '7' || arrayRune[1] == '8' || arrayRune[1] == '9') {
		for i := 1; i < len(arrayRune); i++ {
			arrayRuneN = append(arrayRuneN, arrayRune[i])
		}
		for i, p := range arrayRuneN {
			a := 1
			for c := 1; c < len(arrayRuneN)-i; c++ {
				a = a * 10
			}
			if !(int(p) > 47 && int(p) < 58) {
				return 0
			}
			sf += int(p-48) * a
		}
		return sf
	} else {
		if arrayRune[0] == '-' && (arrayRune[1] == '0' || arrayRune[1] == '1' || arrayRune[1] == '2' || arrayRune[1] == '3' || arrayRune[1] == '4' || arrayRune[1] == '5' || arrayRune[1] == '6' || arrayRune[1] == '7' || arrayRune[1] == '8' || arrayRune[1] == '9') {
			for i := 1; i < len(arrayRune); i++ {
				arrayRuneN = append(arrayRuneN, arrayRune[i])
			}
			for i, p := range arrayRuneN {
				a := 1
				for c := 1; c < len(arrayRuneN)-i; c++ {
					a = a * 10
				}
				if !(int(p) > 47 && int(p) < 58) {
					return 0
				}
				sf += int(p-48) * a
			}
			sf = -sf
			return sf
		} else {
			if arrayRune[0] == '0' || arrayRune[0] == '1' || arrayRune[0] == '2' || arrayRune[0] == '3' || arrayRune[0] == '4' || arrayRune[0] == '5' || arrayRune[0] == '6' || arrayRune[0] == '7' || arrayRune[0] == '8' || arrayRune[0] == '9' {
				for i, p := range arrayRune {
					a := 1
					for c := 1; c < len(s)-i; c++ {
						a = a * 10
					}
					if !(int(p) > 47 && int(p) < 58) {
						return 0
					}
					sf += int(p-48) * a
				}
				return sf
			} else {
				return 0
			}
		}
	}
}
