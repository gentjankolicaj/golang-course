//go:build ignore

package main

import "fmt"

//Pointer : value pointing to some memory address
//Pointer : We get pointer value by using operator &,it gives memory address
//Pointer is different from type that is pointing to
//Operator & : Gives address where value is stored
//Operator * : Gives the value that address is holding (also called de-referencing)
//Everything in GOLANG is pass by value

//Pointer : Good to be used when we have a large group of data, and we don't want to pass it around
//Pointer : To change a value at certain location
func main() {
	var arrayPtr [10]*int // ARRAY THAT HOLDS POINTERS,In golang array has predefined length

	//fill pointer array with pointer values
	for i := 0; i < len(arrayPtr); i++ {
		arrayPtr[i] = &i
	}

	fmt.Printf("%v , %T \n", arrayPtr, arrayPtr)

	for i := 0; i < len(arrayPtr); i++ {
		fmt.Printf("%d , %v , %T \n", i, *arrayPtr[i], *arrayPtr[i])
	}

}
