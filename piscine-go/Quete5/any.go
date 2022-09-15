package main

import "fmt"

func IsNumeric(s string) bool {

	var sbytes = []byte(s)
	for i := 0; i < len(s)-1; i++ {
		if sbytes[i] < 48 || sbytes[i] > 57 {
			return false
		}
	}
	return true
}

func Any(f func(string) bool, a []string) bool {
	for _, e := range a {
		if f(e) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Any(IsNumeric, []string{"Hello", "how", "are", "you"}))
	fmt.Println(Any(IsNumeric, []string{"This", "is", "4", "you"}))
}
