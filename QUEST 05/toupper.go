package piscine

func ToUpper(s string) string {
	var w string
	for i := range s {
		if s[i] > 90 {
			v := s[i] - 32
			w = w + string(v)
		} else {
			w = w + string(s[i])
		}
	}
	return string(w)
}
