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

func TestIsTwoPair(t *testing.T) {
	tests := []struct {
		name           string
		hand           Hand
		communityCards []Card
		want           *HandRank
	}{
		{
			name: "Two pairs in hand",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Hearts},
				{Rank: King, Suit: Diamonds},
				{Rank: Queen, Suit: Clubs},
			},
			want: &HandRank{
				Type: TwoPair,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
					{Rank: King, Suit: Hearts},
					{Rank: King, Suit: Diamonds},
					{Rank: Queen, Suit: Clubs},
				},
			},
		},
		{
			name: "One pair in hand, one in community",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Queen, Suit: Hearts},
				{Rank: Queen, Suit: Diamonds},
				{Rank: Two, Suit: Clubs},
			},
			want: &HandRank{
				Type: TwoPair,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
					{Rank: Queen, Suit: Hearts},
					{Rank: Queen, Suit: Diamonds},
					{Rank: Two, Suit: Clubs},
				},
			},
		},
		{
			name: "Both pairs in community cards",
			hand: Hand{
				Cards: []Card{
					{Rank: Ten, Suit: Spades},
					{Rank: Nine, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Ace, Suit: Hearts},
				{Rank: Ace, Suit: Diamonds},
				{Rank: King, Suit: Hearts},
				{Rank: King, Suit: Clubs},
				{Rank: Two, Suit: Clubs},
			},
			want: &HandRank{
				Type: TwoPair,
				BestHand: []Card{
					{Rank: Ace, Suit: Hearts},
					{Rank: Ace, Suit: Diamonds},
					{Rank: King, Suit: Hearts},
					{Rank: King, Suit: Clubs},
					{Rank: Ten, Suit: Spades},
				},
			},
		},
		{
			name: "No two pair",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Queen, Suit: Hearts},
				{Rank: Queen, Suit: Diamonds},
				{Rank: Jack, Suit: Clubs},
			},
			want: nil,
		},
		{
			name: "Edge case: Three pairs available, should select highest two",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Hearts},
				{Rank: King, Suit: Diamonds},
				{Rank: Queen, Suit: Hearts},
				{Rank: Queen, Suit: Clubs},
			},
			want: &HandRank{
				Type: TwoPair,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
					{Rank: King, Suit: Hearts},
					{Rank: King, Suit: Diamonds},
					{Rank: Queen, Suit: Hearts},
				},
			},
		},
		{
			name: "Edge case: Empty community cards with pair in hand",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{},
			want:           nil,
		},
		{
			name: "Edge case: Multiple high kickers available",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Queen, Suit: Hearts},
				{Rank: Queen, Suit: Diamonds},
				{Rank: Ace, Suit: Clubs},
				{Rank: Jack, Suit: Spades},
			},
			want: &HandRank{
				Type: TwoPair,
				BestHand: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
					{Rank: Queen, Suit: Hearts},
					{Rank: Queen, Suit: Diamonds},
					{Rank: Ace, Suit: Clubs},
				},
			},
		},
		{
			name: "Edge case: Exactly seven cards with three pairs",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Hearts},
				{Rank: King, Suit: Diamonds},
				{Rank: Queen, Suit: Hearts},
				{Rank: Queen, Suit: Clubs},
				{Rank: Jack, Suit: Spades},
			},
			want: &HandRank{
				Type: TwoPair,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
					{Rank: King, Suit: Hearts},
					{Rank: King, Suit: Diamonds},
					{Rank: Queen, Suit: Hearts},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hand.isTwoPair(tt.communityCards)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isTwoPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsThreeOfAKind(t *testing.T) {
	tests := []struct {
		name           string
		hand           Hand
		communityCards []Card
		want           *HandRank
	}{
		{
			name: "Three of a kind in hand",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Ace, Suit: Diamonds},
				{Rank: King, Suit: Hearts},
				{Rank: Queen, Suit: Clubs},
			},
			want: &HandRank{
				Type: ThreeOfAKind,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
					{Rank: Ace, Suit: Diamonds},
					{Rank: King, Suit: Hearts},
					{Rank: Queen, Suit: Clubs},
				},
			},
		},
		{
			name: "Three of a kind with two in community",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Diamonds},
				{Rank: Ace, Suit: Hearts},
				{Rank: Queen, Suit: Clubs},
			},
			want: &HandRank{
				Type: ThreeOfAKind,
				BestHand: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
					{Rank: King, Suit: Diamonds},
					{Rank: Ace, Suit: Hearts},
					{Rank: Queen, Suit: Clubs},
				},
			},
		},
		{
			name: "Three of a kind all in community",
			hand: Hand{
				Cards: []Card{
					{Rank: Ten, Suit: Spades},
					{Rank: Nine, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Hearts},
				{Rank: King, Suit: Diamonds},
				{Rank: King, Suit: Clubs},
				{Rank: Two, Suit: Hearts},
				{Rank: Three, Suit: Clubs},
			},
			want: &HandRank{
				Type: ThreeOfAKind,
				BestHand: []Card{
					{Rank: King, Suit: Hearts},
					{Rank: King, Suit: Diamonds},
					{Rank: King, Suit: Clubs},
					{Rank: Ten, Suit: Spades},
					{Rank: Nine, Suit: Hearts},
				},
			},
		},
		{
			name: "No three of a kind",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Queen, Suit: Hearts},
				{Rank: Queen, Suit: Diamonds},
				{Rank: Jack, Suit: Clubs},
			},
			want: nil,
		},
		{
			name: "Edge case: Multiple high kickers available",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Diamonds},
				{Rank: Ace, Suit: Hearts},
				{Rank: Queen, Suit: Diamonds},
				{Rank: Jack, Suit: Clubs},
				{Rank: Ten, Suit: Spades},
			},
			want: &HandRank{
				Type: ThreeOfAKind,
				BestHand: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
					{Rank: King, Suit: Diamonds},
					{Rank: Ace, Suit: Hearts},
					{Rank: Queen, Suit: Diamonds},
				},
			},
		},
		{
			name: "Edge case: Empty community cards with pair in hand",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Ace, Suit: Hearts},
				},
			},
			communityCards: []Card{},
			want:           nil,
		},
		{
			name: "Edge case: Exactly one kicker available",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Diamonds},
				{Rank: Ace, Suit: Hearts},
			},
			want: &HandRank{
				Type: ThreeOfAKind,
				BestHand: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
					{Rank: King, Suit: Diamonds},
					{Rank: Ace, Suit: Hearts},
				},
			},
		},
		{
			name: "Edge case: Four of a kind available",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: King, Suit: Diamonds},
				{Rank: King, Suit: Clubs},
				{Rank: Ace, Suit: Hearts},
			},
			want: &HandRank{
				Type: ThreeOfAKind,
				BestHand: []Card{
					{Rank: King, Suit: Spades},
					{Rank: King, Suit: Hearts},
					{Rank: King, Suit: Diamonds},
					{Rank: Ace, Suit: Hearts},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hand.isThreeOfAKind(tt.communityCards)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isThreeOfAKind()\ngot  = %+v\nwant = %+v", got, tt.want)
			}
		})
	}
}

