package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// var aruine string = "abcdefghijklmnopqrstuvwxyz" + "\n"
	// for i := 0; i <= 26; i++ {
	// 	z01.PrintRune(rune(aruine[i]))
	// }
	for i := 97; i < 123; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune(rune('\n'))
	// s := "abcdefghijklemnopqrstuvwxyz"
	// for _, p := range s {
	// 	z01.PrintRune(rune(p))
	// }
}
