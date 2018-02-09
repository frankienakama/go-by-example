// goroutine is a lightweight thread of execution
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// synchronous: wait for it to finish before moving on to the next task
// asynchronous: can move on to another task before it finishes

func main() {
	// normal way of calling function f(s) made above, running synchronously
	f("direct")

	// to invoke function in a goroutine, use go f(s)
	// this will execute concurrently (simultaneously) with the calling one
	go f("goroutine")

	// can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// two function calls are running asynchronously in separate goroutines now
	// execution falls through to here
	// Scanln() requires a keypress before the program exits
	fmt.Scanln()
	fmt.Println("done")
}

// direct : 0
// direct : 1
// direct : 2
// goroutine : 0
// goroutine : 1
// goroutine : 2
// going
// <enter> // requires we press a key before the program exits
// done
