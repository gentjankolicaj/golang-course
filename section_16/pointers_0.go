//go:build ignore

package main

import "fmt"

//Pointer : value pointing to some memory address
//Pointer : We get pointer value by using operator &,it gives memory address
//Pointer is different from type that is pointing to
//Operator & : Gives address where value is stored
//Operator * : Gives the value that address is holding (also called de-referencing)
//Everything in GOLANG is pass by value

func main() {
	a := 42

	//Print value
	fmt.Printf("%v , %T \n", a, a)

	//Variable aPtr is of type pointer a
	//Get pointer from a by using & operator
	aPtr := &a
	fmt.Printf("%v , %T \n", aPtr, aPtr) //& gives address where value is stored

	//Get value of aPtr
	fmt.Printf("%v , %T \n", *aPtr, *aPtr) //* gives value stored at address

	*aPtr = 222 //Change value stored at address where aPtr is pointing
	fmt.Println("New update a value ", a)
}
