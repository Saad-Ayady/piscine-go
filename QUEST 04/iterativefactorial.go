package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	r := 1
	for i := 1; i <= nb; i++ {
		r *= i
		if r <= 0 {
			return 0
		}
	}
	return r
}
