//go:build ignored

package main

import (
	"fmt"
	"runtime"
)

func main() {
	printSoftwareInfo()
	printHardwareInfo()
}

func printSoftwareInfo() {
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("COMPILER ", runtime.Compiler)
}

func printHardwareInfo() {
	fmt.Println("ARCH ", runtime.GOARCH)
	fmt.Println("CPU nr: ", runtime.NumCPU())

}
