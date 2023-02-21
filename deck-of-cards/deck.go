package main

import (
	"fmt"
	"math/rand"
	"os"
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
	for _, cards := range d {
		fmt.Println(cards)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("ERROR : ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {

	for i := range d {
		r := rand.Intn(len(d)-1)
		d[i], d[r] = d[r], d[i]
	}
}
