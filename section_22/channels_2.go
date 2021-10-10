//go:build ignored

package main

import "fmt"

//BUFFERED-CHANNELS : c:=make(chan type,bufferLength) ,creation of buffered channel
//BUFFERED-CHANNELS : Allow value to be stored on buffer,based on buffer length

func main() {
	var bufferLength int = 5
	c := make(chan int, bufferLength) //Create a variable of type channel int

	go func() {
		c <- 23  //Send to channel
		c <- 232 //Send to channel
		c <- 21  //Send to channel
	}()

	fmt.Println(<-c) //Receive from channel
	fmt.Println(<-c) //Receive from channel
	fmt.Println(<-c) //Receive from channel
}
