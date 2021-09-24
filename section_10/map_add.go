//go:build ignore

package main

import "fmt"

//Maps are DS with pairs(key,value)
//Maps are unordered list
//Maps template : map[key]value
//Get value : value,ok:=map[key]
//Add pair: map[newKey]=newValue
//Loop using range : key,value:=range map
func main() {
	var m map[int]int = map[int]int{
		1:    213,
		23:   232,
		2321: 930,
		54:   343,
	}

	//Adding new pairs to map
	m[2] = 2
	m[3] = 5
	m[4] = 7

	fmt.Println(m)
	fmt.Println("Range over map")

	for key, value := range m {
		fmt.Println("Key-value : ", key, value)
	}

}
