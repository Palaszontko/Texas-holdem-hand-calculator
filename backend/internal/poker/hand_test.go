package poker

import (
	"reflect"
	"testing"
)

func TestIsPair(t *testing.T) {
	tests := []struct {
		name           string
		hand           Hand
		communityCards []Card
		want           *HandRank
	}{
		{
			name: "Pair of aces in hand",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Hearts},
				{Rank: Queen, Suit: Diamonds},
				{Rank: Jack, Suit: Clubs},
			},
			want: &HandRank{
				Type: Pair,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
					{Rank: King, Suit: Hearts},
					{Rank: Queen, Suit: Diamonds},
					{Rank: Jack, Suit: Clubs},
				},
			},
		},
		{
			name: "Pair with community card",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Ace, Suit: Hearts},
				{Rank: Queen, Suit: Diamonds},
				{Rank: Jack, Suit: Clubs},
			},
			want: &HandRank{
				Type: Pair,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
					{Rank: King, Suit: Hearts},
					{Rank: Queen, Suit: Diamonds},
					{Rank: Jack, Suit: Clubs},
				},
			},
		},
		{
			name: "No pair",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Queen, Suit: Hearts},
				{Rank: Jack, Suit: Diamonds},
				{Rank: Ten, Suit: Clubs},
			},
			want: nil,
		},
		{
			name: "Pair with fewer kickers",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Hearts},
			},
			want: &HandRank{
				Type: Pair,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
					{Rank: King, Suit: Hearts},
				},
			},
		},
		{
			name: "Edge case: Multiple pairs, should return highest",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Two, Suit: Hearts},
				{Rank: Two, Suit: Diamonds},
				{Rank: Ace, Suit: Clubs},
			},
			want: &HandRank{
				Type: Pair,
				BestHand: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
					{Rank: Ace, Suit: Clubs},
					{Rank: Two, Suit: Hearts},
					{Rank: Two, Suit: Diamonds},
				},
			},
		},
		{
			name: "Edge case: Empty community cards",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{},
			want: &HandRank{
				Type: Pair,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
		},
		{
			name: "Edge case: All same suit, different ranks",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: King, Suit: Spades},
				},
			},
			communityCards: []Card{
				{Rank: Queen, Suit: Spades},
				{Rank: Jack, Suit: Spades},
				{Rank: Ten, Suit: Spades},
			},
			want: nil, // Should not identify as pair even though all cards are same suit
		},
		{
			name: "Edge case: Same rank pair in community cards only",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: Queen, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Ace, Suit: Hearts},
				{Rank: Ace, Suit: Diamonds},
				{Rank: Two, Suit: Clubs},
			},
			want: &HandRank{
				Type: Pair,
				BestHand: []Card{
					{Rank: Ace, Suit: Hearts},
					{Rank: Ace, Suit: Diamonds},
					{Rank: King, Suit: Spades},
					{Rank: Queen, Suit: Hearts},
					{Rank: Two, Suit: Clubs},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hand.isPair(tt.communityCards)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
