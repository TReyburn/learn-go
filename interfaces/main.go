package main

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

func main() {

}
