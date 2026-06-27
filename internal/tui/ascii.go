package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// boardTopLine: ┌──────┬──────┬──────┐
func boardTopLine(cols int) string {
	seg := strings.Repeat("─", cellWidth)
	parts := make([]string, cols)
	for i := range parts {
		parts[i] = seg
	}
	return "┌" + strings.Join(parts, "┬") + "┐"
}

// boardMidLine: ├──────┼──────┼──────┤
func boardMidLine(cols int) string {
	seg := strings.Repeat("─", cellWidth)
	parts := make([]string, cols)
	for i := range parts {
		parts[i] = seg
	}
	return "├" + strings.Join(parts, "┼") + "┤"
}

// boardBotLine: └──────┴──────┴──────┘
func boardBotLine(cols int) string {
	seg := strings.Repeat("─", cellWidth)
	parts := make([]string, cols)
	for i := range parts {
		parts[i] = seg
	}
	return "└" + strings.Join(parts, "┴") + "┘"
}

func asciiDivider(width int) string {
	return strings.Repeat("─", width)
}

func joinBoardLines(lines []string) string {
	return strings.Join(lines, "\n")
}

func centerBlock(content string, width int) string {
	if width <= 0 {
		return content
	}
	return lipgloss.Place(width, lipgloss.Height(content), lipgloss.Center, lipgloss.Center, content)
}

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
