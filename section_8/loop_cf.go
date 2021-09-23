//go:build ignore

package main

import (
	"fmt"
)

//Templates of For-loop control flow
//1. for init;condition;increment{}
//2. for  ;condition;increment{}
//3. for  ;condition; {}
//4. for condition{}
//5. for {} => forever until break, or continue to skip to next loop
//6. for index,value:=range of collection{}

func main() {
	MultipleLoops()

	RangeExample()
}

func MultipleLoops() {
	var counter int = 0
	for init := 0; init <= 10; init++ {
		fmt.Println("For-loop template 1, index ", init)
		if init >= 3 {
			counter = init
			break
		}
	}

	for ; true; counter++ {
		fmt.Println("For-loop template 2 ,index ", counter)
		if counter > 5 {
			break
		}
	}

	for true {
		fmt.Println("For-loop template 3 ,index ", counter)
		if counter > 7 {
			break
		}
		counter++
	}

	var i int = 0
	var max int = 100
	for i <= max {
		fmt.Println("For-loop template 4 ,index ", i)
		i++
		if i >= 3 {
			break
		}
	}

}

func RangeExample() {
	var array [10]int
	for index, value := range array {
		fmt.Println("RangeExample : ", index, value)
	}
}
