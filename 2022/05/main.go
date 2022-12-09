package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

type Move struct {
	n    int
	from int
	to   int
}

func loadStacks(filename string) []Stack {
	readFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lines := []string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			break
		}

		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	nStacks := (len(lines[len(lines)-2]) + 1) / 4
	stacks := make([]Stack, nStacks)

	for i := len(lines) - 2; i >= 0; i-- {
		for j := 0; j < nStacks; j++ {
			offset := j*4 + 1
			crate := lines[i][offset : offset+1]
			if crate != " " {
				stacks[j].Push(string(crate))
			}
		}
	}

	return stacks
}

func toInt(raw string) int {
	section, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return section
}

func getMoves(filename string, c chan Move) {
	readFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	skip := true

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if !skip {
			lineSplit := strings.Split(line, " ")

			c <- Move{n: toInt(lineSplit[1]), from: toInt(lineSplit[3]) - 1, to: toInt(lineSplit[5]) - 1}
		}

		if len(line) == 0 {
			skip = false
		}
	}

	readFile.Close()
	close(c)
}

func getTopContainers(stacks []Stack) string {
	top := ""

	for _, stack := range stacks {
		container, ok := stack.Pop()

		if ok {
			top += container
		}
	}

	return top
}

func reorderCrates(filename string, stacks []Stack, magnetic bool) string {
	c := make(chan Move)

	go getMoves(filename, c)

	for move := range c {
		if magnetic {
			temp := Stack{}
			for i := 0; i < move.n; i++ {
				container, _ := stacks[move.from].Pop()

				temp.Push(container)
			}

			for i := 0; i < move.n; i++ {
				container, _ := temp.Pop()

				stacks[move.to].Push(container)
			}
		} else {
			for i := 0; i < move.n; i++ {
				container, _ := stacks[move.from].Pop()

				stacks[move.to].Push(container)
			}
		}
	}

	return getTopContainers(stacks)
}

func main() {
	input := os.Args[1]

	stacks := loadStacks(input)
	moves := reorderCrates(input, stacks, false)
	fmt.Println(moves)

	stacks = loadStacks(input)
	moves = reorderCrates(input, stacks, true)
	fmt.Println(moves)
}
