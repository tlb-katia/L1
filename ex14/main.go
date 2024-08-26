package main

import "fmt"

func main() {
	ch := make(chan int)
	sl := [...]interface{}{1, "dva", true, ch}

	for _, v := range sl {
		getType(v)
	}
}

func getType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	case chan int:
		fmt.Println("chan", v)
	}
}
