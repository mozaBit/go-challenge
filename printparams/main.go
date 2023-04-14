package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		z := os.Args[i]
		zR := []rune(z)
		for _, p := range zR {
			z01.PrintRune(p)
		}
		z01.PrintRune('\n')
	}
}
