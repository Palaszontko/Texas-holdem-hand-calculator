package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Palaszontko/texas-holdem-hand-calculator/backend/internal/poker"
	simulator "github.com/Palaszontko/texas-holdem-hand-calculator/backend/internal/simulation"
)

type SimulationRequest struct {
	PlayerCards    []poker.Card `json:"playerCards"`
	OpponentCards  []poker.Card `json:"opponentCards,omitempty"`
	CommunityCards []poker.Card `json:"communityCards,omitempty"`
	NumIterations  int          `json:"numIterations"`
	NumConcurrent  int          `json:"numConcurrent"`
}

type SimulationResponse struct {
	WinProbability  float64 `json:"winProbability"`
	LoseProbability float64 `json:"loseProbability"`
	TieProbability  float64 `json:"tieProbability"`
	Iterations      int     `json:"iterations"`
}

func SimulationHander(w http.ResponseWriter, r *http.Request) {
	allowedOrigins := []string{"http://localhost:5173",
		"http://localhost:4173",
		"https://texas-holdem-hand-calculator-api.onrender.com"}

	origin := r.Header.Get("Origin")
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			break
		}
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SimulationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if len(req.PlayerCards) != 2 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if len(req.CommunityCards) > 5 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if req.NumIterations <= 0 {
		req.NumIterations = 100_000
	} else if req.NumIterations > 500_000 {
		req.NumIterations = 500_000
	}

	if req.NumConcurrent <= 0 {
		req.NumConcurrent = 8
	} else if req.NumConcurrent > 16 {
		req.NumConcurrent = 16
	}

	fmt.Printf("Request - Player cards: %v, Opponent cards: %v Community cards: %v, Iterations: %d\n",
		req.PlayerCards, req.OpponentCards, req.CommunityCards, req.NumIterations)

	config := simulator.Config{
		PlayerHand:     poker.NewHand(req.PlayerCards...),
		OpponentHand:   poker.NewHand(req.OpponentCards...),
		CommunityCards: req.CommunityCards,
		NumIterations:  req.NumIterations,
		NumConcurrent:  req.NumConcurrent,
	}

	sim := simulator.NewSimulator(config)

	result, err := sim.RunSimulation()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp := SimulationResponse{
		WinProbability:  result.WinProbability,
		LoseProbability: result.LoseProbability,
		TieProbability:  result.TieProbability,
		Iterations:      result.Iterations,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
