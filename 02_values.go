// value types include strings, int, floats, bools

package main

import "fmt"

func main() {

	// concat strings
	fmt.Println("go" + "lang") // golang

	// string, sum
	fmt.Println("1+1 =", 1+1) // 1+1 = 2
	// string, quotient
	fmt.Println("7.0/3.0 =", 7.0/3.0) // 7.0/3.0 = 2.3333333333333335

	// false, can't be both
	fmt.Println(true && false) // false
	// always true
	fmt.Println(true || false) // true
	// is not true
	fmt.Println(!true) // false

}
