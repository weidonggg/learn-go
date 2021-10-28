package main

import (
	"fmt"
	"sync"
)

func main() {
	chanDemo()

}

type worker struct {
	in chan int
	// done chan bool
	// wg *sync.WaitGroup
	done func()
}

func chanDemo() {
	var wg sync.WaitGroup
	channels := make([]worker, 10)

	for i := range channels {
		channels[i] = createWorker(i, &wg)
	}

	for _, w := range channels {

		wg.Add(2)
		go func(w worker) {
			for i := 0; i < 2; i++ {
				// wg.Add(1)

				// 在此情况下，chan 接收方不能同步发送 done，而是应该通过新的协程发送 done，等待接收方接收。
				w.in <- i

			}
		}(w)
		// wg.Add(10)
		// producer(worker)

		// <-worker.done

	}

	wg.Wait()

	// for _, worker := range channels {
	// 	<-worker.done
	// }

}

func createWorker(id int, wg *sync.WaitGroup) worker {

	worker := worker{
		in: make(chan int),
		// done: make(chan bool),
		// wg: wg,
		done: func() {
			wg.Done()
		},
	}

	go doWorker(id, worker)

	return worker
}

func doWorker(id int, w worker) {

	for v := range w.in {
		fmt.Printf("Worker %d Received %d\n", id, v)
		// done chan 必须在该循环内，因为调用方的每个 worker 都在等待 done chan 接收数据，有接收方就必须有发送方。
		// go func() { done <- true }()
		// done <- true
		// go func() { wg.Done() }()
		w.done()

	}

}

func producer(w worker) {

	go func() {
		for i := 0; i < 10; i++ {
			// w.wg.Add(1)
			w.in <- i
		}
	}()
}
