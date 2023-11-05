package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func AbortProgram(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("Time is up. Program has been stopped.")
	os.Exit(1)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Enter duration of program's work in seconds:")
	var i int
	_, err := fmt.Scanf("%d", &i)
	CheckError(err)

	go AbortProgram(i)
	c := make(chan int)

	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			c <- rand.Intn(100)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	select {}
}
