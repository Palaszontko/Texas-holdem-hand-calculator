package simulator

import (
	"github.com/Palaszontko/texas-holdem-hand-calculator/backend/internal/poker"
)

type Config struct {
	PlayerHand     poker.Hand
	OpponentHand   poker.Hand
	CommunityCards []poker.Card
	NumIterations  int
	NumConcurrent  int
}
