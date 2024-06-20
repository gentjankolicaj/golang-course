//go:build ignored

package main

import "fmt"

//NOTE: When working with channels ,we need at least 2 go-routines
func main() {
	c := make(chan int)
	maxIndex := 1400

	go func() {
		for i := 0; i < maxIndex; i++ {
			c <- i
		}
		close(c)
	}()

	for j := 0; j < maxIndex; j++ {
		v, ok := <-c //receive value from channel c
		fmt.Println("value - ", v, " Ok - ", ok)
	}
}
