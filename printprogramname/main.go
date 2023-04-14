package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	z := []rune(os.Args[0])
	for i := 2; i < len(z); i++ {
		z01.PrintRune(z[i])
	}
	z01.PrintRune('\n')
}
