<<<<<<< HEAD
//go:build ignored

=======
>>>>>>> 54a2b972697d4b04b69a676a3ecb75818dc94b81
package main

import (
	"fmt"
)

func main() {

	fmt.Println("1,2,3,4,5 => sum = ", Sum(1, 2, 3, 4, 5))
}

func Sum(x ...int) int {
	var sum int = 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
