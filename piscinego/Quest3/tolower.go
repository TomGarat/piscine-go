package main

import "fmt"

func ToLower(s string) string {
	var sbyte = []byte(s)
	for i := 0; i < len(s); i++ {
		if sbyte[i] > 64 && sbyte[i] < 91 {
			sbyte[i] = sbyte[i] + 32
		}
	}
	return string(sbyte)
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
