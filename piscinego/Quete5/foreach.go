package main

import "fmt"

func printnbr(nb int) int {
	return nb
}

func ForEach(f func(int) int, a []int) {
	for _, e := range a {
		fmt.Print(f(e))
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(printnbr, a)
}
