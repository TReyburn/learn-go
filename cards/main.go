package main

import "fmt"

func main() {
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_")
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
