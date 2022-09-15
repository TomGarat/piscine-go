package main

import "fmt"

func Index(s string, toFind string) int {
	for i1 := 0; i1 < len(s); i1++ {
		for i2 := 0; i2 < len(toFind); i2++ {
			if i1+i2 > len(s)-1 {
				continue
			} else if s[i1+i2] == toFind[i2] {
				return i1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
