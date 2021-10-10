//go:build ignored

package main

import "fmt"

func main() {
	bufferLength := 100
	c := make(chan int, bufferLength)

	go sendToChannel(c, bufferLength)

	receiveFromChannel(c)

}

func sendToChannel(c chan<- int, bufferLength int) {
	for i := 0; i < bufferLength; i++ {
		c <- i
	}
	close(c)
}

func receiveFromChannel(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
