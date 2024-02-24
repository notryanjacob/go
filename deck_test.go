package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	d.saveToFile("temp")
	if len(d) != 16 {
		t.Errorf("Expected deck of length 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected four of clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := readFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected length 16, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
