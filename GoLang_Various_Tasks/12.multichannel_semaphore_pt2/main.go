package main

import (
	"fmt"
)

func main() {
	n := 10

	c := make(chan int)
	done := make(chan bool)

	// Single goroutine is publishing to a channel
	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	// Multiple goroutines are reading from a channel
	for i := 0; i < n; i++ {
		go func(_i int) {
			for n := range c {
				out := fmt.Sprintf("%d says: %d", _i, n)
				fmt.Println(out)
			}
			done <- true
		}(i)
	}

	// Need to wait all of the goroutines to ping the done channel,
	// and then finish the program. This happens on the main thread,
	// otherwise there will be no output.
	for i := 0; i < n; i++ {
		<-done
	}
}
