package main

import (
	"math"

	"github.com/BattlesnakeOfficial/rules"
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

// HeuristicFloodFill calculates a heuristic score based on available space for each ally snake
func HeuristicFloodFill(snapshot agent.GameSnapshot) float64 {
	totalScore := 0.0
	for _, allySnake := range snapshot.YourTeam() {
		rawScore := floodFillScore(snapshot, allySnake)
		transformedScore := transformScore(rawScore)
		totalScore += transformedScore
	}
	return totalScore
}

// transformScore applies a sigmoid-like function to flatten the output to at most 100 points per snake
func transformScore(score float64) float64 {
	maxScore := 100.0
	k := 0.05 // Adjusts the steepness of the curve
	return maxScore * (2 / (1 + math.Exp(-k*score)) - 1)
}

func floodFillScore(snapshot agent.GameSnapshot, snake agent.SnakeSnapshot) float64 {
	visited := make(map[rules.Point]bool)
	queue := []rules.Point{snake.Head()}
	score := 0.0
	snakeTail := snake.Body()[len(snake.Body())-1]

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		// Add to score for each accessible cell
		score++

		// Check if we've reached the snake's tail
		if current == snakeTail {
			score += float64(snake.Length()) * 3
		}

		// Add neighboring cells to the queue
		neighbors := getNeighbors(current, snapshot.Width(), snapshot.Height())
		for _, neighbor := range neighbors {
			if !visited[neighbor] && isValidMove(snapshot, neighbor) {
				queue = append(queue, neighbor)
			}
		}
	}

	return score
}

func getNeighbors(p rules.Point, width, height int) []rules.Point {
	neighbors := []rules.Point{
		{X: p.X, Y: p.Y - 1}, // Up
		{X: p.X, Y: p.Y + 1}, // Down
		{X: p.X - 1, Y: p.Y}, // Left
		{X: p.X + 1, Y: p.Y}, // Right
	}

	validNeighbors := []rules.Point{}
	for _, n := range neighbors {
		if n.X >= 0 && n.X < width && n.Y >= 0 && n.Y < height {
			validNeighbors = append(validNeighbors, n)
		}
	}
	return validNeighbors
}

func isValidMove(snapshot agent.GameSnapshot, p rules.Point) bool {
	// Check if the point is occupied by any snake's body (except tails)
	for _, snake := range snapshot.Snakes() {
		for i, bodyPart := range snake.Body() {
			if i == len(snake.Body())-1 {
				continue // Skip tail
			}
			if bodyPart == p {
				return false
			}
		}
	}

	// Check if the point is a hazard
	for _, hazard := range snapshot.Hazards() {
		if hazard == p {
			return false
		}
	}

	return true
}