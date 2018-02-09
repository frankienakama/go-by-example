// for is Go's only looping construct

package main

import "fmt"

func main() {

	// most basic type of loop
	// single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// 1
	// 2
	// 3

	// initial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// 7
	// 8
	// 9

	// for without a condition will loop until given 'break' or 'return'
	for {
		fmt.Println("loop")
		break
	}
	// loop

	// can also 'continue' to next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // move on to the lext loop, don't print this line
		}
		fmt.Println(n)
	}
	// 1
	// 3
	// 5

}
