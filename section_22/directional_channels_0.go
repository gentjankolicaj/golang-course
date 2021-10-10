//go:build ignored

package main

import "fmt"

//DIRECTIONAL-CHANNELS: With only one direction,send or receive functionality (send-channel & receive-channel)
//SEND-CHANNEL : c:=make(chan <- int,bufferLength)
//RECEIVE-CHANNEL : c:=make(<-chan int,bufferLength)
func main() {

	//SEND CHANNEL
	c := make(chan<- int, 2)

	c <- 23
	c <- 12

	//  fmt.Println(<-c) compile error
	//  fmt.Println(<-c)
	fmt.Println("-------------")
	fmt.Printf("%T \n", c)

}
