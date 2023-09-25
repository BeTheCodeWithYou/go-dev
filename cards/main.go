package main

import (
	"log"
	"os"
)

func main() {
	cards := newDeck()
	//cards.print()

	err := os.WriteFile("testdata/cards", []byte(cards.toString()), 0666)
	if err != nil {
		log.Fatal("error writing to file {}", err)
	}

	//	d1, d2 := cards.hand(3)
	//	d1.print()
	//	d2.print()
}
