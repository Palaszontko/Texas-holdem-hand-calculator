package simulator

import (
	"fmt"
	"time"

	"github.com/Palaszontko/texas-holdem-hand-calculator/backend/internal/poker"
)

func TrySimulatorWithRanomCardsOnHand(iterations int, concurrent int) {
	deck := poker.NewDeck()
	deck.Shuffle()

	playerHand := poker.NewHand(deck.Draw(2)...)
	opponentHand := poker.NewHand(deck.Draw(2)...)

	communityCards := deck.Draw(3)

	config := Config{
		PlayerHand:     playerHand,
		OpponentHand:   opponentHand,
		CommunityCards: communityCards,
		NumIterations:  iterations,
		NumConcurrent:  concurrent,
	}

	sim := NewSimulator(config)

	start := time.Now()

	result, err := sim.RunSimulation()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	duration := time.Since(start)

	fmt.Println("\n🎲 POKER SIMULATION RESULTS")
	fmt.Println("════════════════════════════════════════════")
	fmt.Printf("🎴 Your Hand:       %v\n", playerHand)
	fmt.Printf("🎴 Opponent's Hand: %v\n", opponentHand)
	fmt.Printf("🃏 Community Cards: %v\n\n", communityCards)
	fmt.Printf("📈 RESULTS (from %d simulations):\n", result.Iterations)
	fmt.Printf("🏆 Win:  %.2f%%\n", result.WinProbability*100)
	fmt.Printf("🤝 Tie:  %.2f%%\n", result.TieProbability*100)
	fmt.Printf("❌ Lose: %.2f%%\n", result.LoseProbability*100)
	fmt.Printf("\n⚡ Completed in: %v\n", duration.Round(time.Millisecond))
}
