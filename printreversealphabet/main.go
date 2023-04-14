package main

import "github.com/01-edu/z01"

func main() {
	var aruine string = "\nabcdefghijklmnopqrstuvwxyz"
	for i := 26; i >= 0; i-- {
		z01.PrintRune(rune(aruine[i]))
	}
}
