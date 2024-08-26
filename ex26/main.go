package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abCdefAaf"
	str2 := "ABdelm"
	str3 := "qwerty"

	fmt.Println(checkUnique(str1))
	fmt.Println(checkUnique(str2))
	fmt.Println(checkUnique(str3))

}

func checkUnique(str string) bool {
	newStr := strings.ToLower(str)
	m := make(map[rune]int)

	for _, v := range newStr {
		m[v]++
	}

	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}
