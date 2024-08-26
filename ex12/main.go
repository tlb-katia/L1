package main

import "fmt"

func main() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]int)

	for _, v := range set {
		m[v]++
	}

	for k := range m {
		fmt.Println(k)
	}
}
