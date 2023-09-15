package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type deck which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// opt 1 - log the error and call newdeck()
		// opt 2 - log the error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	for i := range d {
		new_position := rand.Intn(len(d) - 1)
		d[i], d[new_position] = d[new_position], d[i]
	}
}

func (d deck) shuffle2() {
	rand.Shuffle(len(d), func(i, j int) {

		d[i], d[j] = d[j], d[i]
	})
}
