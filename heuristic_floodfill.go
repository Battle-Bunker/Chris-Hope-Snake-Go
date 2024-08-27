package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
)

// HeuristicFloodFill calculates the sum of free spaces available to allied snakes
func HeuristicFloodFill(snapshot agent.GameSnapshot) float64 {
	totalFreeSpaces := 0.0
	for _, allySnake := range snapshot.YourTeam() {
		freeSpaces := floodFill(snapshot, allySnake.Head())
		totalFreeSpaces += float64(freeSpaces)
	}
	return totalFreeSpaces
}

// floodFill performs a flood fill algorithm to count free spaces
func floodFill(snapshot agent.GameSnapshot, start rules.Point) int {
	width, height := snapshot.Width(), snapshot.Height()
	visited := make(map[rules.Point]bool)
	queue := []rules.Point{start}
	count := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true
		count++

		// Check adjacent cells
		for _, dir := range []rules.Point{{X: 0, Y: 1}, {X: 0, Y: -1}, {X: 1, Y: 0}, {X: -1, Y: 0}} {
			next := rules.Point{X: current.X + dir.X, Y: current.Y + dir.Y}
			if isValidMove(snapshot, next, width, height) {
				queue = append(queue, next)
			}
		}
	}

	return count
}

// isValidMove checks if a move to the given point is valid
func isValidMove(snapshot agent.GameSnapshot, p rules.Point, width, height int) bool {
	// Check bounds
	if p.X < 0 || p.X >= width || p.Y < 0 || p.Y >= height {
		return false
	}

	// Check for collision with snakes
	for _, snake := range snapshot.Snakes() {
		for _, bodyPart := range snake.Body() {
			if p == bodyPart {
				return false
			}
		}
	}

	// Check for hazards (optional, depending on game rules)
	for _, hazard := range snapshot.Hazards() {
		if p == hazard {
			return false
		}
	}

	return true
}
