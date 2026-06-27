package tui

import "strings"

const cellWidth = 7
const cellHeight = 7

// 3-wide × 5-tall block-pixel glyphs using full-block █
var pixelDigits = map[int][5]string{
	1: {" █ ", " █ ", " █ ", " █ ", " █ "},
	2: {"███", "  █", "███", "█  ", "███"},
	3: {"███", "  █", "███", "  █", "███"},
	4: {"█ █", "█ █", "███", "  █", "  █"},
	5: {"███", "█  ", "███", "  █", "███"},
	6: {"███", "█  ", "███", "█ █", "███"},
	7: {"███", "  █", " █ ", "█  ", "█  "},
	8: {"███", "█ █", "███", "█ █", "███"},
}

func digitLines(val int) []string {
	lines := make([]string, cellHeight)
	if val == 0 {
		for i := range lines {
			lines[i] = strings.Repeat(" ", cellWidth)
		}
		return lines
	}
	px, ok := pixelDigits[val]
	if !ok {
		for i := range lines {
			lines[i] = strings.Repeat(" ", cellWidth)
		}
		return lines
	}
	lines[0] = strings.Repeat(" ", cellWidth)
	lines[cellHeight-1] = strings.Repeat(" ", cellWidth)
	for i, row := range px {
		lines[1+i] = padCenter(row, cellWidth)
	}
	return lines
}

func blankLines() []string {
	lines := make([]string, cellHeight)
	for i := range lines {
		lines[i] = strings.Repeat(" ", cellWidth)
	}
	return lines
}

func padCenter(s string, width int) string {
	runes := []rune(s)
	rlen := len(runes)
	if rlen >= width {
		return string(runes[:width])
	}
	left := (width - rlen) / 2
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", width-rlen-left)
}
