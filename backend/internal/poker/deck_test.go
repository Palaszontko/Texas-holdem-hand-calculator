package poker

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.Cards) != 52 {
		t.Errorf("expected 52 cards, got %d", len(deck.Cards))
	}

	suitCount := make(map[Suit]int)
	rankCount := make(map[Rank]int)

	for _, card := range deck.Cards {
		suitCount[card.Suit]++
		rankCount[card.Rank]++
	}

	for suit := Clubs; suit <= Spades; suit++ {
		if suitCount[suit] != 13 {
			t.Errorf("expected 13 cards of suit %v, got %d", suit, suitCount[suit])
		}
	}

	for rank := Two; rank <= Ace; rank++ {
		if rankCount[rank] != 4 {
			t.Errorf("expected 4 cards of rank %v, got %d", rank, rankCount[rank])
		}
	}
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	original := make([]Card, 52)
	copy(original, deck.Cards)

	deck.Shuffle()

	if len(deck.Cards) != 52 {
		t.Errorf("shuffle changed deck size to %d", len(deck.Cards))
	}

	matches := 0
	for i := range deck.Cards {
		if deck.Cards[i] == original[i] {
			matches++
		}
	}

	if matches == 52 {
		t.Error("deck was not shuffled")
	}
}

func TestDraw(t *testing.T) {
	deck := NewDeck()
	card := deck.Draw()

	if len(deck.Cards) != 51 {
		t.Errorf("expected 51 cards after draw, got %d", len(deck.Cards))
	}

	for _, c := range deck.Cards {
		if c == card {
			t.Error("drawn card still in deck")
		}
	}
}

func TestString(t *testing.T) {
	deck := NewDeck()
	str := deck.String()

	if str == "" {
		t.Error("string representation is empty")
	}

	for _, card := range deck.Cards {
		cardStr := card.String()
		if cardStr == "" {
			t.Error("card string representation is empty")
		}
	}
}
