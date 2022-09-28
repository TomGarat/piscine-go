package main

import "fmt"

//Write a function that takes a pointer to a pointer to a pointer to an int as argument and gives to this int the value of 1.

func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
}
