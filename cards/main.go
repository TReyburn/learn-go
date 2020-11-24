package main

func main() {
	cards := newDeck()
	cards.saveToFile("test")
}

func newCard() string {
	return "Five of Diamonds"
}
