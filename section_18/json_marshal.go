//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type mobile struct {
	Company string
	Brand   string
	Version string
	Os      string
	Imei    int
}

func main() {

	nexus := mobile{
		Company: "LG",
		Brand:   "Nexus",
		Version: "5",
		Os:      "Android",
		Imei:    232323232,
	}

	iphone6 := mobile{
		Company: "Apple",
		Brand:   "Iphone",
		Version: "6",
		Os:      "iOS",
		Imei:    12112111,
	}

	//Marshal nexus value
	byteSlice0, err0 := json.Marshal(nexus)
	if err0 != nil {
		fmt.Println(err0)
	} else {
		fmt.Println(string(byteSlice0))
	}

	//Marshal iphone6 value
	byteSlice1, err1 := json.Marshal(iphone6)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(string(byteSlice1))
	}

}
