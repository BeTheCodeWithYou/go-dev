package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected 16 but found: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("expected Ace of Spades but found ", d[0])
	}
}
