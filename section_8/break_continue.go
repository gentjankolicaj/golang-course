//go:build ignore

package main

import "fmt"

//Templates of For-loop control flow
//1. for init;condition;increment{}
//2. for  ;condition;increment{}
//3. for  ;condition; {}
//4. for condition{}
//5. for {} => forever until break, or continue to skip to next loop
//6. for index,value:=range of collection{}

func main() {
	counter := 10
	for true {
		fmt.Println("Counter ", counter)
		if counter > 10 {
			fmt.Println("Break condition reached ", counter)
			break
		}
		counter++
	}

	for counter >= -10 {
		counter = counter - 1
		fmt.Println("Counter ", counter)
		if counter == -5 {
			fmt.Println("Continue condition reached ", counter)
			continue
		}
	}

}
