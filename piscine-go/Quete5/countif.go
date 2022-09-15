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

func CountIf(f func(string) bool, tab []string) int {
	count_nb := 0
	for _, e := range tab {
		if f(e) {
			count_nb++
		}
	}
	return count_nb
}

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
