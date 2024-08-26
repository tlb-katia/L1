package main

import "fmt"

func main() {
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7

	fmt.Println(binarySearch(sortedArray, target, 0, len(sortedArray)-1))
}

func binarySearch(a []int, target int, right int, left int) int {
	mid := (left + right) / 2

	if a[mid] == target {
		return mid
	} else if target < a[mid] {
		return binarySearch(a, target, left, mid)
	} else if target > a[mid] {
		return binarySearch(a, target, mid, right)
	} else {
		return -1
	}
}
