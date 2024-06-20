//go:build ignore

package main

import "fmt"

//Important declaration keyword => var identifier_name type

//When we declare a variable and not initialize default values are assigned
//Boolean => false
//Integer => 0
//Double/float => 0.0
//String =>""
//interface{} => nil
//Pointer => nil
//functions => nil
//slices =>nil
//map => nil
//channels => nil

//Global declaration,or file declaration
var flag bool

var id uint

var firstname string

var lastname string

var age int

var salary float64

var other interface{} //empty interface (interface{}) is supertype of all types

func main() {
	FunctionBodyDeclaration()

	//Call get value by identifier
	fmt.Println("Firstname identifier called ", firstname)
	fmt.Println("Other identifier called ", other)

}

func FunctionBodyDeclaration() {
	var flag bool
	var id uint
	var firstname string
	var lastname string
	var age int
	var salary float64
	var other interface{} //empty interface (interface{}) is supertype of all types

}
