//go:build ignored

package main

import (
	"fmt"
	"sync"
)

//FAN-IN: Concurrent design pattern

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

func send(even, odd chan int) {
	for i := 0; i < 30; i++ {
		even <- i
		for j := 2; j < 8; j = j + 2 {
			odd <- j
		}
	}
	close(even)
	close(odd)
}

func receive(even, odd, fanin chan int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
