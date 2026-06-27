package solver

import (
	"container/heap"
	"fmt"

	"github.com/fkvivid/puzzle8/internal/board"
)

type Algorithm int

const (
	BFS Algorithm = iota
	AStar
	Greedy
)

func (a Algorithm) String() string {
	switch a {
	case BFS:
		return "BFS"
	case AStar:
		return "A*"
	case Greedy:
		return "Greedy"
	default:
		return "?"
	}
}

type Result struct {
	Found    bool
	Moves    []board.Dir
	Expanded int
}

func Manhattan(b board.Board) int {
	h := 0
	for i, v := range b {
		if v == 0 {
			continue
		}

		goalI := v - 1
		h += abs(i/board.GridDim-goalI/board.GridDim) + abs(i%GridDim-goalI%GridDim)
	}

	return h
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type node struct {
	state board.Board
	g     int
	path  []board.Dir
	pri   int
	index int
}

type pq []*node

func (p pq) Len() int           { return len(p) }
func (p pq) Less(i, j int) bool { return p[i].pri > p[j].pri }
func (p pq) Swap(i, j int)      { p[i], p[j] = p[j], p[i]; p[i].index = i; p[j].index = j }
func (p *pq) Push(x any)        { n := x.(*node); n.index = len(*p); *p = append(*p, n) }
func (p *pq) Pop() any          { old := *p; n := old[len(old)-1]; *p = old[:len(old)-1]; return n }

func Solve(start board.Board, algo Algorithm) Result {
	if board.IsSolved(start) {
		return Result{Found: true, Expanded: 0}
	}

	switch algo {
	case BFS:
		return solveBFS(start)
	default:
		return solveBestFirst(start, algo)
	}
}

func solveBFS(start board.Board) Result {
	type qnode struct {
		state board.Board
		path  []board.Dir
	}

	queue := []qnode{{state: start, path: nil}}
	visited := map[string]bool{start.Key(): true}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, d := range board.Deltas {
			next, ok := board.ApplyMove(cur.state, d[0], d[1])

			if !ok {
				continue
			}

			key := next.Key()
			if visited[key] {
				continue
			}

			visited[key] = true

			path := append(append([]board.Dir{}, cur.path...), board.DirDirFromDelta(d[0], d[1]))
			if board.IsSolved(next) {
				return Result{Found: true, Moves: path, Expanded: len(visited)}
			}
			queue = append(queue, qnode{states: next, path: path})
		}
	}
}

func solveBestFirst(start board.Board, algo Algorithm) Result {
	visited := map[string]bool{}

	open := &pq{}
	heap.Init(open)

	startPri := 0

	if algo == Greedy {
		startPri = Manhattan(start)
	}
	heap.Push(open, &node{state: start, g: 0, path: nil, pri: startPri})

	for open.Len() > 0 {
		cur := heap.Pop(open).(*node)
		key := cur.state.Key()
		if visited[key] {
			continue
		}

		visited[key] = true

		if board.IsSolved(cur.state) {
			return Result{Found: true, Moves: cur.path, Expanded: len(visited)}
		}

		for _, d := range board.Deltas {
			next, ok := board.ApplyMove(cur.state, d[0], d[1])

			if !ok {
				continue
			}

			if visited[next.Key()] {
				continue
			}

			g := cur.g + 1
			pri := 0
			switch algo {
			case AStar:
				pri = g + Manhattan(next)
			case Greedy:
				pri = Manhattan(next)
			}

			path := append(append([]board.Dir{}, cur.path...), board.DirFromDelta(d[0], d[1]))
			heap.Push(open, &node{state: next, g: g, path: path, pri: pri})
		}
	}

	return Result{Expanded: len(visited)}
}

func PrintResult(r Result, algo Algorithm) {
	if !r.Found {
		fmt.Println("No solution found")
		return
	}
	fmt.Printf("Algorithm: %s\n", algo)
	fmt.Printf("Expanded states: %d\n", r.Expanded)
	fmt.Printf("Moves: %d\n", len(r.Moves))
	for i, m := range r.Moves {
		fmt.Printf("%d. %s\n", i+1, m)
	}
}
