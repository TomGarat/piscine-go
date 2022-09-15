package main

import "fmt"

func ConcatParams(args []string) string {
	r := ""
	for , e := range args {
		r += e + "\n"
	}
	return r
}

func main() {
	test = []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}