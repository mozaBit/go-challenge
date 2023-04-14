package piscine

func UltimateDivMod(a *int, b *int) {
	k := *a / *b
	*b = *a % *b
	*a = k
}
