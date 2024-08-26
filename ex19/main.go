package main

import "fmt"

func main() {
	input := "programmering" // not a typo but a Swedish word
	output := reverseString(input)
	fmt.Println(output)
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		el := runes[i]
		runes[i] = runes[n-1-i]
		runes[n-1-i] = el
	}
	return string(runes)
}
