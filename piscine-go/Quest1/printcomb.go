package main

import (
	"github.com/01-edu/z01"
)

func PrintComb() {

	var l bool = true

	for c := '0'; c <= '9'; c++ {
		for d := '0'; d <= '9'; d++ {
			for u := '0'; u <= '9'; u++ {
				if (c < d) && (d < u) {
					if !l {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(c)
					z01.PrintRune(d)
					z01.PrintRune(u)
					l = false
				}
			}
		}
	}
}

func main() {
	PrintComb()
}
