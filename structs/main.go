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
	contact contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		age: 42,
		contact: contactInfo{
			email: "jim@partytime.io",
			zip: "80303",
		},
	}

	fmt.Printf("%+v", jim)
}
