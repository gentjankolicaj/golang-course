//go:build ignore

package main

import "fmt"

//Maps are DS with pairs(key,value)
//Maps are unordered list
//Maps template : map[key]value
//Get value : value,ok:=map[key]
//Add pair: map[newKey]=newValue
//Loop using range : for key,value:=range map{}
//Delete a pair: by calling builtin function delete(map,Key)

func main() {
	m := GetMap()
	fmt.Println(m)

	delete(m, 1)
	delete(m, 3)
	delete(m, 8)
	fmt.Println(m)

	//verify that 1 is deleted from map
	if val, ok := m[1]; ok == false {
		fmt.Println("Pair with key 1 deleted ", val, ok)
	}

}

func GetMap() map[int]int {
	var m map[int]int = map[int]int{
		1:    213,
		2:    23,
		3:    5,
		4:    5,
		23:   232,
		2321: 930,
		54:   343,
		8:    23834,
	}
	return m
}
