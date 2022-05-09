package dataStructures

import (
	"fmt"
	"gil/dataStructures"
)

func main() {
	numbers := dataStructures.newIntegers()
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
