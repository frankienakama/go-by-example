// maps are the same as hashes/dictionaries

package main

import "fmt"

func main() {

	// create an empty map using 'make'
	// mapname := make(map[key-type]val-type)
	m := make(map[string]int)

	// set key//val pairs using the typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	// printing a map will show all it's k/v pairs
	fmt.Println("map:", m) // map: map[k1:7 k2:13]

	// get value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1: ", v1) // v1:  7

	//  len finds number of k/v pairs in a map
	fmt.Println("len:", len(m)) // len: 2

	// remove a kv pair
	delete(m, "k2")
	fmt.Println("map:", m) // map: map[k1:7]

	// ???????
	_, prs := m["k2"]
	fmt.Println("prs:", prs) // prs: false

	// declare and initialize a map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n) // map: map[foo:1 bar:2]
}

// note: prints as [k:v k:v] with fmt.Println
