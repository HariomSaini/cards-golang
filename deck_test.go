package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Not the expected length of deck, But got- %v ", len(d))
	}
}
