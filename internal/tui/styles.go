package tui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("214")).
			Bold(true)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	menuKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("214")).
			Bold(true)

	menuDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	asciiFrameStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	tileCorrectStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("255")).
				Background(lipgloss.Color("28")).
				Bold(true)

	tileWrongStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245")).
			Background(lipgloss.Color("236")).
			Bold(true)

	tileHighlightStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("255")).
				Background(lipgloss.Color("166")).
				Bold(true)

	hintTileStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			Background(lipgloss.Color("57")).
			Bold(true)

	blankStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("237")).
			Background(lipgloss.Color("234"))

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	winStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("214")).
			Bold(true)

	solveStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	boardFrameStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	hintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("147"))

	cmdStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255"))

	cmdErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("203"))
)
