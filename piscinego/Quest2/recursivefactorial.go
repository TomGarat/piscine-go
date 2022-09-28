package main

import "fmt"

func recursivefactorial(nb int) int {

	if nb > 1 {
		return nb * recursivefactorial(nb-1)
	}
	return 1
}

func main() {
	arg := 4
	fmt.Println(recursivefactorial(arg))
}
