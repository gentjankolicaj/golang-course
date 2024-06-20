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
	quit := make(chan bool)

	printInfo(even, odd, quit)

	go send(even, odd, quit)

	receive(even, odd, quit)

}

func printInfo(args ...interface{}) {
	for _, value := range args {
		fmt.Printf("%T , %v \n", value, value)
	}
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for { //Forever loop until select statement hits last case and returns
		select {
		case v := <-even:
			fmt.Println("From even channel", v)

		case v := <-odd:
			fmt.Println("From odd channel", v)

		case i, ok := <-quit:
			if !ok {
				fmt.Println("From comma-ok ", i, ok)
				return
			} else {
				fmt.Println("From comma-ok ", i, ok)
			}
		}
	}
}
