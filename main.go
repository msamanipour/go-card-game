package main

func main() {
	// Create New Deck
	// cards := newDeck()

	//Deal Cards & print them
	// hand, ramaningCards := deal(cards, 5)

	// hand.print()
	// ramaningCards.print()

	//Save to file
	// cards.saveToFile("my_cards")

	//Get from file
	cards := newDeckFromFile("my_cards")
	cards.print()

}
