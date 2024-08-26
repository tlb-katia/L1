package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Use channel "Done"

func worker1(done chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case <-done:
			fmt.Println("Goroutine Worker1 stopped")
			wg.Done()
			return
		default:
			fmt.Println("Worker1 is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

// Use context

func worker2(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine Worker2 stopped")
			wg.Done()
			return
		default:
			fmt.Println("Worker2 is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

// use close(chan)

func worker3(done chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-done:
			fmt.Println("Goroutine Worker3 stopped")
			wg.Done()
			return
		default:
			fmt.Println("Worker3 is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

//use a timer

func worker4(wg *sync.WaitGroup) {
	stop := time.After(10 * time.Second)
	for {
		select {
		case <-stop:
			fmt.Println("Goroutine Worker4 stopped")
			wg.Done()
			return
		default:
			fmt.Println("Worker4 is working...")
			time.Sleep(1 * time.Second)
		}
	}

}

func main() {
	wg := &sync.WaitGroup{}

	done := make(chan bool)
	wg.Add(1)
	go worker1(done, wg)

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker2(ctx, wg)

	ch := make(chan struct{})
	wg.Add(1)
	go worker3(ch, wg)

	wg.Add(1)
	go worker4(wg)

	time.Sleep(3 * time.Second)
	done <- true
	cancel()
	close(ch)

	wg.Wait()
}
