package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	m := make(map[int]int)

	for _, v := range set1 {
		m[v] += 1
	}

	for _, v := range set2 {
		m[v] += 1
	}

	for k, v := range m {
		if v > 1 {
			fmt.Println(k)
		}
	}
}
