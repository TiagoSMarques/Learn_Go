package main

func main() {
	cards := newDeck()
	// hand, rem_deck := deal(cards, 5)
	// hand.print()
	// rem_deck.print()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}
