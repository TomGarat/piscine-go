package main

import "fmt"

func AppendRange(min, max int) []int {

	var rep []int

	for i := min; i < max; i++ {

		rep = append(rep, i)
	}
	return rep
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
