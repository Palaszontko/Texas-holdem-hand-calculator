package poker

import "fmt"

type Rank int

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

type Card struct {
	Rank Rank
	Suit Suit
}

func (rank Rank) String() string {
	mapRank := map[int]string{2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10", 11: "J", 12: "Q", 13: "K", 14: "A"}
	return mapRank[int(rank)]
}

func (suit Suit) String() string {
	mapSuit := map[int]string{0: "♣", 1: "♦", 2: "♥", 3: "♠"}
	return mapSuit[int(suit)]
}

func (card Card) String() string {
	return fmt.Sprintf("%s%s", card.Rank.String(), card.Suit.String())
}
