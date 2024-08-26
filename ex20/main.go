package main

import (
	"fmt"
	"strings"
)

func main() {
	sent := "snow dog sun"
	fmt.Println(reverseWords(sent))
}

func reverseWords(s string) string {
	splitStr := strings.Split(s, " ")
	l := len(splitStr)

	for i := 0; i < l/2; i++ {
		el := splitStr[i]
		splitStr[i] = splitStr[l-i-1]
		splitStr[l-i-1] = el
	}

	return strings.Join(splitStr, " ")
}
