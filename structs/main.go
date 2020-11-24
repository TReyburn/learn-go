package main

import "fmt"

type contactInfo struct {
	email string
	zip string
}

type person struct {
	firstName string
	lastName string
	age int
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		age: 42,
		contactInfo: contactInfo{
			email: "jim@partytime.io",
			zip: "80303",
		},
	}
	jim.print()
	jim.updateName("Jimothy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}