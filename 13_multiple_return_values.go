// Go has built-in support for multipile return values
// for example, used to return the result and the error vals from a fn

package main

import "fmt"

// (int, int) shows the function returns two ints
func vals() (int, int) {
	return 3, 7
}

func main() {

	// multiple assignment
	// keep in order of the return values
	a, b := vals()
	fmt.Println(a) // 3
	fmt.Println(b) // 7

	// return only a subset, use _
	c, _ := vals()
	fmt.Println(c) // 3

	_, c := vals()
	fmt.Println(c) // 7
}
