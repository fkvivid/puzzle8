package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fkvivid/puzzle8/internal/board"
	"github.com/fkvivid/puzzle8/internal/solver"
)

func mapKey(msg tea.KeyMsg, scr screen, solving bool) tea.Msg {
	switch msg.String() {
	case "ctrl+c", "q":
		return quitMsg{}
	case "esc":
		if scr == screenSolve && !solving {
			return backMsg{}
		}
	case "enter":
		switch scr {
		case screenMenu:
			return startPlayMsg{}
		case screenSolve:
			if !solving {
				return autoplayStartMsg{}
			}
		}
	case "up", "k", "w":
		if scr == screenPlay {
			return moveMsg{board.Up}
		}
	case "down", "j", "s":
		if scr == screenPlay {
			return moveMsg{board.Down}
		}
	case "left", "h", "a":
		if scr == screenPlay {
			return moveMsg{board.Left}
		}
	case "right", "l", "d":
		if scr == screenPlay {
			return moveMsg{board.Right}
		}
	case "n":
		return newGameMsg{}
	case "1":
		return solveRequestMsg{solver.BFS, false}
	case "2":
		return solveRequestMsg{solver.AStar, false}
	case "3":
		return solveRequestMsg{solver.Greedy, false}
	case "4":
		return solveRequestMsg{solver.AStar, true}
	}
	return nil
}
