package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d) != 20 {
		t.Errorf("Expected 20 cards, got %v", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const fileName = "_testFile"
	os.Remove(fileName)
	d := NewDeck()
	d.saveToFile(fileName)

	newDeck := newDeckFromFile(fileName)
	if len(newDeck) != 20 {
		t.Errorf("Expected 20 cards, got %v", len(newDeck))
	}

	os.ReadFile(fileName)
}
