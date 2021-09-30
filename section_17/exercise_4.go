//go:build ignore

package main

import (
	"fmt"
)

//Pointer : value pointing to some memory address
//Pointer : We get pointer value by using operator &,it gives memory address
//Pointer is different from type that is pointing to
//Operator & : Gives address where value is stored
//Operator * : Gives the value that address is holding (also called de-referencing)
//Everything in GOLANG is pass by value

//Pointer : Good to be used when we have a large group of data, and we don't want to pass it around
//Pointer : To change a value at certain location
func main() {
	var ptrSlice []*int = make([]*int, 10, 10) //Pointer slice , hahaha different from slice pointer *[]int

	for i := 0; i < len(ptrSlice); i++ {
		var tmp int = i
		ptrSlice[i] = &tmp
	}

	fmt.Printf("%v , %T \n", ptrSlice, ptrSlice)

	for i := 0; i < len(ptrSlice); i++ {
		fmt.Printf("%d , %v , %T \n", i, *ptrSlice[i], *ptrSlice[i])
	}
}
