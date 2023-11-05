package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

var arrLen = 10

func main() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < arrLen; i++ {
		arr[i] = rand.Intn(100)
	}

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	for i := 0; i < arrLen; i++ {
		ch1 <- arr[i]
		ch2 <- <-ch1 * 2
		fmt.Println(<-ch2)
	}
}
