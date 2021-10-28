package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var c1, c2 = generatorChan(), generatorChan()
	var worker = createWorker(0)
	var timer = time.After(10 * time.Second)
	var tick = time.Tick(time.Second)

	var values []int

	for {
		var activeWorker chan int
		var activeValue int

		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]

		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tick:
			fmt.Printf("Queue len = %d\n", len(values))
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Timeout")
		case <-timer:
			fmt.Println("Bye")
			return

		}
	}

}

func generatorChan() chan int {
	ch := make(chan int)

	go func() {
		rand.Seed(time.Now().UnixNano())
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			ch <- i
			i++
		}
	}()

	return ch
}

func worker(id int, ch chan int) {

	for v := range ch {
		fmt.Printf("Worker %d Received %d\n", id, v)

		time.Sleep(2 * time.Second)
	}
}

func createWorker(id int) chan int {
	ch := make(chan int)

	go worker(id, ch)

	return ch

}
