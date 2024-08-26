package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sendValues(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	value := 0
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		case ch <- value:
			wg.Add(1)
			value++
		}
	}
}

func readValues(ch <-chan int, wg *sync.WaitGroup) {
	for range ch {
		wg.Done()
		fmt.Println("readValues", <-ch)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	N := 1 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), N)
	defer cancel()
	wg.Add(1)

	ch := make(chan int)

	go sendValues(ctx, ch, wg)

	go readValues(ch, wg)

	go func() {
		wg.Wait()
		<-ctx.Done()
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Program exiting")
}
