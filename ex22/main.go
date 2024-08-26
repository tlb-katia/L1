package main

import (
	"fmt"
	"math/big"
)

func add(a, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Add(a, b)
}

func subtract(a, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Sub(a, b)
}

func multiply(a, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Mul(a, b)
}

func divide(a, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Div(a, b)
}

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1048577", 10) // 1048577 = 2^20 + 1
	b.SetString("2097153", 10) // 2097153 = 2^21 + 1

	fmt.Println(add(a, b))
	fmt.Println(subtract(a, b))
	fmt.Println(multiply(a, b))
	fmt.Println(divide(b, a))
}
