package main

import (
	"fmt"
)

func main() {
	for n := range square(square(generate([]int{2, 3}))) {
		fmt.Println(n)
	}
}

func generate(params []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range params {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
