package main

import "fmt"

func main() {
	cards := newDeck()
	// hand, ramaningCards := deal(cards, 5)

	// hand.print()
	// ramaningCards.print()

	fmt.Println([]byte(cards.toString()))

}
