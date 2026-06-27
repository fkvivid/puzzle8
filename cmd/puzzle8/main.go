package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fkvivid/puzzle8/internal/game"
	"github.com/fkvivid/puzzle8/internal/solver"
	"github.com/fkvivid/puzzle8/internal/tui"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  puzzle8")
	fmt.Println("  puzzle8 play")
	fmt.Println("  puzzle8 solve [--bfs|--astar|--greedy] [--auto]")
	fmt.Println("  puzzle8 help")
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

func runTUI() {
	p := tea.NewProgram(tui.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) == 1 {
		runTUI()
		return
	}

	switch os.Args[1] {
	case "play":
		runTUI()

	case "solve":
		args := os.Args[2:]
		algo := parseAlgorithm(args)
		auto := hasFlag(args, "--auto")

		g := game.New()
		result := solver.Solve(g.Board, algo)
		solver.PrintResult(result, algo)

		if auto && result.Found {
			for _, move := range result.Moves {
				fmt.Println(g.Board)
				fmt.Printf(">> %s\n", move)
				g.Move(move)
				time.Sleep(400 * time.Millisecond)
			}
			fmt.Println(g.Board)
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
