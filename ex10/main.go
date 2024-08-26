package main

import "fmt"

func main() {
	temp := [...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float32)

	for _, t := range temp {
		key := int(t/10) * 10
		groups[key] = append(groups[key], t)
	}

	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, v)
	}

}
