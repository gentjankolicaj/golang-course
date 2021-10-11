//go:build ignored

package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat string
	lon string
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Sqrt error : %v %v %v ", se.lat, se.lon, se.err)
}

func main() {
	value, err := sqrt(-122)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 64, sqrtError{"23", "12", errors.New("Test")}
	}
	return 42, nil
}
