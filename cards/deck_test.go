package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}
}

func TestNewDeckFirstValue(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades' but got %v", d[0])
	}
}

func TestNewDeckLastValue(t *testing.T) {
	d := newDeck()

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of 'Four of Clubs' but got %v", d[len(d)-1])
	}
}

func TestSaveAndLoadFromDisk(t *testing.T) {
	_ = os.Remove("_decktesting")

	d := newDeck()
	_ = d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, got %v", len(loadedDeck))
	}

	_ = os.Remove("_decktesting")
}
