package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Position struct {
	row int
	col int
}
type Matrix [][]rune

type Hill struct {
	start Position
	end   Position
	field Matrix
}

func getHill(filename string) Hill {
	hill := Hill{}

	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	rows := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		row := []rune{}

		for i, r := range line {
			if r == 'S' {
				hill.start, r = Position{rows, i}, 'a'
			} else if r == 'E' {
				hill.end, r = Position{rows, i}, 'z'
			}
			row = append(row, r)
		}

		hill.field = append(hill.field, row)
		rows++
	}

	return hill
}

type heuristicFunction func(from, to Position, m Matrix) int
type costFunction func(from, to Position, m Matrix) int
type neighborsFunction func(m Matrix, i, j int) []Position

func Path(start, to Position, m Matrix, neighbors neighborsFunction, cost costFunction, heuristic heuristicFunction) (path []Position, distance int) {
	frontier := &PriorityQueue{}
	heap.Init(frontier)
	heap.Push(frontier, &Node{pos: start, priority: 0})

	cameFrom := map[Position]Position{start: start}
	costSoFar := map[Position]int{start: 0}

	for {
		if frontier.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(frontier).(*Node).pos
		if current == to {
			// Found a path to the goal.
			path := []Position{}
			curr := current
			for curr != start {
				path = append(path, curr)
				curr = cameFrom[curr]
			}
			return path, costSoFar[to]
		}

		for _, neighbor := range neighbors(m, current.row, current.col) {
			newCost := costSoFar[current] + cost(current, neighbor, m)
			if _, ok := costSoFar[neighbor]; !ok || newCost < costSoFar[neighbor] {
				costSoFar[neighbor] = newCost
				priority := newCost + heuristic(neighbor, to, m)
				heap.Push(frontier, &Node{pos: neighbor, priority: priority})
				cameFrom[neighbor] = current
			}
		}
	}
}

type Node struct {
	pos      Position
	priority int
	index    int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func neighbors(m Matrix, i, j int) []Position {
	pos := []Position{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	res := []Position{}
	for _, p := range pos {
		if p.row >= 0 && p.row < len(m) && p.col >= 0 && p.col < len(m[0]) {
			src := m[i][j]
			dest := m[p.row][p.col]
			if dest-src <= 1 {
				res = append(res, p)
			}
		}
	}
	return res
}

func cost(from, to Position, m Matrix) int {
	return 1
}

func heuristic(from, to Position, m Matrix) int {
	return int(m[to.row][to.col] - m[from.row][from.col])
}

func search(v rune, m Matrix) Position {
	for j, l := range m {
		for i, c := range l {
			if c == v {
				return Position{i, j}
			}
		}
	}
	return Position{}
}

func neighbors2(m Matrix, i, j int) []Position {
	n := neighbors(m, i, j)
	if m[i][j] == 'a' {
		a := search('a', m)
		return append(n, a)
	}
	return n
}

func cost2(from, to Position, m Matrix) int {
	if m[from.row][from.col] == 'a' && m[to.row][to.col] == 'a' {
		return 0
	}
	return 1
}

func main() {
	input := os.Args[1]

	hill := getHill(input)
	_, cost := Path(hill.start, hill.end, hill.field, neighbors, cost, heuristic)
	fmt.Println("1:", cost)

	_, cost = Path(hill.start, hill.end, hill.field, neighbors2, cost2, heuristic)
	fmt.Println("2:", cost)
}
