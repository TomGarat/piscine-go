package main

import "fmt"

func ToUpper(s string) string {
	var sbyte = []byte(s)
	for i := 0; i < len(s); i++ {
		if sbyte[i] > 96 && sbyte[i] < 123 {
			sbyte[i] = sbyte[i] - 32
		}
	}
	return string(sbyte)
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
