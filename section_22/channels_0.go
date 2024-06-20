//go:build ignored

package main

import "fmt"

//CHANNELS: Place to send data.Transaction can occur when send & receive happen at same time
//CHANNELS: Have blocking ability till send & receive happen at same time. (blocking)
//CHANNELS:  c:=make(chan type) ,creation of channel

//Func main below blocks because channel c ,has a sender but not a receiver
//Based on the ability of channels ,each transaction must have send & receive happen same time
func main() {
	c := make(chan int) //Create a variable of type channel int

	c <- 23 //Send to channel,=> blocks main-goroutine because of channel ability

	fmt.Println(<-c)
}
