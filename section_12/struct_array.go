//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

type CreditCard struct {
	holder  string
	number  int
	expDate string
	cvv     int
}

func main() {
	var ccArray [7]CreditCard

	for i := 0; i < len(ccArray); i++ {
		ccArray[i] = CreditCard{
			holder:  "Tester",
			number:  4000000000 + i,
			expDate: "01/2" + strconv.Itoa(i),
			cvv:     700 + i,
		}
	}

	for i, v := range ccArray {
		fmt.Printf("-> %d | %v \n", i, v)
	}

}
