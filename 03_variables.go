// variables are explicitly declared and used by the compiler

package main

import "fmt"

func main() {

	// var declares 1+ variables
	var a = "initial"
	fmt.Println(a) // initial

	// can declare multiple vars at once
	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2

	// Go infers the type for initialized vars
	var d = true
	fmt.Println(d) // true

	// vars declared without init are zero-valued
	var e int
	fmt.Println(e) // 0

	// := syntax is shorthand for declaring and initializing a var
	f := "short"
	fmt.Println(f) // short
}
