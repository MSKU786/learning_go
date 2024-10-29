package main

import (
	"os"
	"testing"
)

func TestNewDec(t *testing.T) {
		d := newDeck();

		if (len(d) != 16) {
				t.Errorf("Excpected deck length is 16 but we got %v", len(d));
		}

		if (d[0] != "Ace of Spades") {
			  t.Errorf("Expected Ace of spades but we got %v", d[0]);
		}

		if (d[len(d)-1] != "Four of Clubs") {
				t.Errorf("Expected Four of Clubs but we got %v", d[len(d)-1])
		}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting");

	deck := newDeck()
	deck.saveToFile("_decktesting");

	loadedDeck := newDeckFromFile("_decktesting");

	if (len(loadedDeck) != 16) {
			t.Errorf("Expected 16 cards from file deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting");
}