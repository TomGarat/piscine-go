package main

import "fmt"

func IsPrime(nb int) bool {
	var est_prime bool = true
	if nb == 1 {
		est_prime = false
	} else {
		for i := 2; i < nb-1; i++ {
			if nb%i == 0 {
				est_prime = false
				break
			}
		}
	}
	return est_prime
}

func Map(f func(int) bool, a []int) []bool {
	var slice_return []bool
	for _, e := range a {
		slice_return = append(slice_return, f(e))
	}
	return slice_return
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}
