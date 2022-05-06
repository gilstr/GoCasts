package main

import "fmt"

type student struct {
	firstName string
	lastName  string
}

func (s student) print() {
	fmt.Printf("%+v\n", s)
}

func main() {
	var s1 student
	s1.print()
	s1.firstName = "William"
	s1.print()
}
