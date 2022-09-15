package main

import (
	"github.com/01-edu/z01"
)

func PrintComb() {

	var l bool = true

	for a := -1; a < 9; a++ {
		for b := -1; b < 9; b++ {
			for c := -1; c < 9; c++ {
				for d := 0; d < 9; d++ {

					if a+b < c+d {

						if !l {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						z01.PrintRune(rune(a + 49))
						z01.PrintRune(rune(b + 49))
						z01.PrintRune(' ')
						z01.PrintRune(rune(c + 49))
						z01.PrintRune(rune(d + 49))
						l = false
					}
				}
			}
		}
	}
}

func main() {
	PrintComb()
}
