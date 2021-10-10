//go:build ignored

package main

import (
	"fmt"
)

//SELECT-STATEMENT: select { case a: ... , case b: .... case c: .... return}
//SELECT-STATEMENT: Pull multiple values of multiple channels
func main() {
	bufferLength := 7

	even := make(chan int, bufferLength)
	odd := make(chan int, bufferLength)
	quit := make(chan int, bufferLength)

	printInfo(even, odd, quit)

	go send(even, odd, quit)

	receive(even, odd, quit)

}

func printInfo(args ...chan int) {
	for _, value := range args {
		fmt.Printf("%T , %v \n", value, value)
	}
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- 0
}

func receive(e, o, q chan int) {
	for { //Forever loop until select statement hits last case and returns
		select {
		case v := <-e:
			fmt.Println("From even channel", v)

		case v := <-o:
			fmt.Println("From odd channel", v)

		case v := <-q:
			fmt.Println("From quit channel", v)
			close(q) //Close q channel
			return
		}
	}
}
