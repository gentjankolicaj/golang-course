//go:build ignore

package main

import "fmt"

//Since functions in GOLANG are first class citizens,
//This means we can return them from a function
//Template : func identifier() func() returnType {}

func main() {

	iFunc := intFunc()
	v0 := iFunc()
	fmt.Println(v0)

	//Create variable that holds func value
	fFunc := floatFunc(11)

	//Call func returned
	v1 := fFunc(22)
	fmt.Println(v1)

}

func intFunc() func() int {
	//Create anonymous func & return it
	f := func() int {
		return 999
	}

	return f
}

func floatFunc(f float64) func(f float64) float64 {
	return func(p float64) float64 {
		return f + p
	}
}
