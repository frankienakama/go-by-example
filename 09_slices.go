// way more common than arrays
// slices point to a section of an array
// if the slice exceeds the length of the backing array, a new array is created and the previous one is garbage collected

package main

import "fmt"

func main() {

	// slices are typed only by the elements contained (instead of the number of elements)
	s := make([]string, 3)
	fmt.Println("emp:", s) // emp: [  ]

	// set and get the same way as arrays
	s[0] = "a"                // set index 0 to "a"
	s[1] = "b"                // "b"
	s[2] = "c"                // "c"
	fmt.Println("set:", s)    // set: [a b c]
	fmt.Println("get:", s[2]) // get: c

	// len finds length
	fmt.Println("len:", len(s)) // len: 3

	// add new values to the slice with append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) // apd: [a b c d e f]

	// slices can be copied
	c := make([]string, len(s)) // empty slice with the length of s
	copy(c, s)                  // copy into c from s
	fmt.Println("cpy:", c)      // cpy: [a b c d e f]

	// "slice" operator takes values from index 2 to 4 (5 excluded)
	l := s[2:5]
	fmt.Println("sl1:", l) // sl1: [c d e]

	// everything before index 5 (exclusive)
	l = s[:5]
	fmt.Println("sl2:", l) // sl2: [a b c d e]

	// everything after 2 (inclusive)
	l = s[2:]
	fmt.Println("sl3", l) // sl3 [c d e f]

	// declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // dcl: [g h i]

	// multi-dimensional slices--inner slices can be of varying lengths (unlike multi-dimensional arrays)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD) // 2d: [[0] [1 2] [2 3 4]]

}

// note: slices are different types than arrays, but rendered the same by fmt.Println
