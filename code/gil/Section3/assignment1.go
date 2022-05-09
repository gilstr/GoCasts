package dataStructures

import "fmt"

func main() {
	for i := range [10]int{} {
		fmt.Printf("%v: ", i)
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
