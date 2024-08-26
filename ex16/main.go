package main

import "fmt"

func main() {
	sl := []int{1231, 1, 0, 2, 323, 4, 5, 256, 47, 8, 9, 245210, 8, 0, 10}
	fmt.Println(quickSort(sl))
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[len(nums)/2]

	biggerNums := []int{}
	equalNums := []int{}
	smallerNums := []int{}

	for _, x := range nums {
		if x > pivot {
			biggerNums = append(biggerNums, x)
		} else if x == pivot {
			equalNums = append(equalNums, x)
		} else {
			smallerNums = append(smallerNums, x)
		}
	}
	return append(append(quickSort(smallerNums), equalNums...), quickSort(biggerNums)...)
}
