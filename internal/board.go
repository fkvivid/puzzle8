package board

import "fmt"

const GridDim = 3

type Board [9]int

var Goal = Board{1, 2, 3, 4, 5, 6, 7, 8, 0}

type Dir int

const (
	Up Dir = iota
	Down
	Left
	Right
)

func (d Dir) String() string {
	switch d {
	case Up:
		return "UP"
	case Down:
		return "DOWN"
	case Left:
		return "LEFT"
	case Right:
		return "RIGHT"
	default:
		return "?"
	}
}

func FindBlank(b Board) int {
	for i, v := range b {
		if v == 0 {
			return i
		}
	}
	return -1
}

func IsSolved(b Board) bool {
	return b == Goal
}

func (b Board) Key() string {
	s := make([]byte, 9)
	for i, v := range b {
		s[i] = byte('0' + v)
	}
	return string(s)
}

func ApplyMove(b Board, drow, dcol int) (Board, bool) {
	blank := FindBlank(b)
	row := blank/GridDim + drow
	col := blank%GridDim + dcol

	if row < 0 || row >= GridDim || col < 0 || col >= GridDim {
		return b, false
	}

	next := b
	target := row*GridDim + col
	next[blank], next[target] = next[target], next[blank]
	return next, true
}

func DirFromDelta(drow, dcol int) Dir {
	switch {
	case drow == -1:
		return Up
	case drow == 1:
		return Down
	case dcol == -1:
		return Left
	case dcol == 1:
		return Right
	default:
		return -1
	}
}

var Deltas = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func (b Board) String() string {
	out := ""
	for i, v := range b {
		if i > 0 && i%GridDim == 0 {
			out += "\n"
		}
		if v == 0 {
			out += " "
		} else {
			out += fmt.Sprintf("%d", v)
		}

	}

	return out
}
