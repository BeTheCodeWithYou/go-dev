package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("expected deck of 20 cards but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Club" {
		t.Errorf("Expected Five of Club but got %v", len(d)-1)
	}
}

func TestNewDeckSaveToFileAndLoadDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck)!=20 {
		t.Errorf("Expected deck of 20 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
