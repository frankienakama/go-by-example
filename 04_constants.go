package main

import "fmt"
import "math"

// const declares a constant value
const s string = "constant"

func main() {

	fmt.Println(s) // constant

	// const statement can appear anywhere a var statement can
	const n = 500000000

	// const expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d) // 6e+11

	// numeric const has no type until given one
	fmt.Println(int64(d)) // 600000000000

	// number can be given a type by using it in a context that requires one (math.Sin expects a float64)
	fmt.Println(math.Sin(n)) // -0.28470407323754404

}
