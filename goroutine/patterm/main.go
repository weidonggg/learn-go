package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(s string) chan string {
	c := make(chan string)

	go func() {
		i := 0

		rand.Seed(time.Now().UnixNano())
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("Service: %s, Message: %d", s, i)
			i++
		}
	}()

	return c

}

func fanIn(c1, c2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-c1
		}

	}()

	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
			}
		}
	}()

	return c
}

func fanInBySelectSlice(chs ...chan string) chan string {
	c := make(chan string)

	for _, ch := range chs {
		go func(m chan string) {
			for {
				c <- <-m
			}
		}(ch)
	}
	return c
}

func main() {

	m1 := msgGen("Service1")
	m2 := msgGen("Service2")

	m := fanInBySelectSlice(m1, m2)
	for {
		fmt.Println(<-m)
	}

}
