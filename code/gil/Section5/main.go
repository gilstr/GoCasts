package main

import "fmt"

func main() {
	scores := map[string]float64{
		"math":             80.0,
		"computer science": 90,
		"history":          90,
	}

	// var scores map[string]float64

	// scores := make(map[string]float64)

	fmt.Println(scores)
	deleteFromMap(scores, "history")
	fmt.Println(scores)
	printMap(scores)

}

func deleteFromMap(m map[string]float64, k string) {
	delete(m, k)
}

func printMap(m map[string]float64) {
	for s, f := range m {
		println(s, ":", f)
	}

}
