//go:build ignored

package main

import (
	"fmt"
	"sync"
)

//DIRECTIONAL-CHANNELS: With only one direction,send or receive functionality (send-channel & receive-channel)
//SEND-CHANNEL : c:=make(chan <- int,bufferLength)
//RECEIVE-CHANNEL : c:=make(<-chan int,bufferLength)
//NOTE: In channels we can go from general to specific not the other way around.
func main() {
	bufferLength := 5
	c := make(chan int, bufferLength)

	var wg sync.WaitGroup
	wg.Add(2) //Because we have 2 go-routines

	go send(c, &wg)
	go receive(c, &wg)

	wg.Wait() //Wait for go-routines to finish

	fmt.Println("About to exit")

}

func send(cs chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		cs <- i
	}
	wg.Done()
}

func receive(cr <-chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, " : ", <-cr)
	}
	wg.Done()
}
