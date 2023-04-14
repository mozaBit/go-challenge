package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		z := os.Args[i]
		zR := []rune(z)
		for j := 0; j < len(zR); j++ {
			z01.PrintRune(zR[j])
		}
		z01.PrintRune('\n')
	}
}
