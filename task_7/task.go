package main

import (
	"fmt"
	"sync"

	"github.com/dchest/uniuri"
)

func main() {
	var mu sync.Mutex
	data := make(map[int]string)
	numGoroutines := 5
	var wg sync.WaitGroup

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			data[id] = uniuri.New()

			fmt.Printf("Goroutine %d has stopped\n", id)
		}(i)
	}

	wg.Wait()

	fmt.Println("Map data:")
	for key, value := range data {
		fmt.Printf("Key - %d | Value - %s\n", key, value)
	}
}
