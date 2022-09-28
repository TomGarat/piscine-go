package main

import "fmt"

func IsUpper(s string) bool {

	var alphaMaj string = "ABCDEFGHIJKLMOPQRSTUVWXYZ"
	var lettreMajPresente bool = false
	for _, char := range s {
		lettreMajPresente = false
		for _, lettre := range alphaMaj {
			if char == lettre {
				lettreMajPresente = true
			}
		}
		if !lettreMajPresente {
			return false
		}
	}
	return lettreMajPresente
}

func main() {
	fmt.Println(IsUpper("HEaLLO"))
	fmt.Println(IsUpper("HELLO!"))

}
