//go:build ignored

package main

import (
	"fmt"

	"github.com/gentjankolicaj/golang-course/tree/main/section_27/dog"
)

type data struct {
	id       int
	flag     bool
	message  string
	testData []float64
}

type canine struct {
	name string
	age  int
}

func main() {
	var dataArray [3]data
	for i := 0; i < len(dataArray); i++ {
		dataArray[i] = data{
			id:       i,
			flag:     true,
			message:  "Testing array struct",
			testData: []float64{2, 2, 3, 4, 2, 32, 4, 3, 2.342, 2, 2.3},
		}
	}

	for i, v := range dataArray {
		fmt.Printf("index %d , value %v \n", i, v)
	}

	rex := canine{
		name: "Rex",
		age:  2,
	}

	fmt.Println(dog.FromHumanToDog(rex.age))

}
