package main

import "fmt"

func main() {
	p := new(int)   // p, of type *int, points to an unnamed int variable
	fmt.Println(*p) // "0"
	*p = 2          // sets the unnamed int to 2
	fmt.Println(*p)
}

func newInt() *int {
	return new(int)
}

// Identical behaviours ^v

func newInt2() *int {
	var dummyInt int
	return &dummyInt
}

func delta(old, new int) int {
	return new - old
}
