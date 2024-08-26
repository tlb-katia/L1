package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(deleteElement(nums, 4))
}

func deleteElement(s []int, ind uint) []int {
	return append(s[:ind], s[ind+1:]...)
}
