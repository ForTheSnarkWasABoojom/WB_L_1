package main

import (
	"fmt"
	"time"
)

func SleepForNextNSeconds(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("Sleep started.")
	SleepForNextNSeconds(10)
	fmt.Println("Sleep ended.")
}
