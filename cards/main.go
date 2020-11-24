package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards, _ = deal(cards, 5)
	cards.print()
}
