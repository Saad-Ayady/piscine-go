package piscine

import "github.com/01-edu/z01"

func printStr(s string) {
    for _, r := range s {
        z01.PrintRune(r)
	}
}

func generateComb(current string, start, n int, result *[]string) {
    if len(current) == n {
        *result = append(*result, current)
        return
    }
    for i := start; i <= 9; i++ {
        generateComb(current+string(rune('0'+i)), i+1, n, result)
    }
}

func PrintCombN(n int) {
    if n <= 0 || n >= 10 {
        return
    }
    var result []string
    generateComb("", 0, n, &result)
    for i, comb := range result {
        if i > 0 {
            z01.PrintRune(',')
            z01.PrintRune(' ')
        }
        printStr(comb)
    }
    z01.PrintRune('\n')
}
