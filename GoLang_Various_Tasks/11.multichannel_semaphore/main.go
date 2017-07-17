package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	// This is the semaphore channel. It expects two confirmations to close c and self-close.
	done := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}

		done <- true
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}

		done <- true
	}()

	go func() {

		// receive confirmations in a separate goroutine,
		// so the print below is not blocked
		<-done
		<-done
		close(c)
		close(done)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
