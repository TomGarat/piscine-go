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

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))

}
