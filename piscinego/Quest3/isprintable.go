package main

import "fmt"

func IsPrintable(s string) bool {

	var sbytes = []byte(s)
	for i := 0; i < len(s)-1; i++ {
		if sbytes[i] < 32 || sbytes[i] > 127 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}
