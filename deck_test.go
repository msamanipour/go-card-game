package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("String length must be 16 but now is %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//Remove all _decktesting file from before
	os.Remove("_decktesting")

	//Create newDeck and save to file
	deck := newDeck()
	deck.saveToFile("_decktesting")

	//Load saved file from hard
	loadedDeck := newDeckFromFile("_decktesting")

	//Check if loaded successful
	if len(loadedDeck) != 16 {
		t.Errorf("String length must be 16 but now is %v", len(loadedDeck))
	}

	//remove testing file
	os.Remove("_decktesting")
}
