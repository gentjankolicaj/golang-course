//go:build ignored

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("Cpu nr: ", runtime.NumCPU())
	fmt.Println("Go-routines : ", runtime.NumGoroutine())

	var counter int = 0
	const gs int = 100

	var waitGroup sync.WaitGroup //define wait-group variable
	waitGroup.Add(gs)            //Set nr of wait-groups

	for i := 0; i < gs; i++ {
		go func() { //Each call creates a new go-routine (thread)
			v := counter //Read value

			time.Sleep(time.Second) //Go-routine sleep for 7 seconds
			runtime.Gosched()       //Yields processor to allow other go-routines to run

			v++
			counter = v      //Write value
			waitGroup.Done() //Set goroutine as done in wait-group
		}() // Call anon function with ()
	}

	waitGroup.Wait() //Wait till all goroutines have finished executing

	//NOTE: Since all go-routines are in race-condition,value of counter is going to be 1
	fmt.Println("Go-routines : ", runtime.NumGoroutine())
	fmt.Println("Count value : ", counter)

}
