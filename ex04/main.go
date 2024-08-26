package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	res := geneartor(ctx)

	for i := 0; i < 10; i++ {
		go worker(ctx, res)
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT)
	<-signalCh
	cancel()
}

func geneartor(ctx context.Context) chan int {
	out := make(chan int)
	var i int
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				out <- i
				i++

			}
		}

	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func worker(ctx context.Context, in chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for v := range in {
				fmt.Println(v)
			}
		}
	}
}
