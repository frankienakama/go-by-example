// range iterates over elements

package main

import "fmt"

func main() {

	// range used to sum the numbers in a slice
	// range on arrays and slices provides the index and value for each entry

	// if we want the index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i) // index: 1
		}
	}

	// don't need the indexes here, so it is ignored using _
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) // sum: 9

	// range on maps iterate over k/v pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
		// a -> apple
		// b -> banana
	}

	// iterate over just the keys in a map
	for k := range kvs {
		fmt.Println("key:", k)
		// key: a
		// key: b
	}

	// range on strings iterates over the Unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
		// 0 103
		// 1 111
	}
}
