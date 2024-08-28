package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/Battle-Bunker/cyphid-snake/server"
	"github.com/BattlesnakeOfficial/rules/client"
)

func main() {

	metadata := client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "zuthan",
		Color:      "#FF0000",
		Head:       "evil",
		Tail:       "nr-booster",
	}

	portfolio := agent.NewPortfolio(
		agent.NewHeuristic(10.0, "alive", HeuristicAliveAllies),
		agent.NewHeuristic(1.0, "team-health", HeuristicHealth),
		agent.NewHeuristic(10.0, "food-proximity", HeuristicFoodProximity),
		agent.NewHeuristic(1.0, "length", HeuristicRelativeLength),
		agent.NewHeuristic(5.0, "floodfill", HeuristicFloodFill),
	)

	snakeAgent := agent.NewSnakeAgentWithTemp(portfolio, 5.0, metadata)
	server := server.NewServer(snakeAgent)

	server.Start()
}
