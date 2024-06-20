//go:build ignored

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//NOTE : To check race condition ,from shell call
// $go run -race fileName.go
//NOTE : DATA-RACE or RACE-CONDITION happens when multiple go-routines try and access on variable.
//ATOMIC : Used to prevent race-condition
func main() {
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("Cpu nr: ", runtime.NumCPU())
	fmt.Println("Go-routines : ", runtime.NumGoroutine())

	var counter int64 = 0
	const gs int = 100

	var waitGroup sync.WaitGroup //define wait-group variable
	waitGroup.Add(gs)            //Set nr of wait-groups

	for i := 0; i < gs; i++ {
		go func() { //Each call creates a new go-routine (thread)
			atomic.AddInt64(&counter, 1)                             //Write to counter in atomic way
			fmt.Println(i, "- counter ", atomic.LoadInt64(&counter)) //Read from counter in atomic way

			runtime.Gosched() //Yields processor to allow other go-routines to run

			waitGroup.Done() //Set goroutine as done in wait-group
		}() // Call anon function with ()
	}

	waitGroup.Wait() //Wait till all goroutines have finished executing

	//NOTE: Since in all go-routines is used atomic package with its functions
	//All read-write operations are atomic & counter value is 100
	fmt.Println("Go-routines : ", runtime.NumGoroutine())
	fmt.Println("Count value : ", counter)

}
