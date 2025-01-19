package poker

import (
	"math/rand"
	"strings"
)

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	var deck Deck
	for rank := Two; rank <= Ace; rank++ {
		for suit := Clubs; suit <= Spades; suit++ {
			deck.Cards = append(deck.Cards, Card{Rank: rank, Suit: suit})
		}
	}
	return deck
}

func (deck *Deck) Shuffle() {
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

func (deck *Deck) Draw(amount int) []Card {
	cards := deck.Cards[:amount]
	deck.Cards = deck.Cards[amount:]
	return cards
}

func (deck Deck) String() string {
	sb := strings.Builder{}

	sb.WriteString("Deck: ")
	for _, card := range deck.Cards {
		sb.WriteString(card.String())
		sb.WriteString(" ")
	}

	return sb.String()
}
