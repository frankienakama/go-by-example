// pointers 'point to' values stored elsewhere in computer memory

// this example uses two functions
package main

import "fmt"

// zeroval() has an int param, so args will be passed to it by value
func zeroval(ival int) {

	ival = 0

}

// zeroptr() has an *int param, so it takes an int pointer as an argument
func zeroptr(iptr *int) {

	// dereferences pointer from memory address to the current value at that address
	// assigning a value to a dereferenced pointer changes the value at the referenced address
	*iptr = 0

}

func main() {

	i := 1
	fmt.Println("initial:", i) // initial: 1

	zeroval(i)
	fmt.Println("zeroval:", i) // zeroval: 1

	// the &i syntax gives the memory address of i (aka a pointer to i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // zeroptr: 0

	// pointers can be printed too
	fmt.Println("pointer:", &i) // pointer: 0xc4200160a8

}

// zeroval doesn't change the i in main, but zeroptr does, because it has a reference to the memory address for that var

// pointer analogy: analogy, a page number in a book's index could be considered a pointer to the corresponding page; dereferencing such a pointer would be done by flipping to the page with the given page number and reading the text found on the indexed page
