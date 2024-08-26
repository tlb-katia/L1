package main

import "fmt"

func main() {
	a := 174
	b := 13

	a += b
	b = a - b
	a -= b

	fmt.Println(a, b)
}
