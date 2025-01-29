package simulator

import (
	"sync"

	"github.com/Palaszontko/texas-holdem-hand-calculator/backend/internal/poker"
)

type Simulator struct {
	config Config
}

func NewSimulator(config Config) *Simulator {
	return &Simulator{config: config}
}

func (s *Simulator) RunSimulation() (*Result, error) {
	results := make(chan poker.Result, s.config.NumIterations)

	var wg sync.WaitGroup

	iterationsPerWorker := s.config.NumIterations / s.config.NumConcurrent

	for i := 0; i < s.config.NumConcurrent; i++ {
		wg.Add(1)
		go s.simulationWorker(iterationsPerWorker, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	wins, losses, ties := 0, 0, 0

	for result := range results {
		switch result {
		case poker.Win:
			wins++
		case poker.Lose:
			losses++
		case poker.Tie:
			ties++
		}
	}

	totalIterations := float64(wins + losses + ties)

	return &Result{
		WinProbability:  float64(wins) / totalIterations,
		LoseProbability: float64(losses) / totalIterations,
		TieProbability:  float64(ties) / totalIterations,
		Iterations:      int(totalIterations),
	}, nil

}

func (s *Simulator) simulationWorker(iterations int, results chan<- poker.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < iterations; i += 1 {
		result := s.runSingleSimulation()
		results <- result
	}
}

func (s *Simulator) runSingleSimulation() poker.Result {
	deck := poker.NewDeck()
	deck = s.removeKnownCards(deck)
	deck.Shuffle()

	remainingCommunityCards := 5 - len(s.config.CommunityCards)
	remainingOpponentCards := 2 - len(s.config.OpponentHand.Cards)

	communityCards := append(s.config.CommunityCards, deck.Draw(remainingCommunityCards)...)
	opponentHand := poker.Hand{Cards: append(s.config.OpponentHand.Cards, deck.Draw(remainingOpponentCards)...)}

	result, _, _ := poker.CompareHands(s.config.PlayerHand, opponentHand, communityCards)

	return result
}

func (s *Simulator) removeKnownCards(deck poker.Deck) poker.Deck {
	knownCards := make(map[poker.Card]bool)

	for _, card := range s.config.PlayerHand.Cards {
		knownCards[card] = true
	}

	for _, card := range s.config.OpponentHand.Cards {
		knownCards[card] = true
	}

	for _, card := range s.config.CommunityCards {
		knownCards[card] = true
	}

	var remainingCards []poker.Card

	for _, card := range deck.Cards {
		if !knownCards[card] {
			remainingCards = append(remainingCards, card)
		}
	}

	return poker.Deck{Cards: remainingCards}

}
