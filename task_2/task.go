package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	for i := 0; i < 5; i++ {
		curr_sq := i
		go func() {
			fmt.Printf("The square of %d is %d\n", arr[curr_sq], arr[curr_sq]*arr[curr_sq])
		}()
	}
	time.Sleep(10000000)
}
