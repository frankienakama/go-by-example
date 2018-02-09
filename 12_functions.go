package main

import "fmt"

// takes two ints (a and b) and returns their sum as an int (the type following the params)
func plus(a int, b int) int {

	// Go requires explicit returns
	return a + b

}

// multiple consecutive parameters of the same type can share a type name as the last parameter
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// call functions like normal (assign to var, call using var name)
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res) // 1 + 2 = 3

	res = plusPlus(1, 2, 3)         // change value of res to plusPlus fn
	fmt.Println("1 + 2 + 3 =", res) // 1 + 2 + 3 = 6

}
