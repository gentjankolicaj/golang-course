//go:build ignored

package main

import "fmt"

//DIRECTIONAL-CHANNELS: With only one direction,send or receive functionality (send-channel & receive-channel)
//SEND-CHANNEL : c:=make(chan <- int,bufferLength)
//RECEIVE-CHANNEL : c:=make(<-chan int,bufferLength)
func main() {

	//RECEIVE-CHANNEL
	c := make(<-chan int, 2)

	// c<-23 Compile error because can't sent to receive-only channel
	//  c<-12

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-------------")
	fmt.Printf("%T \n", c)

}
