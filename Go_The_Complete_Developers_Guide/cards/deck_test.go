package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected len 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {

		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {

		t.Errorf("Expected first card to be King of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loaded_deck := newDeckFromFile("_decktesting")
	if len(loaded_deck) != 16 {

		t.Errorf("Expected len 16, but got %v", len(d))
	}
	os.Remove("_decktesting")

}
