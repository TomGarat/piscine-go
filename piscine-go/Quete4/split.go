package main

import "fmt"

func Split(s, sep string) []string {
	var result = []string{""}
	islice := 0
	sep_trouve := true
	for i := 0; i < len(s)-1; i++ {
		for i2 := 0; i2 < len(sep); i2++ {
			if string(sep[i2]) != string(s[i+i2]) {
				sep_trouve = false
			}
		}
		if sep_trouve {
			islice++
			result = append(result, "")
			i += len(sep) - 1
		} else {
			result[islice] += string(s[i])
		}
		sep_trouve = true
	}
	return result
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
