package tui

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fkvivid/puzzle8/internal/solver"
)

func (m model) execCommand(raw string) (tea.Model, tea.Cmd) {
	parts := strings.Fields(raw)
	if len(parts) == 0 {
		return m, nil
	}
	name := strings.ToLower(parts[0])

	switch name {
	case "q", "quit":
		return m, tea.Quit

	case "n", "new":
		return m.apply(newGameMsg{})

	case "h", "help":
		m.cmdError = "cmds: :solve :bfs :greedy  :show-moves N  :autoplay  :new  :q"
		return m, nil

	// solve commands: compute full path, show as followable hints
	case "solve", "astar":
		return m.startHintsAlgo(solver.AStar, 0)

	case "bfs":
		return m.startHintsAlgo(solver.BFS, 0)

	case "greedy":
		return m.startHintsAlgo(solver.Greedy, 0)

	case "autoplay", "auto":
		if m.solveResult.Found {
			return m.apply(autoplayStartMsg{})
		}
		return m.apply(solveRequestMsg{algo: solver.AStar, autoplay: true})

	case "show-moves":
		n, err := parseN(parts, 1)
		if err != nil {
			m.cmdError = "usage: :show-moves N"
			return m, nil
		}
		return m.startHintsAlgo(solver.AStar, n)
	}

	// :show-moves-N  (hyphenated variant)
	if strings.HasPrefix(name, "show-moves-") {
		nStr := strings.TrimPrefix(name, "show-moves-")
		n, err := strconv.Atoi(nStr)
		if err != nil || n <= 0 {
			m.cmdError = "usage: :show-moves-N  (N > 0)"
			return m, nil
		}
		return m.startHintsAlgo(solver.AStar, n)
	}

	m.cmdError = "unknown command — try :help"
	return m, nil
}

// startHintsAlgo solves with the given algorithm and shows results as followable hints.
// n=0 means show the full solution path.
func (m model) startHintsAlgo(algo solver.Algorithm, n int) (tea.Model, tea.Cmd) {
	if m.game.IsWon() {
		m.cmdError = "puzzle already solved"
		return m, nil
	}
	m.hintSolving = true
	m.hintMoves = nil
	m.cmdError = ""
	// stay on the play screen so the player can follow along
	if m.screen == screenSolve {
		m.screen = screenPlay
	}
	return m, hintSolveCmd(m.game.Board, algo, n)
}

func parseN(parts []string, idx int) (int, error) {
	if idx >= len(parts) {
		return 0, strconv.ErrSyntax
	}
	n, err := strconv.Atoi(parts[idx])
	if err != nil || n <= 0 {
		return 0, strconv.ErrSyntax
	}
	return n, nil
}
