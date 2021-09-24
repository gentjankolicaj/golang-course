//go:build ignore

package main

import "fmt"

func main() {
	var favSport string = "football"

	switch favSport {
	case "volleyball":
		fmt.Println("Volleyball")
	case "baseball":
		fmt.Println("Baseball")
	case "football":
		fmt.Println("Football")
	case "tennis":
		fmt.Println("Tennis")
	default:
		fmt.Println("Golf")
	}

}
