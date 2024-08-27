package main

import (
  "math"

  "github.com/Battle-Bunker/cyphid-snake/agent"
  "github.com/samber/lo"
)

func HeuristicRelativeLength(snapshot agent.GameSnapshot) float64 {
  allies := snapshot.YourTeam()
  opponents := snapshot.Opponents()

  if len(opponents) == 0 {
    return float64(lo.SumBy(allies, func(s agent.SnakeSnapshot) int { return s.Length() }))
  }

  smallestOpponent := lo.MinBy(opponents, func(a, b agent.SnakeSnapshot) bool {
    return a.Length() < b.Length()
  }).Length()
  largestOpponent := lo.MaxBy(opponents, func(a, b agent.SnakeSnapshot) bool {
    return a.Length() > b.Length()
  }).Length()

  scoreSnake := func(snake agent.SnakeSnapshot) float64 {
    relativeLength := float64(snake.Length() - smallestOpponent + 5)
    sensitiveRange := float64(largestOpponent - smallestOpponent + 10)

    // S-curve function
    score := 1 / (1 + math.Exp(-4*(relativeLength/sensitiveRange-0.5)))

    // Scale the score to a reasonable range
    return score * 100
  }

  return lo.SumBy(allies, scoreSnake)
}