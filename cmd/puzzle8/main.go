package main

import (
	"fmt"
	"os"

	"github.com/fkvivid/puzzle8/internal/game"
	"github.com/fkvivid/puzzle8/internal/solver"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  eighth-puzzle")
	fmt.Println("  eighth-puzzle play")
	fmt.Println("  eighth-puzzle solve [--bfs|--astar|--greedy] [--auto]")
	fmt.Println("  eighth-puzzle help")
}

func parseAlgorithm(args []string) solver.Algorithm {
	algo := solver.BFS
	for _, arg := range args {
		switch arg {
		case "--bfs":
			algo = solver.BFS
		case "--astar":
			algo = solver.AStar
		case "--greedy":
			algo = solver.Greedy
		}
	}
	return algo
}

func hasFlag(args []string, flag string) bool {
	for _, a := range args {
		if a == flag {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) == 1 {
		game.New().Run()
		return
	}

	switch os.Args[1] {
	case "play":
		game.New().Run()

	case "solve":
		args := os.Args[2:]
		algo := parseAlgorithm(args)
		auto := hasFlag(args, "--auto")

		g := game.New()
		result := solver.Solve(g.Board, algo)
		solver.PrintResult(result, algo)

		if auto && result.Found {
			g.PlayMoves(result.Moves)
		}
		if !result.Found {
			os.Exit(1)
		}

	case "help", "--help", "-h":
		printHelp()

	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println()
		printHelp()
		os.Exit(1)
	}
}
