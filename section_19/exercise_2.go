//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Id   int
	Name string
}

//Unmarshall json
func main() {
	jsonString := `[{"Id":4,"Name":"Four"},{"Id":3,"Name":"Three"},{"Id":1,"Name":"One"},{"Id":2,"Name":"Two"},{"Id":3,"Name":"Three"},{"Id":1,"Name":"One"},{"Id":4,"Name":"Four"}]`

	var persons []person

	err := json.Unmarshal([]byte(jsonString), &persons)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	for i, v := range persons {
		fmt.Printf("Person idx %v | value %v | type %T \n", i, v, v)
	}

}
