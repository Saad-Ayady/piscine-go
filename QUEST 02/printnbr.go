package piscine

import "github.com/01-edu/z01"

func MyFun(nb int) {
	c := '0'
	if nb == 0 {
		z01.PrintRune(c)
		return
	}
	for i := 1; i <= nb%10; i++ {
		c++
	}
	for i := -1; i >= nb%10; i-- {
		c++
	}
	if nb/10 != 0 {
		MyFun(nb / 10)
	}
	z01.PrintRune(c)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	MyFun(n)
}
