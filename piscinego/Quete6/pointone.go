package main

import "fmt"

// Language: go
//Write a function that takes a pointer to an int as argument and gives to this int the value of 1.

func PointOne(n *int) {
	*n = 1
}

func main() {
	n := 0
	PointOne(&n)
	fmt.Println(n)
}
