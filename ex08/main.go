package main

import "fmt"

func main() {
	fmt.Println(shiftBit(4, 17))
}

func shiftBit(n uint, num int) int {
	if n == 0 {
		return num
	}
	return num | (1 << (n - 1))
}
