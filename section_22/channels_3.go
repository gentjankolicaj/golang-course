//go:build ignored

package main

import "fmt"

//UNSUCCESSFUL-BUFFERED-CHANNELS : c:=make(chan type,bufferLength) ,creation of buffered channel
//UNSUCCESSFUL-BUFFERED-CHANNELS : Allow value to be stored on buffer,based on buffer length

func main() {
	var bufferLength int = 2
	c := make(chan int, bufferLength) //Create a variable of type channel int

	c <- 23  //Send to channel
	c <- 232 //Send to channel
	c <- 21  //Send to channel => blocking because buffer length is 2

	fmt.Println(<-c) //Receive from channel
	fmt.Println(<-c) //Receive from channel
	fmt.Println(<-c) //Receive from channel
}
