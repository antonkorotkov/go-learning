package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		t.Errorf("Expected deck to be 28 cards but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected ace of spades but got %v", d)
	}

	if d[len(d)-1] != "Seven of Clubs" {
		t.Errorf("Expected Seven of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 28 {
		t.Errorf("Expected 28 but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
