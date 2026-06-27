package tui

import (
	"github.com/fkvivid/puzzle8/internal/board"
	"github.com/fkvivid/puzzle8/internal/solver"
)

type screen int

const (
	screenMenu screen = iota
	screenPlay
	screenSolve
	screenWin
)

type moveMsg struct{ dir board.Dir }
type newGameMsg struct{}
type startPlayMsg struct{}
type quitMsg struct{}
type backMsg struct{}

type solveRequestMsg struct {
	algo     solver.Algorithm
	autoplay bool
}

type solveDoneMsg struct {
	result   solver.Result
	algo     solver.Algorithm
	autoplay bool
}

type autoplayStartMsg struct{}
type autoplayTickMsg struct{ index int }

type tickMsg struct{}

type hintDoneMsg struct {
	moves []board.Dir
}
