// Go supports methods based on struct types

package main

import "fmt"

type rect struct {
	width, height int
}

// this area has a 'receiver type' of *rect
func (r *rect) area() int {

	return r.width * r.height

}

// methods can be defined for either pointer or value receiver types

// value receiver:
func (r rect) perim() int {

	return 2*r.width + 2*r.height

}

// call the two methods defined for the struct
func main() {

	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())   // area: 50
	fmt.Println("perim:", r.perim()) // perim: 30

	// Go automatically handles conversion between values and pointers for method calls
	// use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct

	rp := &r
	fmt.Println("area:", rp.area())   // area: 50
	fmt.Println("perim:", rp.perim()) // perim: 30

}
