package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Test",
		lastName:  "Test2",
		contactInfo: contactInfo{
			email:   "test@example.com",
			zipcode: 111,
		},
	}

	alex.updateName("Ant")
	alex.print()
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Println(p)
}
