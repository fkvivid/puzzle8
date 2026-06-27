package game

import (
	"math/rand"
	"time"

	"github.com/fkvivid/puzzle8/internal/board"
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

func (g *Game) Move(dir board.Dir) bool {
	switch dir {
	case board.Up:
		return g.TryMove(-1, 0)
	case board.Down:
		return g.TryMove(1, 0)
	case board.Left:
		return g.TryMove(0, -1)
	case board.Right:
		return g.TryMove(0, 1)
	default:
		return false
	}
}

func (g *Game) IsWon() bool {
	return board.IsSolved(g.Board)
}

func (g *Game) Reset() {
	g.Randomize()
}

func (g *Game) SetBoard(b board.Board) {
	g.Board = b
}
