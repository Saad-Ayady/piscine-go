package piscine

func NRune(s string, n int) rune {
	if n < 0 || n > len(s) {
		return rune(0)
	}
	return rune(s[n-1])
}
