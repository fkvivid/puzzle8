package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func renderSolve(m model) string {
	if m.solving {
		return solveStyle.Render("  computing with " + m.solveAlgo.String() + "…")
	}

	if !m.solveResult.Found {
		return solveStyle.Render("  no solution found  ·  Esc to go back")
	}

	algo := solveStyle.Render(fmt.Sprintf("  algorithm  %s", m.solveAlgo))
	expanded := solveStyle.Render(fmt.Sprintf("  expanded   %d states", m.solveResult.Expanded))
	length := solveStyle.Render(fmt.Sprintf("  solution   %d moves", len(m.solveResult.Moves)))
	actions := helpStyle.Render("  Enter auto-play · Esc back")

	return lipgloss.JoinVertical(lipgloss.Left,
		solveStyle.Render("  solution found"),
		algo, expanded, length, "", actions,
	)
}

func overlayWin(boardContent string, moves int, width, height int) string {
	banner := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("214")).
		Padding(0, 3).
		Render(
			winStyle.Render(fmt.Sprintf("solved in %d moves!", moves)) +
				"\n" +
				helpStyle.Render("  N  new game    Q  quit"),
		)
	content := lipgloss.JoinVertical(lipgloss.Center, boardContent, "", banner)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}
