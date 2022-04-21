package main

import "fmt"

func main() {
	numbers := integers{getInt(0), getInt(1), getInt(2)}
	numbers = append(numbers, getInt(len(numbers)))

	numbers.print()

	printHorizontalLine()

	numbers2 := newIntegers()

	printHorizontalLine()

	numbers2.print()

	printHorizontalLine()

	numbers2.subset(1, 5).print()

	printHorizontalLine()

	numbers2.reset()
	numbers2.print()
}

func getInt(x int) int {
	return x + 1
}

func printHorizontalLine() {
	fmt.Println("-----------------------------------")
}
