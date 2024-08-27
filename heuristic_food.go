package main

import (
	"math"

	"github.com/BattlesnakeOfficial/rules"
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicFoodProximity calculates a score based on how close the team's snakes are to their nearest food
func HeuristicFoodProximity(snapshot agent.GameSnapshot) float64 {
	totalScore := 0.0
	maxDistance := snapshot.Width() + snapshot.Height() // Maximum possible Manhattan distance

	for _, snake := range snapshot.YourTeam() {
		nearestFoodDistance := maxDistance
		snakeHead := snake.Head()

		for _, foodPoint := range snapshot.Food() {
			distance := manhattanDistance(snakeHead, foodPoint)
			if distance < nearestFoodDistance {
				nearestFoodDistance = distance
			}
		}

		// Invert the distance so that closer food results in a higher score
		snakeScore := maxDistance - nearestFoodDistance
		totalScore += float64(snakeScore)
	}

	return totalScore
}

// manhattanDistance calculates the Manhattan distance between two points
func manhattanDistance(p1, p2 rules.Point) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}