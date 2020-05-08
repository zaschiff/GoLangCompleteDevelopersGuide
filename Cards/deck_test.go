package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

/*
*************************************************** NOTES ****************************************************
	we can write files to test our go package but these tests are not like testing with RSpec, mocha, jasmine,
	selenium, etc!

	To make a test file we need to create a new file the specifically ends with '_test.go'!!!!
	To run all the tests in the given package, we can run 'go test' at the command line.
	Test functions need to start with Test with a capitalize T

	usually when creating tests we ask: How do we know what to test? go makes this straight forward.
	deciding what to test using the deck.go code. Results are dependent to developer.
		- newDeck()
			-> length of deck
			-> first card in deck is specific
			-> last card in deck is specific
*/
