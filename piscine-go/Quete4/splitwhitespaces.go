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

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
