//go:build ignored

package main

import "fmt"

//CHANNELS: Place to send data.Transaction can occur when send & receive happen at same time
//CHANNELS: Have blocking ability till send & receive happen at same time. (blocking)
//CHANNELS:  c:=make(chan type) ,creation of channel
//CHANNELS :  send : c<-value, receive : <-c

//Function below doesn't block because we have send & receive on channel c
//Sender : go-routine with anon function
//Receiver : main-go-routing with fmt.Println()
func main() {
	c := make(chan int) //Create a variable of type channel int

	go func() {
		c <- 23 //Send to channel
	}()

	fmt.Println(<-c) //Receive from channel
}
