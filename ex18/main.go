package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	count int
}

func main() {
	count := counter{}
	wg := sync.WaitGroup{}

	for i := 7; i < 100; i++ {
		wg.Add(1)
		go count.inc(&wg)
	}
	wg.Wait()

	fmt.Println("count:", count.count)
}

func (c *counter) inc(wg *sync.WaitGroup) {
	c.mu.Lock()
	c.count++
	wg.Done()
	c.mu.Unlock()
}
