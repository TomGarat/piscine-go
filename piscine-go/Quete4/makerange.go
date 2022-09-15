package main

import "fmt"

func MakeRange(min, max int) []int {

	iv := min

	if max > min {
		var rep []int = make([]int, max-min)
		for i := 0; iv < max; i++ {
			iv++
			rep[i] = iv - 1
		}
		return rep
	} else {
		return []int{}
	}

}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
