package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.m[key]
	return val, ok
}

func main() {
	sm := NewSafeMap()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			sm.Set(fmt.Sprintf("key%d", i), i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			val, ok := sm.Get(fmt.Sprintf("key%d", i))
			if ok {
				fmt.Printf("Got key%d: %d\n", i, val)
			} else {
				fmt.Printf("Key key%d not found\n", i)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	wg.Wait()

	fmt.Println("All goroutines finished")
}
