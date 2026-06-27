# puzzle8

A terminal sliding-puzzle game written in Go. Slide the numbered tiles into order using keyboard controls, or let the built-in solvers compute a path for you to follow step by step.

```
╔═══════════════════════════════╗
║       P U Z Z L E   8        ║
╚═══════════════════════════════╝

  ┌───────┬───────┬───────┐
  │  ███  │  █ █  │       │
  │  █ █  │  █ █  │       │
  │  ███  │  ███  │       │
  │  █ █  │    █  │       │
  │  ███  │    █  │       │
  ├───────┼───────┼───────┤
  │   █   │  ███  │  ███  │
  │   █   │    █  │  █    │
  │   █   │  ███  │  ███  │
  │   █   │    █  │    █  │
  │   █   │  ███  │  ███  │
  ├───────┼───────┼───────┤
  │  ███  │  ███  │  ███  │
  │  █    │  █ █  │    █  │
  │  ███  │  ███  │  ███  │
  │  █ █  │  █ █  │    █  │
  │  ███  │  ███  │  ███  │
  └───────┴───────┴───────┘
```

## Install

### Homebrew (macOS / Linux)

Requires [Homebrew](https://brew.sh):

```sh
brew tap fkvivid/tap
brew install puzzle8
```

If Homebrew asks you to trust the tap first:

```sh
brew trust fkvivid/tap
brew install puzzle8
```

Then run `puzzle8` from any terminal.

### Go install

Requires Go on your PATH (`$(go env GOPATH)/bin` must be in `$PATH`):

```sh
go install github.com/fkvivid/puzzle8/cmd/puzzle8@latest
```

### Build from source

```sh
git clone https://github.com/fkvivid/puzzle8
cd puzzle8
go build -o puzzle8 ./cmd/puzzle8
./puzzle8
```

## Usage

```sh
puzzle8          # launch TUI
puzzle8 solve    # solve a random board and print moves
puzzle8 help     # show CLI options
```

### CLI flags for `solve`

| Flag       | Description              |
|------------|--------------------------|
| `--bfs`    | Use BFS (default)        |
| `--astar`  | Use A*                   |
| `--greedy` | Use Greedy best-first    |
| `--auto`   | Print board at each step |

## Controls

### In-game keys

| Key | Action |
|-----|--------|
| `↑ ↓ ← →` or `w a s d` or `h j k l` | Slide a tile |
| `N` | Shuffle and start a new game |
| `1` | Open solve screen — BFS |
| `2` | Open solve screen — A* |
| `3` | Open solve screen — Greedy |
| `4` | Auto-solve with A* (animated) |
| `Q` / `Ctrl+C` | Quit |

### Command mode

Press `:` at any time to open the command bar (like Neovim).  
Type a command and press `Enter` to run it, or `Esc` to cancel.

| Command | Description |
|---------|-------------|
| `:solve` | Compute the optimal A* solution and show it as a followable path |
| `:bfs` | Same using BFS |
| `:greedy` | Same using Greedy best-first |
| `:show-moves N` | Show only the next N hint moves (e.g. `:show-moves 5`) |
| `:show-moves-N` | Shorthand: `:show-moves-3` |
| `:autoplay` | Auto-play the last computed solution |
| `:new` / `:n` | New shuffled game |
| `:quit` / `:q` | Quit |
| `:help` / `:h` | Print command reference in the status bar |

### Following a solution path

After any solve command, the status bar shows the full path as arrows:

```
path (22) → ↓ ↓ → ↑ ↑ ← ↓ ↓ ← …+12  (tile 7 right)
```

- The number in parentheses is how many moves remain.
- The arrows are the next moves to make (up to 10 shown; `…+N` means N more follow).
- `(tile 7 right)` tells you which tile to move and in which direction.
- The tile that should move next is highlighted in **purple** on the board.
- Each correct move pops the front of the path automatically.
- Making a wrong move clears the path — run the command again to re-compute.

## Tile colours

| Colour | Meaning |
|--------|---------|
| Green background | Tile is in its correct goal position |
| Dark background | Tile is out of place |
| Orange background | Tile that was just moved |
| Purple background | Tile the next hint move will slide |

## Algorithms

| Algorithm | Guarantee | Notes |
|-----------|-----------|-------|
| BFS | Optimal (fewest moves) | Explores states level by level |
| A* | Optimal | Uses Manhattan-distance heuristic; much faster than BFS |
| Greedy | Not optimal | Fast but may produce longer paths |

All three algorithms operate on the standard 8-puzzle (3×3 grid, tiles 1–8, blank at position 9). The A* solver typically expands fewer than 5 000 states for any solvable board.

## License

[MIT](LICENSE)
