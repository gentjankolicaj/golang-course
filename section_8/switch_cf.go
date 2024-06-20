//go:build ignore

package main

import "fmt"

func main() {
	SwitchWithInt()
	SwitchWithFloat()
	SwitchWithBool()
	SwitchWithString()

	SwitchWithIntAndBreak()
}

func SwitchWithInt() {
	fmt.Println("SwitchWithInt")
	var val int = 10
	switch val {
	case 0:
		fmt.Println("Case matched ", 0)
	case 1:
		fmt.Println("Case matched ", 1)
	case 2:
		fmt.Println("Case matched ", 2)
	case 3:
		fmt.Println("Case matched ", 3)
	case 4:
		fmt.Println("Case matched ", 4)
	default:
		fmt.Println("Case not matched ", val)
	}
}

func SwitchWithFloat() {
	fmt.Println("SwitchWithFloat")
	var val float64 = 3.14
	switch val {
	case 0.1:
		fmt.Println("Case matched ", 0)
	case 1.66:
		fmt.Println("Case matched ", 1)
	case 2.47:
		fmt.Println("Case matched ", 2)
	case 3.14:
		fmt.Println("Case matched ", 3)
	case 4.0:
		fmt.Println("Case matched ", 4)
	default:
		fmt.Println("Case not matched ", val)
	}

}

func SwitchWithBool() {
	fmt.Println("SwitchWithBool")
	var val bool = true
	switch val {
	case true:
		fmt.Println("Case matched ", true)
	case false:
		fmt.Println("Case matched ", false)
	default:
		fmt.Println("Case not matched ", val)
	}
}

func SwitchWithString() {
	fmt.Println("SwitchWithString")
	str := "Hello SwitchWithString"
	switch str {
	case "Hello":
		fmt.Println("Case matched ", true)
	case "SwitchWithString":
		fmt.Println("Case matched ", false)
	default:
		fmt.Println("Case not matched ", str)
	}
}

func SwitchWithIntAndBreak() {
	fmt.Println("SwitchWithIntAndBreak")
	var val int = 10
	switch val {
	case 10:
		fmt.Println("Case matched ", 10)
		val = 11
		break
	case 11:
		fmt.Println("Case matched ", 11)
		break
	case 2:
		fmt.Println("Case matched ", 2)
		break
	case 3:
		fmt.Println("Case matched ", 3)
		break
	case 4:
		fmt.Println("Case matched ", 4)
		break
	default:
		fmt.Println("Case not matched ", val)
	}
}
