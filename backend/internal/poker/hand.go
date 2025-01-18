package poker

import (
	"fmt"
	"sort"
)

type Hand struct {
	Cards []Card
}

func (hand Hand) String() string {
	return fmt.Sprintf("Hand: %s", hand.Cards)
}

type HandRank struct {
	Type     HandRankType
	BestHand []Card
}

type HandRankType int

const (
	HighCard HandRankType = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

func (hand *Hand) groupCardsByRank(cards []Card) map[Rank]int {
	groupedCards := make(map[Rank]int)
	for _, card := range cards {
		groupedCards[card.Rank]++
	}
	return groupedCards
}

func (hand *Hand) groupCardsBySuit(cards []Card) map[Suit]int {
	groupedCards := make(map[Suit]int)
	for _, card := range cards {
		groupedCards[card.Suit]++
	}
	return groupedCards
}

func sortCardsByRank(cards *[]Card) {
	sort.Slice(*cards, func(i, j int) bool {
		return (*cards)[i].Rank < (*cards)[j].Rank
	})
}

func (hand *Hand) Test(communityCards []Card) {

	if result := hand.isPair(communityCards); result != nil {
		fmt.Println("Found Pair")
		fmt.Printf("Best Hand: %s\n", result.BestHand)
	}

}

func (hand *Hand) isPair(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)
	groupedCards := hand.groupCardsByRank(allCards)

	for _, card := range allCards {
		if groupedCards[card.Rank] == 2 {
			pairCards := make([]Card, 0, 2)
			kickers := make([]Card, 0, 3)

			for _, c := range allCards {
				if c.Rank == card.Rank {
					pairCards = append(pairCards, c)
				} else {
					kickers = append(kickers, c)
				}
			}

			sort.Slice(kickers, func(i, j int) bool {
				return kickers[i].Rank > kickers[j].Rank
			})
			kickers = kickers[:min(3, len(kickers))]

			return &HandRank{
				Type:     Pair,
				BestHand: append(pairCards, kickers...),
			}
		}
	}
	return nil
}
