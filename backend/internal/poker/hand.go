package poker

import (
	"fmt"
	"slices"
	"sort"
)

type Hand struct {
	Cards []Card
}

func NewHand(cards ...Card) Hand {
	return Hand{Cards: cards}
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
		return (*cards)[i].Rank > (*cards)[j].Rank
	})
}

func containsRanks(cards []Card, ranks []Rank) bool {
	rankSet := make(map[Rank]bool)
	for _, card := range cards {
		rankSet[card.Rank] = true
	}

	for _, rank := range ranks {
		if !rankSet[rank] {
			return false
		}
	}
	return true
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

	if result := hand.isThreeOfAKind(communityCards); result != nil {
		fmt.Println("Found Three of a Kind")
		fmt.Printf("Best Hand: %s\n", result.BestHand)
	}

	if result := hand.isStraight(communityCards); result != nil {
		fmt.Println("Found Straight")
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

			sortCardsByRank(&kickers)

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
	var pairs [][]Card
	var pairRanks []Rank

	for rank := Ace; rank >= Two; rank-- {
		if groupedCards[rank] == 2 {
			var pairCards []Card
			for _, c := range allCards {
				if c.Rank == rank && len(pairCards) < 2 {
					pairCards = append(pairCards, c)
				}
			}
			pairs = append(pairs, pairCards)
			pairRanks = append(pairRanks, rank)
		}
	}

	if len(pairs) >= 2 {
		bestPairs := pairs[:2]

		var kicker Card
		if len(pairs) > 2 {
			kicker = pairs[2][0]
		} else {
			usedRanks := map[Rank]bool{
				pairRanks[0]: true,
				pairRanks[1]: true,
			}

			for _, c := range allCards {
				if !usedRanks[c.Rank] && (kicker.Rank == 0 || c.Rank > kicker.Rank) {
					kicker = c
				}
			}
		}

		bestHand := append(bestPairs[0], bestPairs[1]...)
		if kicker.Rank != 0 {
			bestHand = append(bestHand, kicker)
		}

		return &HandRank{
			Type:     TwoPair,
			BestHand: bestHand,
		}
	}

	return nil
}

func (hand *Hand) isThreeOfAKind(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)
	groupedCards := hand.groupCardsByRank(allCards)

	var threeOfAKindCards []Card
	var threeOfAKindRank Rank

	for rank := Ace; rank >= Two; rank-- {
		if groupedCards[rank] >= 3 {
			threeOfAKindRank = rank

			for _, c := range allCards {
				if c.Rank == rank && len(threeOfAKindCards) < 3 {
					threeOfAKindCards = append(threeOfAKindCards, c)
				}
			}
			break
		}
	}

	if len(threeOfAKindCards) == 3 {
		var kickers []Card
		for rank := Ace; rank >= Two; rank-- {
			if rank != threeOfAKindRank {
				for _, c := range allCards {
					if c.Rank == rank {
						kickers = append(kickers, c)
						break
					}
				}
				if len(kickers) == 2 {
					break
				}
			}
		}

		bestHand := append(threeOfAKindCards, kickers...)

		return &HandRank{
			Type:     ThreeOfAKind,
			BestHand: bestHand,
		}
	}

	return nil
}

func (hand *Hand) isStraight(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)

	rankPresent := make(map[Rank]bool)
	for _, card := range allCards {
		rankPresent[card.Rank] = true
	}

	for highCard := Ace; highCard >= Six; highCard-- {
		straight := true
		var straightCards []Card

		for i := 0; i < 5; i++ {
			currentRank := highCard - Rank(i)
			if !rankPresent[currentRank] {
				straight = false
				break
			}
			for _, card := range allCards {
				if card.Rank == currentRank {
					straightCards = append(straightCards, card)
					break
				}
			}
		}

		if straight {
			slices.Reverse(straightCards)
			return &HandRank{
				Type:     Straight,
				BestHand: straightCards,
			}
		}
	}

	if rankPresent[Ace] && rankPresent[Two] && rankPresent[Three] &&
		rankPresent[Four] && rankPresent[Five] {
		var straightCards []Card
		for _, rank := range []Rank{Ace, Two, Three, Four, Five} {
			for _, card := range allCards {
				if card.Rank == rank {
					straightCards = append(straightCards, card)
					break
				}
			}
		}
		return &HandRank{
			Type:     Straight,
			BestHand: straightCards,
		}
	}

	return nil
}

