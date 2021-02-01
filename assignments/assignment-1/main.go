package main

import "fmt"

func main() {

	// We could do one single loop to iterate from 0 to 10 and print,
	// but since the exercise explicitly says we should create a
	// slice of ints, I've created a first loop to populate the
	// slice and a for/range to check whether the current item is
	// odd or even and print.

	collection := []int{}

	for i := 0; i <= 10; i++ {
		collection = append(collection, i)
	}

	for _, i := range collection {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
