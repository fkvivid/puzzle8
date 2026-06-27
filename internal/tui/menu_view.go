package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func renderMenu() string {
	titleBox := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("214")).
		Padding(0, 6).
		Render(
			titleStyle.Render("P U Z Z L E   8"),
		)

	subtitle := subtitleStyle.Render("slide the tiles into order")

	rows := []struct{ key, desc string }{
		{"Enter", "start playing"},
		{"N", "new shuffle"},
		{"", ""},
		{"1 / 2 / 3", "solve BFS / A★ / Greedy (solve screen)"},
		{"4", "auto-solve with A★"},
		{"", ""},
		{":", "command mode"},
		{"  :solve", "show full solution path to follow"},
		{"  :bfs", "show BFS solution path"},
		{"  :greedy", "show Greedy solution path"},
		{"  :show-moves N", "show next N hint moves"},
		{"", ""},
		{"Q", "quit"},
	}

	lines := make([]string, len(rows))
	for i, r := range rows {
		if r.key == "" {
			lines[i] = ""
			continue
		}
		k := menuKeyStyle.Render(padRight(r.key, 7))
		d := menuDescStyle.Render(r.desc)
		lines[i] = "  " + k + "  " + d
	}

	menuContent := lipgloss.JoinVertical(lipgloss.Left, lines...)
	menuBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(0, 1).
		Render(menuContent)

	return lipgloss.JoinVertical(lipgloss.Center,
		titleBox,
		subtitle,
		"",
		menuBox,
	)
}
