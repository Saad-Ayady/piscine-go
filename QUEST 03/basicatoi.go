package piscine

func BasicAtoi(s string) int {
	result := 0
	multiplier := 1
	for i := len(s) - 1; i >= 0; i-- {
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		result += digit * multiplier
		multiplier *= 10
	}
	return result
}
