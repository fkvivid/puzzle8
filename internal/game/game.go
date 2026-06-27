package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fkvivid/puzzle8/internal/board"
	"github.com/fkvivid/puzzle8/internal/solver"
	"github.com/fkvivid/puzzle8/internal/terminal"
)

type Game struct {
	Board board.Board
}

func New() *Game {
	g := &Game{}
	g.Randomize()
	return g
}

func (g *Game) Randomize() {
	g.Board = board.Goal
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	moves := board.Deltas
	opposite := [4]int{1, 0, 3, 2}
	lastDir := -1

	for i := 0; i < 100; i++ {
		dir := rng.Intn(4)
		if dir == lastDir {
			dir = opposite[dir]
		}
		d := moves[dir]
		if next, ok := board.ApplyMove(g.Board, d[0], d[1]); ok {
			g.Board = next
			lastDir = opposite[dir]
		}
	}
}

func (g *Game) TryMove(drow, dcol int) bool {
	next, ok := board.ApplyMove(g.Board, drow, dcol)
	if !ok {
		return false
	}
	g.Board = next
	return true
}

func (g *Game) HandleKey(k terminal.Key) {
	switch k {
	case terminal.KeyUp:
		g.TryMove(-1, 0)
	case terminal.KeyDown:
		g.TryMove(1, 0)
	case terminal.KeyLeft:
		g.TryMove(0, -1)
	case terminal.KeyRight:
		g.TryMove(0, 1)
	}
}

func (g *Game) PlayMoves(moves []board.Dir) {
	for _, m := range moves {
		terminal.ClearScreen()
		fmt.Println(g.Board.String())
		fmt.Printf(">> %s\n", m)
		switch m {
		case board.Up:
			g.TryMove(-1, 0)
		case board.Down:
			g.TryMove(1, 0)
		case board.Left:
			g.TryMove(0, -1)
		case board.Right:
			g.TryMove(0, 1)
		}
		time.Sleep(400 * time.Millisecond)
	}
}

func (g *Game) runSolver(algo solver.Algorithm, autoPlay bool) {
	result := solver.Solve(g.Board, algo)
	solver.PrintResult(result, algo)

	if autoPlay && result.Found {
		fmt.Println("\nAuto-playing...")
		terminal.WaitEnter()
		g.PlayMoves(result.Moves)
	}
}

func (g *Game) handleCommand(line string) {
	switch line {
	case "help":
		fmt.Println("Commands:")
		fmt.Println("  solve-bfs")
		fmt.Println("  solve-astar")
		fmt.Println("  solve-greedy")
		fmt.Println("  auto-bfs")
		fmt.Println("  auto-astar")
		fmt.Println("  auto-greedy")
		fmt.Println("  help")
	case "solve-bfs":
		g.runSolver(solver.BFS, false)
	case "solve-astar":
		g.runSolver(solver.AStar, false)
	case "solve-greedy":
		g.runSolver(solver.Greedy, false)
	case "auto-bfs":
		g.runSolver(solver.BFS, true)
	case "auto-astar":
		g.runSolver(solver.AStar, true)
	case "auto-greedy":
		g.runSolver(solver.Greedy, true)
	default:
		fmt.Println("Unknown command. Type :help")
	}
}

func (g *Game) Run() {
	for {
		terminal.ClearScreen()
		fmt.Println(g.Board.String())
		fmt.Println("WASD/arrows move | : commands | q quit")

		key, err := terminal.ReadKey()
		if err != nil {
			fmt.Println("input error:", err)
			return
		}

		switch key {
		case terminal.KeyQuit:
			return
		case terminal.KeyCommand:
			fmt.Print(": ")
			line, err := terminal.ReadLine()
			if err != nil {
				return
			}
			g.handleCommand(line)
			terminal.WaitEnter()
		default:
			g.HandleKey(key)
		}
	}
}
