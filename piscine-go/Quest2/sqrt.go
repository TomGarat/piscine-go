package main

import "fmt"

func Sqrt(nb int) int {

	for i := 2; i < nb; i++ {
		nb = nb / i

		if nb < 2 {
			return 0
		}

	}
	return nb

}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
