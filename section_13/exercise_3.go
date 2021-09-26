//go:build ignore

package main

import (
	"fmt"
)

type color struct {
	red, green, blue int
}
type vehicle struct {
	doors int
	color color
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t1 := truck{
		vehicle: vehicle{doors: 4,
			color: color{
				red:   100,
				green: 120,
				blue:  201,
			}},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{doors: 5,
			color: color{
				red:   123,
				green: 232,
				blue:  155,
			}},
		luxury: true,
	}

	fmt.Println("Truck-struct ", t1)
	fmt.Println("Sedan-struct ", s1)
}
