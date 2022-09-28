package main

import "fmt"

func IsPrime(nb int) bool {

	var estPrime bool = true

	if nb == 1 {
		estPrime = false
	} else {
		for i := 2; i < nb-1; i++ {
			if nb%i == 0 {
				estPrime = false
			}
		}
	}
	return estPrime

}

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}
