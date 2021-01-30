package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSilce(mySlice)

	fmt.Println(mySlice)

	fmt.Println(&mySlice)

	fmt.Println(*&mySlice)
}

func updateSilce(s []string) {
	s[0] = "Bye"

	s = append(s, "xpto")
}
