package piscine

func StrLen(s string) int {
	rS := []rune(s)
	i := 0
	for i = range rS {
		i++
	}
	return i
}
