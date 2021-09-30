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

type person struct {
	first string
	last  string
}

func main() {
	//Define p1 of type person
	var p1 person

	p1 = person{
		first: "john",
		last:  "doe",
	}

	fmt.Printf("%v , %T \n", p1, p1)
	changeMe(&p1)

	//IMPORTANT:TO DEREFERENCE A STRUCT I used (*value).field
	//For example
	p1Ptr := &p1
	fmt.Println((*p1Ptr).first, (*p1Ptr).last) //de-referencing a pointer struct
}

func changeMe(p *person) {
	fmt.Println("changeMe(p *person) called \n")
	*p = person{
		first: "Changed-first",
		last:  "Changed-last",
	}
}
