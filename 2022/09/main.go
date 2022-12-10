package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Bridge struct {
	nodes     []Position
	positions map[string]Position
}

func toInt(raw string) int {
	section, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return section
}

func move(pos Position, mov string, n int) Position {
	position := Position{pos.x, pos.y}
	if mov == "R" {
		position.x += n
	} else if mov == "L" {
		position.x -= n
	} else if mov == "U" {
		position.y += n
	} else if mov == "D" {
		position.y -= n
	} else {
		panic("Unknown move: " + mov)
	}
	return position
}

func moveRightOrLeft(tail Position, delta int) Position {
	if delta > 0 {
		return move(tail, "R", 1)
	} else if delta < 0 {
		return move(tail, "L", 1)
	}
	return tail
}

func moveUpOrDown(tail Position, delta int) Position {
	if delta > 0 {
		return move(tail, "U", 1)
	} else if delta < 0 {
		return move(tail, "D", 1)
	}
	return tail
}

func moveToNextPosition(head Position, tail Position) Position {
	xDelta, yDelta := head.x-tail.x, head.y-tail.y

	if xDelta*xDelta > 1 && yDelta == 0 {
		return moveRightOrLeft(tail, xDelta)
	} else if yDelta*yDelta > 1 && yDelta == 0 {
		return moveUpOrDown(tail, yDelta)
	} else if yDelta*yDelta > 1 && xDelta*xDelta > 1 || yDelta*yDelta > 1 || xDelta*xDelta > 1 {
		temp := moveRightOrLeft(tail, xDelta)
		return moveUpOrDown(temp, yDelta)
	}

	return tail
}

func getTailUniqueSpots(filename string, nodes int) int {
	bridge := Bridge{nodes: make([]Position, nodes), positions: make(map[string]Position)}

	for i := 0; i < nodes; i++ {
		bridge.nodes[i] = Position{0, 0}
	}

	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	bridge.positions[strconv.Itoa(bridge.nodes[0].x)+"-"+strconv.Itoa(bridge.nodes[0].y)] = bridge.nodes[0]

	for fileScanner.Scan() {
		line := fileScanner.Text()

		split := strings.Split(line, " ")
		mov, n := split[0], toInt(split[1])

		for i := 0; i < n; i++ {

			bridge.nodes[0] = move(bridge.nodes[0], mov, 1)

			for j := 1; j < nodes; j++ {
				tail := bridge.nodes[j]
				newPosition := moveToNextPosition(bridge.nodes[j-1], tail)

				if newPosition.x != tail.x || newPosition.y != tail.y {
					bridge.nodes[j] = newPosition

					if j == nodes-1 {
						bridge.positions[strconv.Itoa(newPosition.x)+"-"+strconv.Itoa(newPosition.y)] = tail
					}
				}
			}
		}

	}

	return len(bridge.positions)
}

func main() {
	input := os.Args[1]

	moves := getTailUniqueSpots(input, 2)
	fmt.Println(moves)

	moves = getTailUniqueSpots(input, 10)
	fmt.Println(moves)
}
