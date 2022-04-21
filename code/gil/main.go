package main

import "fmt"

func main() {
	numbers := integers{getInt(0), getInt(1), getInt(2)}
	numbers = append(numbers, getInt(len(numbers)))

	i := 0
	i = 1
	fmt.Println(i)

	numbers.print()

	numbers2 := newIntegers()
	numbers2.print()
}

func getInt(x int) int {
	return x + 1
}
