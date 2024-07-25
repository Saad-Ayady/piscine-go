package piscine

func AlphaCount(s string) int {
	c := 0
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			c++
		}
	}
	return c
}
