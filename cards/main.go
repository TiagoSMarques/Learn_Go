package main

func main() {
	cards := newDeck()
	hand, rem_deck := deal(cards, 5)
	hand.print()
	rem_deck.print()

	cards.saveToFile("my_cards")

	cards2 := newDeckFromFile("my_cards")

	cards2.shuffle()
	cards2.shuffle2()
	cards2.print()

}
