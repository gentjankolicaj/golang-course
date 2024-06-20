//go:build ignore

package main

import (
	"fmt"
)

func main() {

	//Switch with no switch-expression=> switch{} => switch true{} by default
	switch {
	case false:
		fmt.Println("Should not print")
	case true:
		fmt.Println("Should print 1")
	case true:
		fmt.Println("Should not print-2")
		break
	case true:
		fmt.Println("Should not print-3")
	default:
		fmt.Println("Should not print default.")
	}
}
