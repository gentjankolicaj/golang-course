//go:build ignored

package main

import (
	"fmt"
	"sync"
)

func main() {
	c0 := make(chan int)
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 10000; i++ {
			c0 <- i
		}
		close(c0) //Close channel c0 after finishing populate
	}()

	go fanOut(c0, c1, c2)

	for v := range c2 {
		fmt.Println("c2 channel value -", v)
	}

}

func fanOut(c0, c1, c2 chan int) {
	var wg sync.WaitGroup
	wg.Add(2) //because of 2  go-routines

	//Receive from c0 and send to c1
	go func() {
		for value := range c0 { //Range over c0 to get all values of c0
			c1 <- value
		}
		close(c1)
		wg.Done()
	}()

	//Receive from c0 and send to c2
	go func() {
		for value := range c0 { //Range over c0 to get all values of c0
			c2 <- value
		}
		close(c2)
		wg.Done()
	}()

	wg.Wait() //Make parent go-routine to wait for child go-routines to finish executing,
	//It is a blocking call

}
