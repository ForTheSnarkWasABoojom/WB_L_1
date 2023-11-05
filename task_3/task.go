package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	c := make(chan int)

	for i := 0; i < 5; i++ {
		curr_sq := i
		go func() {
			c <- arr[curr_sq] * arr[curr_sq]
		}()
	}
	time.Sleep(10000000)
	final_sum := 0
	for i := 0; i < 5; i++ {
		final_sum += <-c
	}
	fmt.Println(final_sum)
}
