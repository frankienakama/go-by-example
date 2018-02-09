// interfaces are named collections of method signatures

package main

import (
	"fmt"
	"math"
)

// basic interface for geometry shapes
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// to implement an interface, need to implement all methods in the interface

// implement 'geometry' on rects
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// implement 'geometry' on circles
func (c circle) perim() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}

// if a var has an interface type, can call methods that are in the named interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// rect and circle struct types both implement the geometry interface, can use instances of these structs as arguments to 'measure'
	measure(r)
	measure(c)
}
