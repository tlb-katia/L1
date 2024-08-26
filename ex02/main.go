package main

import (
	"fmt"
	"sync"
)

/*
Here we run several goroutines
We get the result of calculations in the main routine
To do this, we need to use a chanel
*/

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	out := make(chan int)

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

	for v := range out {
		fmt.Println(v)
	}
}

/*
A simpler approach
We only use one additional goroutine and print results from that routine
No need for channels
*/

//func main() {
//	arr := [...]int{2, 4, 6, 8, 10}
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//
//	go func() {
//		defer wg.Done()
//		for _, v := range arr {
//			fmt.Println(v * v)
//		}
//	}()
//	wg.Wait()
//}
