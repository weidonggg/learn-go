package main

import (
	"fmt"
	"time"
)

func main() {
	//
	// buffer channel
	//
	// bufferChannelDemo()
	channelClose()
}

func worker(id int, ch chan int) {

	// for {
	// 	v, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("Worker %d Received %d\n", id, v)
	// }

	for v := range ch {
		fmt.Printf("Worker %d Received %d\n", id, v)
	}
}

func createWorker(id int) chan int {
	ch := make(chan int)

	go worker(id, ch)

	return ch
}

func bufferChannelDemo() {

	c := make(chan int, 3)

	go worker(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	time.Sleep(time.Second)

}

func channelClose() {

	c := make(chan int)

	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	close(c)
	time.Sleep(time.Millisecond)

}
