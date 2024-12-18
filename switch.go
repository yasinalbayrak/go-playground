package main

import "fmt"

func main() {
	var heads, tails int

	switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
	}
}

func coinflip() string {
	return "heads"
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
