package poker

import "testing"

func TestRankString(t *testing.T) {
	tests := []struct {
		name     string
		rank     Rank
		expected string
	}{
		{"Test Ace", Ace, "A"},
		{"Test King", King, "K"},
		{"Test Queen", Queen, "Q"},
		{"Test Jack", Jack, "J"},
		{"Test Ten", Ten, "10"},
		{"Test Two", Two, "2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rank.String(); got != tt.expected {
				t.Errorf("Rank.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSuitString(t *testing.T) {
	tests := []struct {
		name     string
		suit     Suit
		expected string
	}{
		{"Test Clubs", Clubs, "♣"},
		{"Test Diamonds", Diamonds, "♦"},
		{"Test Hearts", Hearts, "♥"},
		{"Test Spades", Spades, "♠"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.suit.String(); got != tt.expected {
				t.Errorf("Suit.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCardString(t *testing.T) {
	tests := []struct {
		name     string
		card     Card
		expected string
	}{
		{
			name:     "Test Ace of Spades",
			card:     Card{Rank: Ace, Suit: Spades},
			expected: "A♠",
		},
		{
			name:     "Test Two of Hearts",
			card:     Card{Rank: Two, Suit: Hearts},
			expected: "2♥",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.String(); got != tt.expected {
				t.Errorf("Card.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestHandString(t *testing.T) {
	tests := []struct {
		name     string
		hand     Hand
		expected string
	}{
		{
			name:     "Empty Hand",
			hand:     Hand{Cards: []Card{}},
			expected: "Hand: []",
		},
		{
			name:     "One Card Hand",
			hand:     Hand{Cards: []Card{{Ace, Spades}}},
			expected: "Hand: [A♠]",
		},
		{
			name:     "Two Card Hand",
			hand:     Hand{Cards: []Card{{Ace, Spades}, {King, Hearts}}},
			expected: "Hand: [A♠ K♥]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hand.String(); got != tt.expected {
				t.Errorf("Hand.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}
