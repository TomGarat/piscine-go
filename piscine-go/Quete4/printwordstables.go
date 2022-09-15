package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	var result = []string{""}
	islice := 0
	for _, e := range s {
		if string(e) == " " {
			islice++
			result = append(result, "")
		} else {
			result[islice] += string(e)
		}
	}
	return result
}

func PrintWordsTables(a []string) {
	for _, e := range a {
		fmt.Println(e)
	}
}

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}
