//go:build ignore

package main

import (
	"fmt"
)

//Maps are DS with pairs(key,value)
//Maps are unordered list
//Maps template : map[key]value

func main() {
	var m map[int]string //Created map
	fmt.Println(m)

	m2 := map[int]string{
		1: "James Bond",
		2: "Agent47",
		3: "Captain America",
		4: "Black widow",
	}
	fmt.Println(m2)

	//Access some value by key
	//If you enter a key and there is no value stored on map for that key,
	//Then returned value is zero-value
	value, ok := m[1]
	fmt.Println(value, ok)

	value1, ok1 := m2[3]
	fmt.Println(value1, ok1)

	//Another expression with value & ok flag
	//this is called ; ok idiom used to distinguish from zero-value
	if value2, ok2 := m2[4]; ok2 {
		fmt.Println("Comma-OK idiom ", value2, ok2)
	}
}
