//go:build ignored

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//NOTE : To check race condition ,from shell call
// $go run -race fileName.go
//NOTE : DATA-RACE or RACE-CONDITION happens when multiple go-routines try and access on variable.
//MUTEX : Is used to prevent RACE-CONDITION and this is done by calling methods Lock() & Unlock()
func main() {
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("Cpu nr: ", runtime.NumCPU())
	fmt.Println("Go-routines : ", runtime.NumGoroutine())

	var counter int = 0
	const gs int = 100

	var waitGroup sync.WaitGroup //define wait-group variable
	waitGroup.Add(gs)            //Set nr of wait-groups

	var mutex sync.Mutex //Define mutex variable

	for i := 0; i < gs; i++ {
		go func() { //Each call creates a new go-routine (thread)
			mutex.Lock() //Lock execution of lines below to one go-routine at a time

			v := counter      //Read value
			runtime.Gosched() //Yields processor to allow other go-routines to run
			v++
			counter = v //Write value

			mutex.Unlock()   //Unlocks the locked lines to be executed by other go-routines
			waitGroup.Done() //Set goroutine as done in wait-group
		}() // Call anon function with ()
	}

	waitGroup.Wait() //Wait till all goroutines have finished executing

	//NOTE: Since in all go-routines is used Mutex for lock & unlock of variable read-write
	//value of counter is going to be 100
	fmt.Println("Go-routines : ", runtime.NumGoroutine())
	fmt.Println("Count value : ", counter)

}
