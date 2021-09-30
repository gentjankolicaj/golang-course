//go:build ignore

package main

import "fmt"

//Callback : is passing a function as argument (sometimes called functional programing)
//Function can :
//1.Assigned to variable
//2.Returned from a function
//3.Be anonymous
//4.Passed as argument to another function
//Callback template : func identifier(fx func(params) returnType,t0 type,t1 type ...) returnType{}

func main() {
	slice := []int{2, 3, 3, 5, 2, 4, 5, 6, 7, 8, 95, 5, 23, 343}
	total := sum(slice...) //Pass a slice of int as argument

	//Callback below with function even & sum
	totalEven := even(sum, slice...)

	//Callback with another function odd & negativeSum
	totalOdd := odd(negativeSum, slice...)

	fmt.Println("\n")
	fmt.Println("Total ", total)
	fmt.Println("Even ", totalEven)
	fmt.Println("Odd ", totalOdd)

	isEven(total)
}

func sum(x ...int) int {
	fmt.Printf("x %T , %v ", x, x)
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

//This is func has a callback,meaning that a functions is passed to him as argument
//Func passed as argument is fx
func even(fx func(slice ...int) int, slice ...int) int {
	var evenSlice []int
	for _, v := range slice {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}

	//Pass even slice at function argument
	value := fx(evenSlice...)
	return value
}

func negativeSum(slice ...int) int {
	var totalNeg int = 0
	for _, v := range slice {
		totalNeg -= v
	}
	return totalNeg
}
func odd(fx func(slice ...int) int, slice ...int) int {
	var oddSlice []int
	for _, v := range slice {
		if v%2 != 0 {
			oddSlice = append(oddSlice, v)
		}
	}
	var negativeSum int = 0
	//Pass oddSlice at
	negativeSum = fx(oddSlice...)
	return negativeSum
}

func isEven(i int) bool {
	return i%2 == 0
}
