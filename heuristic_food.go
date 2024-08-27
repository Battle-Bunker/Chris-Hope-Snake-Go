package main

import (
	"github.com/BattlesnakeOfficial/rules"
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/samber/lo"
	"math"
)

// manhattanDistance calculates the Manhattan distance between two points
func manhattanDistance(p1, p2 rules.Point) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}

// calculateFoodProximityScore calculates the food proximity score for a single snake
func calculateFoodProximityScore(snake agent.SnakeSnapshot, food []rules.Point) float64 {
	head := snake.Head()

	// Calculate the sum of inverse Manhattan distances to all food
	foodScore := lo.SumBy(food, func(f rules.Point) float64 {
		dist := manhattanDistance(head, f)
		return 1.0 / float64(dist+1)
	})

	// If the snake has just eaten (health is 100), add 1 to simulate a food at distance 0
	if snake.Health() == 100 {
		foodScore += 1.0
	}

	return foodScore * 100 // consider each food to be worth 100 points
}

// HeuristicFoodProximity calculates the food proximity heuristic score for the entire team
func HeuristicFoodProximity(snapshot agent.GameSnapshot) float64 {
	food := snapshot.Food()

	// Calculate the sum of food proximity scores for each allied snake
	return lo.SumBy(snapshot.YourTeam(), func(snake agent.SnakeSnapshot) float64 {
		return calculateFoodProximityScore(snake, food)
	})
}