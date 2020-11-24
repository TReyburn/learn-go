package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
}

func main() {
	// don't do this
	solveig := person{"Solveig", "Delabroye", 28}
	// do this
	fmt.Println(solveig)
	travis := person{
		firstName: "Travis",
		lastName:  "Reyburn",
		age:       30,
	}
	fmt.Println(travis)

	var alex person
	fmt.Printf("%+v", alex)

	alex.firstName = "Alex"
	alex.lastName = "Smith"

	fmt.Printf("%+v", alex)
}
