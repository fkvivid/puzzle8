package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/fkvivid/puzzle8/internal/board"
)

func renderBoard(b board.Board, lastMoved int, hintTile int) string {
	var lines []string
	lines = append(lines, boardTopLine(board.GridDim))

	for row := 0; row < board.GridDim; row++ {
		var cells [][]string
		for col := 0; col < board.GridDim; col++ {
			i := row*board.GridDim + col
			cells = append(cells, renderCell(b[i], i, lastMoved, hintTile))
		}
		lines = append(lines, boardDataRows(cells)...)
		if row < board.GridDim-1 {
			lines = append(lines, boardMidLine(board.GridDim))
		}
	}

	lines = append(lines, boardBotLine(board.GridDim))
	return joinBoardLines(lines)
}

func renderCell(val int, index int, lastMoved int, hintTile int) []string {
	style := blankStyle
	content := digitLines(val)

	if val != 0 {
		style = tileWrongStyle
		if board.Goal[index] == val {
			style = tileCorrectStyle
		}
		if index == hintTile {
			style = hintTileStyle
		}
		if index == lastMoved {
			style = tileHighlightStyle
		}
	}

	styled := make([]string, cellHeight)
	for i := 0; i < cellHeight; i++ {
		styled[i] = style.Width(cellWidth).Align(lipgloss.Center).Render(content[i])
	}
	return styled
}

func boardDataRows(cells [][]string) []string {
	rows := make([]string, cellHeight)
	for h := 0; h < cellHeight; h++ {
		parts := make([]string, len(cells))
		for i, cell := range cells {
			parts[i] = cell[h]
		}
		rows[h] = "│" + strings.Join(parts, "│") + "│"
	}
	return rows
}

// nextHintTile returns the index of the tile that will move given the next hint direction.
func nextHintTile(b board.Board, hints []board.Dir) int {
	if len(hints) == 0 {
		return -1
	}
	blank := board.FindBlank(b)
	var drow, dcol int
	switch hints[0] {
	case board.Up:
		drow, dcol = -1, 0
	case board.Down:
		drow, dcol = 1, 0
	case board.Left:
		drow, dcol = 0, -1
	case board.Right:
		drow, dcol = 0, 1
	}
	row := blank/board.GridDim + drow
	col := blank%board.GridDim + dcol
	if row < 0 || row >= board.GridDim || col < 0 || col >= board.GridDim {
		return -1
	}
	return row*board.GridDim + col
}
