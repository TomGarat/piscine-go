package main

import "fmt"

func FindNextPrime(nb int) int {
	var est_prime bool = true
	for i := nb; i < nb+100; i++ {
		est_prime = true
		for i2 := 2; i2 < i-1; i2++ {
			if i%i2 == 0 {
				est_prime = false
			}
		}
		if est_prime {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}
