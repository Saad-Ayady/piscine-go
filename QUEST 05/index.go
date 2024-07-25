package piscine

func Index(s string, toFind string) int {
	re := -1
	for i := 0; i < len(s); i++ {
		if string(s[i]) != string(toFind[0]) {
			continue
		} else {
			re = i
			break
		}
	}
	return re
}
