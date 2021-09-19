//go:build ignore

package main

//New type declaration
type color int

//GLOBAL Variable declaration
var red int = 20
var blue color = 23

//Variable conversion => a=type(b) where type is underlying type b

func main() {

	red = int(color) //Conversion ,convert from color-type to int-type
}
