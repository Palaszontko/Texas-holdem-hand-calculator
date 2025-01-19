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

	if result := hand.isTwoPair(communityCards); result != nil {
		fmt.Println("Found Two Pair")
		fmt.Printf("Best Hand: %s\n", result.BestHand)
	}

}

func (hand *Hand) isPair(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)
	groupedCards := hand.groupCardsByRank(allCards)

	for card := Ace; card >= Two; card -= 1 {
		if groupedCards[card] == 2 {

			var pairCards []Card
			var kickers []Card

			for _, c := range allCards {
				if c.Rank == card {
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

func (hand *Hand) isTwoPair(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)
	groupedCards := hand.groupCardsByRank(allCards)

	usedCards := map[Card]bool{}

	var pairs [][]Card
	var kickers []Card
	for card := Ace; card >= Two; card -= 1 {
		if groupedCards[card] == 2 {
			var pairCards []Card
			for _, c := range allCards {
				if c.Rank == card {
					pairCards = append(pairCards, c)
					usedCards[c] = true
				} else if !usedCards[c] {
					kickers = append(kickers, c)
				}
				if len(pairCards) == 2 {
					pairs = append(pairs, pairCards)
					break
				}
			}
		}
	}

	if len(pairs) == 2 {
		sortCardsByRank(&kickers)
		kickers = kickers[:min(1, len(kickers))]
		return &HandRank{
			Type:     TwoPair,
			BestHand: append(append(pairs[0], pairs[1]...), kickers...),
		}
	}

	return nil

}
