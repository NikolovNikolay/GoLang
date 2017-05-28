package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var num uint64

	for i := 0; i < 50; i++ {
		go atomicIncrement(&num, 1)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	time.Sleep(time.Second)

	result := atomic.LoadUint64(&num)
	fmt.Println("num:", result)
}

func atomicIncrement(num *uint64, delta uint64) {
	for {
		atomic.AddUint64(num, delta)
		time.Sleep(time.Millisecond)
	}
}
