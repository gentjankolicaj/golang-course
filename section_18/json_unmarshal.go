//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type Mobile struct {
	Company string `json:"Company"`
	Brand   string `json:"Brand"`
	Version string `json:"Version"`
	Os      string `json:"Os"`
	Imei    int    `json:"Imei"`
}

func main() {
	str := "[{\"Company\":\"LG\",\"Brand\":\"Nexus\",\"Version\":\"5\",\"Os\":\"Android\",\"Imei\":232323232},{\"Company\":\"Apple\",\"Brand\":\"Iphone\",\"Version\":\"6\",\"Os\":\"iOS\",\"Imei\":12112111}]"

	//Convert rune into slice of bytes
	data := []byte(str)

	var mobileSlice []Mobile

	err := json.Unmarshal(data, &mobileSlice)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range mobileSlice {
		fmt.Printf("%v , %v , %T \n", i, v, v)
	}

}
