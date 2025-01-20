# Texas Hold'em Hand Calculator

A poker simulator leveraging Monte Carlo methods to calculate winning probabilities in Texas Hold'em poker games. The system employs advanced statistical modeling and parallel processing to achieve high accuracy in hand strength evaluation.

## Overview

The system implements a Monte Carlo simulation engine that:

- Runs hundreds of thousands of randomized poker scenarios
- Utilizes parallel processing for optimal performance
- Provides statistical confidence through large sample sizes
- Simulates realistic poker scenarios by considering all possible opponent hands

## Core Features

- Advanced Monte Carlo probability engine
- Complete poker hand evaluation (High Card through Royal Flush)
- Statistical analysis of winning/losing/tie scenarios
- Multiple-thread simulation processing
- Configurable simulation depth
- Built-in test simulation function for quick verification

### Test Simulation Function

The system includes a built-in test function `TrySimulatorWithRanomCardsOnHand` that demonstrates the Monte Carlo simulation capabilities:

```go
func TrySimulatorWithRanomCardsOnHand(iterations int, concurrent int) {
	deck := poker.NewDeck()
	deck.Shuffle()

	playerHand := poker.NewHand(deck.Draw(2)...)

	communityCards := deck.Draw(5)

	config := Config{
		PlayerHand:     playerHand,
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
	fmt.Printf("🃏 Community Cards: %v\n\n", communityCards)
	fmt.Printf("📈 RESULTS (from %d simulations):\n", result.Iterations)
	fmt.Printf("🏆 Win:  %.2f%%\n", result.WinProbability*100)
	fmt.Printf("🤝 Tie:  %.2f%%\n", result.TieProbability*100)
	fmt.Printf("❌ Lose: %.2f%%\n", result.LoseProbability*100)
	fmt.Printf("\n⚡ Completed in: %v\n", duration.Round(time.Millisecond))
}
```
```
	🎲 POKER SIMULATION RESULTS
════════════════════════════════════════════
🎴 Your Hand:       Hand: [Q♦ K♠]
🃏 Community Cards: [5♦ 3♣ J♣ Q♠ 5♣]

📈 RESULTS (from 500000 simulations):
🏆 Win:  81.43%
🤝 Tie:  0.51%
❌ Lose: 18.07%

⚡ Completed in: 1.054s
```

## API Endpoints

### POST /api/simulation

Executes the Monte Carlo simulation for given cards.

#### Request Body

```json
{
  "playerCards": [
    { "Rank": 14, "Suit": 0 },
    { "Rank": 13, "Suit": 0 }
  ],
  "communityCards": [{ "Rank": 10, "Suit": 0 }],
  "numIterations": 100000,
  "numConcurrent": 8
}
```

#### Response

```json
{
  "winProbability": 0.65,
  "loseProbability": 0.3,
  "tieProbability": 0.05,
  "iterations": 100000
}
```

### GET /api/health

Health check endpoint.

## Technical Details

### Statistical Model Parameters

- Default iteration count: 100,000 simulations
- Maximum iteration count: 500,000 simulations
- Concurrent workers: 8-16 threads
- Confidence level: Increases with iteration count

### Card Representation

- Ranks: 2-14 (2 through Ace)
- Suits: 0-3 (Clubs, Diamonds, Hearts, Spades)

## Performance Optimizations

- Parallel execution of Monte Carlo simulations
- Optimized card comparison algorithms
- Efficient memory management for large simulation sets
- Configurable thread count for different hardware capabilities

## Configuration

### Monte Carlo Parameters

- `numIterations`: Simulation depth (higher = more accurate)
- `numConcurrent`: Thread count for parallel processing

### Server Configuration

- Default port: 8080
- CORS enabled for localhost:5173

## Limitations

- Maximum 500,000 iterations per request (statistical accuracy vs. performance trade-off)
- Maximum 16 concurrent workers (hardware optimization)
- Heads-up scenarios only (1v1 probability calculations)

## TODO

### In Progress

- [ ] Svelte frontend setup
- [ ] Interactive card picker component
- [ ] Real-time simulation results display
- [ ] Mobile responsiveness

### Future Ideas

- [ ] Add multiple opponents simulation
