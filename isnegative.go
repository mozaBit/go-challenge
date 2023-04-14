package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	var p string = "F\n\nT"
	if nb >= 0 {
		for i := 0; i <= 1; i++ {
			z01.PrintRune(rune(p[i]))
		}
	} else {
		for i := 3; i >= 2; i-- {
			z01.PrintRune(rune(p[i]))
		}
	}
}
