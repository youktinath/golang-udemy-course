package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck of length 20, got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades' but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	d := newDeck()
	d.saveToFile("_decktesting.txt")
	loadedDeck := newDeckFromFile("_decktesting.txt")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck of length 20, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")
}