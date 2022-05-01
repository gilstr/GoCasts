package main

import (
	"fmt"
)

func main() {
	numbers := newIntegers()
	numbers.print()
	numbers.shuffle()
	numbers.print()
}

func getInt(x int) int {
	return x + 1
}

func printHorizontalLine() {
	fmt.Println("-----------------------------------")
}
