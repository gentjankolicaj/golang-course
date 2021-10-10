//go:build ignored

package main

import "fmt"

//DIRECTIONAL-CHANNELS: With only one direction,send or receive functionality (send-channel & receive-channel)
//SEND-CHANNEL : c:=make(chan <- int,bufferLength)
//RECEIVE-CHANNEL : c:=make(<-chan int,bufferLength)
func main() {
	bufferLength := 5

	c := make(chan int, bufferLength)
	cs := make(chan<- int, bufferLength) //Send only channel
	cr := make(<-chan int, bufferLength) //Receive only channel

	fmt.Printf("%T \n", cs)
	fmt.Printf("%T \n", cr)
	fmt.Printf("-----------------------\n")
	fmt.Printf("Conversion of channels,from general to specific \n")
	fmt.Printf("%T ,receive \n", (<-chan int)(c))
	fmt.Printf("%T ,send \n ", (chan<- int)(c))

}
