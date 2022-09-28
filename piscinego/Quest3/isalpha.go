package main

import "fmt"

func IsAlpha(s string) bool {

	if s == "" {
		return true
	}
	var alphaAlpha string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var lettreAlphaPresente bool = false
	for _, char := range s {
		lettreAlphaPresente = false
		for _, lettre := range alphaAlpha {
			if char == lettre {
				lettreAlphaPresente = true
			}
		}
	}
	return lettreAlphaPresente
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
	fmt.Println(IsAlpha(""))

}
