package piscine

func Swap(a *int, b *int) {
	k := 0
	k = *a
	*a = *b
	*b = k
}
