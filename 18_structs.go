// structs are typed collections of field
// useful for grouping data together to form records

package main

import "fmt"

// the 'person' struct has name and age fields
type person struct {
	name string
	age  int
}

func main() {

	// syntax to create a new struct
	fmt.Println(person{"Bob", 20}) // {Bob 20}

	// can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30}) // {Alice 30}

	// omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"}) // {Fred 0}

	// & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}

	// access struct fields with .
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Sean

	// can also dereference struct pointers with .
	sp := &s
	fmt.Println(sp.age) // 50

	// structs are mutable (value can be changed)
	sp.age = 51
	fmt.Println(sp.age) // 51

}
