package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Error("Expected 16 decks, but got ", len(d))
	}
}

func TestSaveToFileAndReadFile(t *testing.T) {
	os.Remove("_decktesting.txt")
	deck := newDeck()
	deck.saveToFile("_decktesting.txt")
	loadedDeck := readFile("_decktesting.txt")
	if len(loadedDeck) != 16 {
		t.Error("Expected 16 cards in deck, got ", len(loadedDeck))
	}

}
