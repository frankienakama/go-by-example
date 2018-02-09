// switch statements express conditionals across many branches

package main

import "fmt"
import "time"

func main() {

	// basic switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// Write 2 as two

	// can use commas to separate multiple expressions in teh same case statement
	// default is optional
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	// It's a weekday

	// switch without an expression is an alt way to express if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	// It's after noon

	// (type) switch compares types instead of values
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool") // I'm a bool
		case int:
			fmt.Println("I'm an int") // I'm an int
		default:
			fmt.Printf("Don't know type %T\n", t) // Don't know type string
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
