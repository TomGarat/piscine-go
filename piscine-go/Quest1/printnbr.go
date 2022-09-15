package main

import "github.com/01-edu/z01"

func main() {
	PrintNbr(-1230123)
}

func PrintNbr(number int) {
	if number < 0 {
		z01.PrintRune('-')
		number = -number
	}
	if number > 9 {
		PrintNbr(number / 10)
	}
	z01.PrintRune(rune(number%10 + 48))
}
