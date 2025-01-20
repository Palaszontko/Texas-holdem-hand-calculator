package poker

type Result int

const (
	Win Result = iota
	Lose
	Tie
)

func CompareHands(hand1 Hand, hand2 Hand, communityCards []Card) (Result, HandRank, HandRank) {
	result1 := hand1.EvaluateHandStrenght(communityCards)
	result2 := hand2.EvaluateHandStrenght(communityCards)

	if result1.Type > result2.Type {
		return Win, *result1, *result2
	} else if result1.Type < result2.Type {
		return Lose, *result2, *result1
	} else {
		switch compareHandWithTie(result1, result2) {
		case Win:
			return Win, *result1, *result2
		case Lose:
			return Lose, *result2, *result1
		default:
			return Tie, *result1, *result2
		}
	}
}

func compareHandWithTie(result1 *HandRank, result2 *HandRank) Result {
	switch result1.Type {
	case HighCard:
		return compareHandsIterative(result1.BestHand, result2.BestHand)
	case Pair:
		return compareHandsIterative(result1.BestHand, result2.BestHand)
	case TwoPair:
		return compareHandsIterative(result1.BestHand, result2.BestHand)
	case ThreeOfAKind:
		return compareHandsIterative(result1.BestHand, result2.BestHand)
	case Straight:
		return compareHandsIterative(result1.BestHand, result2.BestHand)
	case Flush:
		return compareHandsIterative(result1.BestHand, result2.BestHand)
	case FullHouse:
		return compareHandsIterative(result1.BestHand, result2.BestHand)
	case FourOfAKind:
		return compareHandsIterative(result1.BestHand, result2.BestHand)
	case StraightFlush:
		return compareHandsIterative(result1.BestHand, result2.BestHand)
	default:
		return Tie
	}

}

func compareHandsIterative(cards1 []Card, cards2 []Card) Result {
	cards1Copy, cards2Copy := make([]Card, len(cards1)), make([]Card, len(cards2))
	copy(cards1Copy, cards1)
	copy(cards2Copy, cards2)

	sortCardsByRank(&cards1Copy)
	sortCardsByRank(&cards2Copy)

	for i := 0; i < len(cards1Copy); i++ {
		if cards1Copy[i].Rank > cards2Copy[i].Rank {
			return Win
		} else if cards1Copy[i].Rank < cards2Copy[i].Rank {
			return Lose
		}
	}

	return Tie
}
