//go:build ignored

package main

import (
	"fmt"
)

//DIRECTIONAL-CHANNELS: With only one direction,send or receive functionality (send-channel & receive-channel)
//SEND-CHANNEL : c:=make(chan <- int,bufferLength)
//RECEIVE-CHANNEL : c:=make(<-chan int,bufferLength)
//NOTE: In channels we can go from general to specific not the other way around.
//RANGE: Will keep looping to over channel till is closed.We have to be careful because if channel is not closed it loops-forever
//CLOSE-CHANNEL: close(c) ,we close by using builtin function close(type)
func main() {
	bufferLength := 5
	c := make(chan int, bufferLength)

	go send(c)

	for v := range c {
		fmt.Printf("%T , %v \n", v, v)
	}

	fmt.Println("About to exit")

}

func send(cs chan<- int) {
	for i := 0; i < 1000; i++ {
		cs <- i
	}
	close(cs) //Close channel
}
