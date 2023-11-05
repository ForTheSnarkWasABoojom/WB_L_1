package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/dchest/uniuri"
)

func main() {
	dataChannel := make(chan string)
	var wg sync.WaitGroup
	numWorkers, err := strconv.Atoi(os.Args[1])
	handleError(err)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go createWorker(i, dataChannel, &wg)
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-signalChannel:
			close(dataChannel)
			wg.Wait()
			fmt.Println("All workers has stopped. Program completed")
			return
		default:
			data := uniuri.New()
			dataChannel <- data
			time.Sleep(time.Second)
		}
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func createWorker(id int, dataChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d has started\n", id)
	for data := range dataChannel {
		fmt.Printf("Worker %d received data: %s\n", id, data)
	}
	fmt.Printf("Worker %d has stopped\n", id)
}
