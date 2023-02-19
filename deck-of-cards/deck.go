package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuite := []string{"Spades", "Diamond", "Hearts", "Club"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuite {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) printDeck() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	 return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
