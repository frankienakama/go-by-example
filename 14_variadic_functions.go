// variadic functions can be called with any number of trailing arguments
// fmt.Println is a variadic fn

package main

import "fmt"

// takes and arbitrary number of ints as arguments
func sum(nums ...int) {

	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func main() {

	// called like normal functions, but with however many arguments
	sum(1, 2)    // [1 2] 3
	sum(1, 2, 3) // [1 2 3] 6

	// if you already have a slice with the values you need, use the slice!
	nums := []int{1, 2, 3, 4}
	sum(nums...) // [1 2 3 4] 10

}
