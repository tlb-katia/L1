package main

import (
	"fmt"
	"sync"
)

const arrSize = 5

func main() {
	var sum int
	arr := [arrSize]int{2, 4, 6, 8, 10}
	out1 := square(arr)

	for v := range out1 {
		sum += v
	}

	fmt.Println(sum)
}

func square(arr [arrSize]int) chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}

	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			out <- v * v
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
