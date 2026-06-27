package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fkvivid/puzzle8/internal/board"
	"github.com/fkvivid/puzzle8/internal/game"
	"github.com/fkvivid/puzzle8/internal/solver"
)

type model struct {
	game   *game.Game
	screen screen
	width  int
	height int
	moves  int

	solving     bool
	solveAlgo   solver.Algorithm
	solveResult solver.Result

	autoplay      bool
	autoplayMoves []board.Dir

	lastMoved int
	elapsed   time.Duration
	started   time.Time

	// command mode
	cmdMode  bool
	cmdBuf   string
	cmdError string

	// hint / show-moves
	hintMoves   []board.Dir
	hintSolving bool
}

func NewModel() tea.Model {
	return model{
		game:    game.New(),
		screen:  screenMenu,
		started: time.Now(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg{} })
}

func solveCmd(b board.Board, algo solver.Algorithm, autoplay bool) tea.Cmd {
	return func() tea.Msg {
		result := solver.Solve(b, algo)
		return solveDoneMsg{result: result, algo: algo, autoplay: autoplay}
	}
}

// hintSolveCmd solves with the given algorithm and returns up to n moves (0 = all).
func hintSolveCmd(b board.Board, algo solver.Algorithm, n int) tea.Cmd {
	return func() tea.Msg {
		result := solver.Solve(b, algo)
		if !result.Found {
			return hintDoneMsg{moves: nil}
		}
		moves := result.Moves
		if n > 0 && n < len(moves) {
			moves = moves[:n]
		}
		return hintDoneMsg{moves: moves}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		if m.cmdMode {
			return m.handleCmdKey(msg)
		}
		if msg.String() == ":" && !m.autoplay {
			m.cmdMode = true
			m.cmdBuf = ""
			m.cmdError = ""
			return m, nil
		}
		if mapped := mapKey(msg, m.screen, m.solving); mapped != nil {
			return m.apply(mapped)
		}

	case solveDoneMsg:
		m.solving = false
		m.solveResult = msg.result
		m.solveAlgo = msg.algo
		if msg.autoplay && msg.result.Found {
			return m.startAutoplay(msg.result.Moves)
		}

	case hintDoneMsg:
		m.hintSolving = false
		if len(msg.moves) == 0 {
			m.cmdError = "no solution found from current position"
		} else {
			m.hintMoves = msg.moves
			m.cmdError = ""
		}

	case autoplayTickMsg:
		if msg.index >= len(m.autoplayMoves) {
			m.autoplay = false
			m.screen = screenWin
			return m, nil
		}
		before := m.game.Board
		m.game.Move(m.autoplayMoves[msg.index])
		m.moves++
		m.lastMoved = movedTileIndex(before, m.game.Board)
		return m, autoplayTickCmd(msg.index + 1)

	case tickMsg:
		m.elapsed = time.Since(m.started)
		return m, tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg{} })
	}

	return m, nil
}

func (m model) handleCmdKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "ctrl+c":
		m.cmdMode = false
		m.cmdBuf = ""
		m.cmdError = ""
	case "enter":
		cmd := m.cmdBuf
		m.cmdMode = false
		m.cmdBuf = ""
		return m.execCommand(cmd)
	case "backspace":
		if len(m.cmdBuf) > 0 {
			m.cmdBuf = m.cmdBuf[:len(m.cmdBuf)-1]
		}
	default:
		s := msg.String()
		if len(s) == 1 && s[0] >= ' ' {
			m.cmdBuf += s
		}
	}
	return m, nil
}

func (m model) apply(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case moveMsg:
		if m.autoplay {
			return m, nil
		}
		before := m.game.Board
		if m.game.Move(msg.dir) {
			m.moves++
			m.lastMoved = movedTileIndex(before, m.game.Board)

			// advance or clear hints
			if len(m.hintMoves) > 0 {
				if m.hintMoves[0] == msg.dir {
					m.hintMoves = m.hintMoves[1:]
				} else {
					m.hintMoves = nil
				}
			}

			if m.game.IsWon() {
				m.screen = screenWin
			}
		}

	case newGameMsg:
		m.game.Reset()
		m.moves = 0
		m.lastMoved = -1
		m.autoplay = false
		m.autoplayMoves = nil
		m.hintMoves = nil
		m.hintSolving = false
		m.cmdError = ""
		m.screen = screenPlay
		m.started = time.Now()

	case startPlayMsg:
		m.screen = screenPlay
		m.started = time.Now()

	case quitMsg:
		return m, tea.Quit

	case backMsg:
		m.screen = screenPlay
		m.solving = false

	case solveRequestMsg:
		m.screen = screenSolve
		m.solving = true
		m.solveAlgo = msg.algo
		m.lastMoved = -1
		m.hintMoves = nil
		m.hintSolving = false
		return m, solveCmd(m.game.Board, msg.algo, msg.autoplay)

	case autoplayStartMsg:
		if m.solveResult.Found {
			return m.startAutoplay(m.solveResult.Moves)
		}
	}

	return m, nil
}

func (m model) startAutoplay(moves []board.Dir) (tea.Model, tea.Cmd) {
	m.autoplay = true
	m.autoplayMoves = moves
	m.moves = 0
	m.hintMoves = nil
	m.screen = screenPlay
	m.lastMoved = -1
	m.started = time.Now()
	return m, autoplayTickCmd(0)
}

func movedTileIndex(before, after board.Board) int {
	for i := range before {
		if before[i] != 0 && before[i] != after[i] {
			for j := range after {
				if after[j] == before[i] {
					return j
				}
			}
		}
	}
	return -1
}

func (m model) View() string {
	return m.render()
}
