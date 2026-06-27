package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type Key int

const (
	KeyNone Key = iota
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyQuit
	KeyCommand
)

func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

func ReadKey() (Key, error) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)

	if err != nil {
		return KeyNone, err
	}

	defer term.Restore(fd, oldState)

	buf := make([]byte, 3)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		return KeyNone, err
	}
	switch {
	case n >= 1 && (buf[0] == 'q' || buf[0] == 'Q'):
		return KeyQuit, nil
	case n >= 1 && buf[0] == ':':
		return KeyCommand, nil
	case n >= 1 && (buf[0] == 'w' || buf[0] == 'W'):
		return KeyUp, nil
	case n >= 1 && (buf[0] == 's' || buf[0] == 'S'):
		return KeyDown, nil
	case n >= 1 && (buf[0] == 'a' || buf[0] == 'A'):
		return KeyLeft, nil
	case n >= 1 && (buf[0] == 'd' || buf[0] == 'D'):
		return KeyRight, nil
	case n >= 3 && buf[0] == 27 && buf[1] == '[':
		switch buf[2] {
		case 'A':
			return KeyUp, nil
		case 'B':
			return KeyDown, nil
		case 'C':
			return KeyRight, nil
		case 'D':
			return KeyLeft, nil
		}
	}
	return KeyNone, nil
}

func ReadLine() (string, error) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return "", err
	}
	term.Restore(fd, oldState)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	return strings.TrimSpace(line), err
}
func WaitEnter() {
	fmt.Print("Press Enter to continue...")
	_, _ = ReadLine()
}
