//go:build ignored

package main

import (
	"fmt"
	"runtime"
)

//NOTE: func init runs before main .Usually we use for initialization
func init() {
	fmt.Println("Func init() invoked()...")
}

func main() {
	printHardwareInfo()

	//Concurrent design
	//We lunch go-routines
	go foo()
	bar()
}

func foo() {
	for i := 0; i < 17; i++ {
		fmt.Println("foo-", i)
	}
}

func bar() {
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