func TestIsStraight(t *testing.T) {
	tests := []struct {
		name           string
		hand           Hand
		communityCards []Card
		want           *HandRank
	}{
		{
			name: "Straight in hand higher than board",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: Queen, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Jack, Suit: Hearts},
				{Rank: Ten, Suit: Diamonds},
				{Rank: Nine, Suit: Clubs},
			},
			want: &HandRank{
				Type: Straight,
				BestHand: []Card{
					{Rank: Nine, Suit: Clubs},
					{Rank: Ten, Suit: Diamonds},
					{Rank: Jack, Suit: Hearts},
					{Rank: Queen, Suit: Hearts},
					{Rank: King, Suit: Spades},
				},
			},
		},
		{
			name: "Ace-high straight",
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
			want: &HandRank{
				Type: Straight,
				BestHand: []Card{
					{Rank: Ten, Suit: Clubs},
					{Rank: Jack, Suit: Diamonds},
					{Rank: Queen, Suit: Hearts},
					{Rank: King, Suit: Hearts},
					{Rank: Ace, Suit: Spades},
				},
			},
		},
		{
			name: "Wheel straight (A-5)",
			hand: Hand{
				Cards: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Two, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Three, Suit: Hearts},
				{Rank: Four, Suit: Diamonds},
				{Rank: Five, Suit: Clubs},
			},
			want: &HandRank{
				Type: Straight,
				BestHand: []Card{
					{Rank: Ace, Suit: Spades},
					{Rank: Two, Suit: Hearts},
					{Rank: Three, Suit: Hearts},
					{Rank: Four, Suit: Diamonds},
					{Rank: Five, Suit: Clubs},
				},
			},
		},
		{
			name: "No straight - missing card",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: Queen, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Jack, Suit: Hearts},
				{Rank: Ten, Suit: Diamonds},
				{Rank: Eight, Suit: Clubs}, // Missing Nine
			},
			want: nil,
		},
		{
			name: "Multiple possible straights - should pick highest",
			hand: Hand{
				Cards: []Card{
					{Rank: Nine, Suit: Spades},
					{Rank: Eight, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Seven, Suit: Hearts},
				{Rank: Six, Suit: Diamonds},
				{Rank: Five, Suit: Clubs},
				{Rank: Four, Suit: Spades},
				{Rank: Three, Suit: Hearts},
			},
			want: &HandRank{
				Type: Straight,
				BestHand: []Card{
					{Rank: Five, Suit: Clubs},
					{Rank: Six, Suit: Diamonds},
					{Rank: Seven, Suit: Hearts},
					{Rank: Eight, Suit: Hearts},
					{Rank: Nine, Suit: Spades},
				},
			},
		},
		{
			name: "Edge case: Empty community cards",
			hand: Hand{
				Cards: []Card{
					{Rank: Nine, Suit: Spades},
					{Rank: Eight, Suit: Hearts},
				},
			},
			communityCards: []Card{},
			want:           nil, // Can't make straight with just 2 cards
		},
		{
			name: "Edge case: All same suit, different consecutive ranks",
			hand: Hand{
				Cards: []Card{
					{Rank: King, Suit: Spades},
					{Rank: Queen, Suit: Spades},
				},
			},
			communityCards: []Card{
				{Rank: Jack, Suit: Spades},
				{Rank: Ten, Suit: Spades},
				{Rank: Nine, Suit: Spades},
			},
			want: &HandRank{
				Type: Straight,
				BestHand: []Card{
					{Rank: Nine, Suit: Spades},
					{Rank: Ten, Suit: Spades},
					{Rank: Jack, Suit: Spades},
					{Rank: Queen, Suit: Spades},
					{Rank: King, Suit: Spades},
				},
			},
		},
		{
			name: "Edge case: Straight in community cards only",
			hand: Hand{
				Cards: []Card{
					{Rank: Two, Suit: Spades},
					{Rank: Three, Suit: Hearts},
				},
			},
			communityCards: []Card{
				{Rank: Nine, Suit: Hearts},
				{Rank: Eight, Suit: Diamonds},
				{Rank: Seven, Suit: Clubs},
				{Rank: Six, Suit: Spades},
				{Rank: Five, Suit: Hearts},
			},
			want: &HandRank{
				Type: Straight,
				BestHand: []Card{
					{Rank: Five, Suit: Hearts},
					{Rank: Six, Suit: Spades},
					{Rank: Seven, Suit: Clubs},
					{Rank: Eight, Suit: Diamonds},
					{Rank: Nine, Suit: Hearts},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hand.isStraight(tt.communityCards)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}
