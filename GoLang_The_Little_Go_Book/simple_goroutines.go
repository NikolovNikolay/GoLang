package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker represents a worker
type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		select {
		case data := <-c:
			fmt.Printf("worker %d got %d\n", w.id, data)
		case <-time.After(time.Millisecond * 10):
			fmt.Println("Break time")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	c := make(chan int, 100)

	for i := 1; i < 5; i++ {
		w := &Worker{id: i}
		go w.process(c)
	}

	for {
		select {
		case c <- rand.Int():
		case t := <-time.After(time.Millisecond * 100):
			fmt.Println("timed out at", t)
		default:
			//this can be left empty to silently drop the data
			fmt.Println("dropped")
		}

		time.Sleep(time.Millisecond * 50)
	}
}
