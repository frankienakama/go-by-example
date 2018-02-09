// Go supports anonymous functions, which can form closures
// useful when you want to define a function inline without having to name it

package main

import "fmt"

// intSeq reurns another function, defined anonymously in the body
func intSeq() func() int {

	i := 0
	// the returned function 'closes over' the var i to form a closure
	return func() int {

		i++
		return i

	}

}

func main() {

	// nextInt is called, assigning the returned anon function to nextInt
	nextInt := intSeq()

	// this new function captures its own i value, updated each time nextInt is called
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// confirm the new state is unique to nextInt by creating and testing a new one
	newInts := intSeq()
	fmt.Println(newInts()) // 1

}
