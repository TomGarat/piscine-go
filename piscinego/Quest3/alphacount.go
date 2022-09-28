package main

import "fmt"

func AlphaCount(s string) int {
	var compteur int = 0
	var alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i_string := 0; i_string < len(s)-1; i_string++ {
		for i_alpha := 0; i_alpha < len(alphabet)-1; i_alpha++ {
			if alphabet[i_alpha] == s[i_string] {
				compteur++
			}
		}
	}
	return compteur
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
