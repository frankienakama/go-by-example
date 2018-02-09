// communicate errors via an explicit, separate return value (unlike exceptions in Ruby)
// makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks

package main

import "fmt"
import "errors"

// by convention, errors are the last return value and has type 'error'
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic 'error' value with the given error message
		return -1, errors.New("can't work with 42")
	}
	// a nil value in the error position indicates that there was no error
	return arg + 3, nil
}

// can use custom types as errors by implementing the Error() method
// use custom type to explicitly represent an argument error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob) // f1 worked: 10
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// this uses &argError syntax to build a new struct
		// supplies two values for the two fields 'arg' and 'prob'
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// two loops test out each of the error-returning functions
	// use the use of an inline error check on the 'if' like is a common idiom in Go
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
		// f1 failed: can't work with 42
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
		// f2 failed: 42 - can't work with it
	}

	// to programmatically use the data in a custom error, you'll need to get the error as an instance of the custom error type via type assertion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)  // 42
		fmt.Println(ae.prob) // can't work with it
	}
}
