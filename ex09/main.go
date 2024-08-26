package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := generator()
	ch2 := multiply(ch1)

	for v := range ch2 {
		fmt.Println(v)
	}

}

func generator() chan int {
	ch := make(chan int)
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))
	for _, v := range arr {
		go func(v int) {
			ch <- v
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func multiply(in chan int) chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}

	for v := range in {
		wg.Add(1)
		go func(v int) {
			out <- v * 2
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
