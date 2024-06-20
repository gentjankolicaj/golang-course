//go:build ignored

package main

import (
	"encoding/json"
	"log"
)

type phone struct {
	Brand   string
	Version string
}

func main() {
	nexus := phone{"Nexus", "1.20"}

	byteSlice, err := json.Marshal(nexus) //NOTE : is a bad practice to not handle error
	if err != nil {
		log.Fatalln(err) //Prints error & shuts process down
	} else {
		log.Println("OK-", string(byteSlice))
	}

}
