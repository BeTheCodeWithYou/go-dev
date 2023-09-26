package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {

	deck := deck{}

	cardSuite := []string{"Spades", "Diamond", "Heart", "Club"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five"}

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

func (d deck) saveToFile(s string) error {
	return os.WriteFile(s, []byte(d.toString()), 0666)
}

func newDeckFromFile(f string) deck {
	bs, err := os.ReadFile(f)
	if err != nil {
		log.Fatal("Error in reading from file {}", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
