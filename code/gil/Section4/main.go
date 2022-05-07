package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (s student) print() {
	fmt.Printf("%+v\n", s)
}

func (s *student) updateFirstName(newFirstName string) {
	s.firstName = newFirstName

}
func main() {
	s1 := student{
		firstName: "William",
		lastName:  "Clinton",
		contactInfo: contactInfo{
			email:   "bill@whitehouse.org",
			zipCode: 100000,
		},
	}

	s1.print()
	s1.updateFirstName("Bill")
	s1.print()
}
