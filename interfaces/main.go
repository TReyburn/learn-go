package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func (englishBot) getGreeting() string {
	// Well just pretend we have some very custom logic here
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	// And we'll keep pretending here
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
