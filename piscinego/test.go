package main

import "fmt"

func main() {
	a := 4
	maFonction(&a)
	fmt.Println(a)
}

func maFonction(p *int) *int {
	*p *= *p
	return p

}
