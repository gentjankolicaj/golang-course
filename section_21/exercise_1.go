//go:build ignored

package main

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("Before main : init() invoked")
}

func main() {
	//Define local level scope waitGroup
	var waitGroup sync.WaitGroup

	//Add 2 goroutines in waitGroup
	waitGroup.Add(2)

	//Launch 2 go-routines
	go routine0(&waitGroup)
	go routine1(&waitGroup)

	//Call method wait to make main-routine wait 2 newly created routines
	waitGroup.Wait()
	fmt.Println("2 New goroutines finished.")

}

func routine0(wg *sync.WaitGroup) {
	fmt.Println("Invoked func routine0()")
	wg.Done() //Remove routine from wait group
}

func routine1(wg *sync.WaitGroup) {
	fmt.Println("Invoked func routine1()")
	wg.Done() //Remove routine from wait group
}
