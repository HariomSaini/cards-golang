package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Not the expected length of deck, But got- %v ", len(d))
	}
	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected Card is Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "King Of Clubs" {
		t.Errorf("Expected Card is King of Clubs but got %v", d[len(d)-1])
	}
}
