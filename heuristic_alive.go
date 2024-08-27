package main

import (
  "github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicAliveAllies calculates a heuristic score based on the number of alive ally snakes
func HeuristicAliveAllies(snapshot agent.GameSnapshot) float64 {
  aliveAllies := 0
  for _, allySnake := range snapshot.YourTeam() {
    if allySnake.Health() > 0 {
      aliveAllies++
    }
  }
  return float64(100 * aliveAllies)
}