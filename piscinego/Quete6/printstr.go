package main

func PrintStr(s string) {
	for _, r := range s {
		print(string(r))
	}
}

func main() {
	PrintStr("Hello World!")
}
