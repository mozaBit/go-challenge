package main

import "github.com/01-edu/z01"

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

type point struct {
	x int
	y int
}

func main() {
	sx := "x = "
	sy := "y = "
	es := ", "
	points := &point{}

	setPoint(points)
	printStr(sx)
	PrintNbr2(points.x)
	printStr(es)
	printStr(sy)
	PrintNbr2(points.y)
	z01.PrintRune('\n')
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(rune(r))
	}
}

func PrintNbr2(n int) {
	t := 1
	if n < 0 {
		t = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		f := (n / 10) * t
		if f != 0 {
			PrintNbr2(f)
		}
		k := (n % 10 * t) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}
