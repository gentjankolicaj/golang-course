//go:build ignored

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//NOTE:Package sync has synchronization vars
//Synchronization variables are used to orchestrate execution in multiple goroutines
//This variable since is declared on package has package scope
var wg sync.WaitGroup

func main() {
	printHardwareInfo()

	//Concurrent design and we lunch go-routines

	//Using synchronization primitive,to orchestrate routines
	wg.Add(1) //Done method is opposite of Add
	go foo()

	bar()
	wg.Wait() //Wait till all things added are done
}

func foo() {
	for i := 0; i < 17; i++ {
		fmt.Println("foo-", i)
	}
	wg.Done() //Invoke done method on wait group
}

func bar() {
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Println("bar-", i)
	}
}

func printHardwareInfo() {
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("ARCH ", runtime.GOARCH)
	fmt.Println("CPU ", runtime.NumCPU())
	fmt.Println("GO-routines ", runtime.NumGoroutine())
}
