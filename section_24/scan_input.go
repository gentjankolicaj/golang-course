//go:build ignored

package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string

	fmt.Println("Type name : ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Type surname : ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Type age : ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println("Answer 1 ", answer1)
	fmt.Println("Answer 2 ", answer2)
	fmt.Println("Answer 3 ", answer3)

}
