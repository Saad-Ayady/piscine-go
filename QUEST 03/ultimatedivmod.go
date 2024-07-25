package piscine

func UltimateDivMod(a *int, b *int) {
	i := *a
	*a = i / *b
	*b = i % *b
}
