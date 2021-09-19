//go:build ignore

package main

//WE can create our custom types using existing primitive types
//GLOBAL DECLARATION : Below I have declared 5 new types
type boolean bool
type whole_number uint
type integer_number int
type real_number float64
type complex_number complex128
type student struct {
	id   uint
	name string
}

func main() {

	var flag boolean
	var a whole_number
	var b integer_number
	var c real_number
	var d complex_number
	var newStudent student

	//I call function to avoid compile error:GOLANG doesn't allow non-used declared variables
	avoidCompileError(flag, a, b, c, d, newStudent)

	LocalCustomType()

	var i int
	var j integer_number
	//ALSA
	// j=i => produces compile error because i is of type int
	//and j is of type integer_number which in turn is of type int

	//NOTE: GO IS ABOUT TYPES and is statically typed

}

func LocalCustomType() {
	//LOCAL TYPE declaration
	type boolean bool
	type whole_number uint
	type integer_number int
	type real_number float64
	type complex_number complex128
	type student struct {
		id   uint
		name string
	}

	//VARIABLE DECLARATION OF NEW TYPE
	var flag boolean
	var a whole_number
	var b integer_number
	var c real_number
	var d complex_number
	var newStudent student

	//I call function to avoid compile error:GOLANG doesn't allow non-used declared variables
	avoidCompileError(flag, a, b, c, d, newStudent)

}

func avoidCompileError(t ...interface{}) {

}
