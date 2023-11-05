package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value++
}

func main() {
	counter := &Counter{}

	var wg sync.WaitGroup
	numGoroutines := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Final counter: %d\n", counter.Value)
}
