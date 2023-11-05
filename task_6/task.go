package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	done := make(chan bool)
	var wg1 sync.WaitGroup

	go func() {
		defer wg1.Done()
		fmt.Println("Goroutine 1 started")
		<-done
		fmt.Println("Goroutine 1 ended")
	}()

	wg1.Add(1)
	done <- true
	wg1.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	var wg2 sync.WaitGroup

	go func() {
		defer wg2.Done()
		fmt.Println("Goroutine 2 started")
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 2 ended by call")
		}
	}()

	wg2.Add(1)
	cancel()
	wg2.Wait()

	var wg3 sync.WaitGroup

	go func() {
		defer wg3.Done()
		fmt.Println("Goroutine 3 started")
		runtime.Goexit()
	}()

	wg3.Add(1)
	wg3.Wait()
	fmt.Println("Goroutine 3 ended")

	var wg4 sync.WaitGroup

	go func() {
		defer wg4.Done()
		fmt.Println("Goroutine 4 started")
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Goroutine 4 ended with panic call")
			}
		}()
		panic("Emulate panic to terminate a goroutine")
	}()

	wg4.Add(1)
	wg4.Wait()
}
