package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const autoplayDelay = 150 * time.Millisecond

func autoplayTickCmd(index int) tea.Cmd {
	return tea.Tick(autoplayDelay, func(time.Time) tea.Msg {
		return autoplayTickMsg{index: index}
	})
}
