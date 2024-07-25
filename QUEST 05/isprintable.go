package piscine

func IsPrintable(s string) bool {
	for i := range s {
		if s[i] <= 31 || s[i] == 127 {
			return false
		}
	}
	return true
}
