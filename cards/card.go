package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {

	deck := deck{}

	cardSuite := []string{"Ace", "Diamond", "Heart", "Club"}
	cardValue := []string{"One", "Two", "Three", "Four", "Five"}

	for _, s := range cardSuite {
		for _, v := range cardValue {
			deck = append(deck, v+" of "+s)
		}
	}
	return deck
}

func (d deck) hand(n int) (deck, deck) {
	return d[:n], d[n:]
}

func (d deck) print() {

	for i, deck := range d {
		fmt.Println(i, deck)
	}
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}

