package main

import (
	"fmt"
	"github.com/Palaszontko/texas-holdem-hand-calculator/backend/internal/api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/api/health", handlers.HealthCheckHandler)
	http.HandleFunc("/api/simulation", handlers.SimulationHander)

	fmt.Println("Starting server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