func (hand *Hand) isFlush(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)
	groupedCards := hand.groupCardsBySuit(allCards)

	for suit, amount := range groupedCards {
		if amount >= 5 {
			sortCardsByRank(&allCards)

			var bestHand []Card
			for _, card := range allCards {
				if card.Suit == suit && len(bestHand) < 5 {
					bestHand = append(bestHand, card)
				}
			}

			return &HandRank{
				Type:     Flush,
				BestHand: bestHand,
			}
		}
	}

	return nil
}

func (hand *Hand) isFullHouse(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)
	groupedCards := hand.groupCardsByRank(allCards)

	var threeOfAKindRanks []Rank
	var pairRanks []Rank

	for rank, amount := range groupedCards {
		if amount >= 3 {
			threeOfAKindRanks = append(threeOfAKindRanks, rank)
		}
		if amount == 2 {
			pairRanks = append(pairRanks, rank)
		}
	}

	if len(threeOfAKindRanks) >= 2 {
		slices.SortFunc(threeOfAKindRanks, func(a, b Rank) int {
			return int(b) - int(a)
		})

		var bestHand []Card
		for _, card := range allCards {
			if card.Rank == threeOfAKindRanks[0] {
				bestHand = append(bestHand, card)
			}
		}
		count := 0
		for _, card := range allCards {
			if card.Rank == threeOfAKindRanks[1] && count < 2 {
				bestHand = append(bestHand, card)
				count++
			}
		}

		return &HandRank{
			Type:     FullHouse,
			BestHand: bestHand,
		}
	}

	if len(threeOfAKindRanks) == 1 && len(pairRanks) >= 1 {
		var bestHand []Card

		for _, card := range allCards {
			if card.Rank == threeOfAKindRanks[0] {
				bestHand = append(bestHand, card)
			}
		}

		slices.SortFunc(pairRanks, func(a, b Rank) int {
			return int(b) - int(a)
		})
		count := 0
		for _, card := range allCards {
			if card.Rank == pairRanks[0] && count < 2 {
				bestHand = append(bestHand, card)
				count++
			}
		}

		return &HandRank{
			Type:     FullHouse,
			BestHand: bestHand,
		}
	}

	return nil
}

func (hand *Hand) isFourOfAKind(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)
	groupedCards := hand.groupCardsByRank(allCards)

	for card, amount := range groupedCards {
		if amount == 4 {
			var fourOfAKindCards []Card
			for _, c := range allCards {
				if c.Rank == card {
					fourOfAKindCards = append(fourOfAKindCards, c)
				}
			}

			var kicker Card
			sortCardsByRank(&allCards)
			fmt.Println(allCards)
			for _, c := range allCards {
				if c.Rank != card {
					kicker = c
					break
				}
			}

			return &HandRank{
				Type:     FourOfAKind,
				BestHand: append(fourOfAKindCards, kicker),
			}
		}
	}

	return nil
}

func (hand *Hand) isStraightFlush(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)
	groupedCards := hand.groupCardsBySuit(allCards)

	for suit, amount := range groupedCards {
		if amount >= 5 {
			var flushCards []Card
			for _, card := range allCards {
				if card.Suit == suit {
					flushCards = append(flushCards, card)
				}
			}

			handCopy := NewHand()

			if result := handCopy.isStraight(flushCards); result != nil {
				fmt.Println("Straight Flush")
				sortCardsByRank(&allCards)
				fmt.Printf("allCards: %v\n", allCards)
				fmt.Printf("flushCards: %v\n", flushCards)

				return &HandRank{
					Type:     StraightFlush,
					BestHand: result.BestHand,
				}
			}
		}
	}

	return nil
}

func (hand *Hand) isRoyalFlush(communityCards []Card) *HandRank {
	allCards := append(hand.Cards, communityCards...)
	groupedCards := hand.groupCardsBySuit(allCards)

	for suit, amount := range groupedCards {
		if amount >= 5 {
			var flushCards []Card
			for _, card := range allCards {
				if card.Suit == suit {
					flushCards = append(flushCards, card)
				}
			}

			if containsRanks(flushCards, []Rank{Ace, King, Queen, Jack, Ten}) {
				sortCardsByRank(&flushCards)

				return &HandRank{
					Type:     RoyalFlush,
					BestHand: flushCards[:5],
				}
			}
		}
	}

	return nil
}
