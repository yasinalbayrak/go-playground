package main

import (
	"bufio"
	"fmt"
	"os"
)

// Make is similar to new keyword in C
// maps in GO are reference type, they are pointers to actual data.
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	myMap := map[string]int{
		"go":     1,
		"python": 2,
	}

	println(myMap)

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
