package main

import "fmt"

func IsLower(s string) bool {

	var alphaMin string = "abcdefghijklmnopqrstuvwxyz"
	var lettreMinPresente bool = false
	for _, char := range s {
		lettreMinPresente = false
		for _, lettre := range alphaMin {
			if char == lettre {
				lettreMinPresente = true
			}
		}
	}
	return lettreMinPresente
}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}
