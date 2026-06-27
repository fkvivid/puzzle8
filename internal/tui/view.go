package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/fkvivid/puzzle8/internal/board"
)

func (m model) render() string {
	if m.width == 0 {
		return "Loading..."
	}

	var body string
	switch m.screen {
	case screenMenu:
		body = lipgloss.Place(m.width, m.height-6, lipgloss.Center, lipgloss.Center, renderMenu())

	case screenPlay, screenWin:
		hint := nextHintTile(m.game.Board, m.hintMoves)
		boardContent := renderBoard(m.game.Board, m.lastMoved, hint)
		framed := boardFrameStyle.Render(boardContent)
		if m.screen == screenWin {
			body = overlayWin(framed, m.moves, m.width, m.height-6)
		} else {
			body = lipgloss.Place(m.width, m.height-6, lipgloss.Center, lipgloss.Center, framed)
		}

	case screenSolve:
		boardContent := renderBoard(m.game.Board, m.lastMoved, -1)
		framed := boardFrameStyle.Render(boardContent)
		centered := lipgloss.Place(m.width, m.height-10, lipgloss.Center, lipgloss.Center, framed)
		body = lipgloss.JoinVertical(lipgloss.Left, centered, renderSolve(m))
	}

	footer := m.renderFooter()
	divider := asciiFrameStyle.Render(asciiDivider(min(m.width, 48)))
	return lipgloss.JoinVertical(lipgloss.Left, body, divider, footer)
}

func (m model) renderFooter() string {
	var lines []string

	// status line
	var status string
	switch m.screen {
	case screenMenu:
		status = "  press Enter to play"
	case screenPlay:
		if m.autoplay {
			status = fmt.Sprintf("  auto-playing…  move %d", m.moves)
		} else {
			status = fmt.Sprintf("  moves: %d", m.moves)
		}
	case screenWin:
		status = fmt.Sprintf("  solved in %d moves", m.moves)
	case screenSolve:
		if m.solving {
			status = "  computing…"
		} else if m.solveResult.Found {
			status = fmt.Sprintf("  %d-move solution ready", len(m.solveResult.Moves))
		} else {
			status = "  no solution found"
		}
	}
	lines = append(lines, statusStyle.Render(status))

	// hints bar
	if m.hintSolving {
		lines = append(lines, hintStyle.Render("  hints: computing…"))
	} else if len(m.hintMoves) > 0 {
		lines = append(lines, renderHintsBar(m.game.Board, m.hintMoves))
	}

	// command / error bar
	if m.cmdMode {
		lines = append(lines, cmdStyle.Render("  :"+m.cmdBuf+"▌"))
	} else if m.cmdError != "" {
		lines = append(lines, cmdErrorStyle.Render("  "+m.cmdError))
	} else {
		lines = append(lines, helpStyle.Render(m.renderHelp()))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

const maxVisibleHints = 10

func dirArrow(d board.Dir) (arrow, name string) {
	switch d {
	case board.Up:
		return "↑", "up"
	case board.Down:
		return "↓", "down"
	case board.Left:
		return "←", "left"
	case board.Right:
		return "→", "right"
	}
	return "?", "?"
}

func renderHintsBar(b board.Board, hints []board.Dir) string {
	total := len(hints)

	// how many to show inline
	show := hints
	more := ""
	if total > maxVisibleHints {
		show = hints[:maxVisibleHints]
		more = fmt.Sprintf(" …+%d", total-maxVisibleHints)
	}

	arrows := make([]string, len(show))
	firstArrow, firstName := dirArrow(hints[0])
	arrows[0] = firstArrow
	for i := 1; i < len(show); i++ {
		a, _ := dirArrow(show[i])
		arrows[i] = a
	}

	// which tile moves next
	var nextDesc string
	nextTile := nextHintTile(b, hints)
	if nextTile >= 0 {
		nextDesc = fmt.Sprintf("  (tile %d %s)", b[nextTile], firstName)
	}

	count := fmt.Sprintf("(%d) ", total)
	return hintStyle.Render(fmt.Sprintf("  path %s%s%s%s", count, strings.Join(arrows, " "), more, nextDesc))
}

func (m model) renderHelp() string {
	switch m.screen {
	case screenMenu:
		return "  arrows/wasd · n new · : commands · q quit"
	case screenPlay:
		return "  arrows/wasd move · n new · 1-4 solve · : commands · q quit"
	case screenWin:
		return "  n new game · q quit"
	case screenSolve:
		if m.solving {
			return "  please wait…"
		}
		return "  Enter auto-play · Esc back · q quit"
	default:
		return ""
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
