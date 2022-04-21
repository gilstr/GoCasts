package main

import "fmt"

type integers []int

func newIntegers() integers {
	ints := integers{}

	for i := 1; i < 10; i++ {
		for _, j := range []int{10, 100, 1000} {
			ints = append(ints, i*j)
		}
	}
	return ints
}

func (ints integers) print() {
	for i, v := range ints {
		fmt.Println(i, ":", v)
	}
}
