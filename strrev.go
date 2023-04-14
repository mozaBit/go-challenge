package piscine

func StrRev(s string) string {
	ruS := []rune(s)
	var r string
	for i := range ruS {
		r = string(ruS[i]) + r
	}
	return r
}
