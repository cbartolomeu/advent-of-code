package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(raw string) int {
	res, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return res
}

func sum(arr []int) int {
	res := 0
	for _, it := range arr {
		res += it
	}
	return res
}

func processAdd(cycle int, r int, signals []int, v int, ctr []string) (int, int, []int, []string) {
	cycle, r, signals, ctr = processNoop(cycle, r, signals, ctr)

	c := cycle + 1

	pointer := c - 1
	if pointer%40 > r+1 || pointer%40 < r-1 {
		ctr[pointer/40] += "."
	} else {
		ctr[pointer/40] += "#"
	}

	if c == 20 || (c-20)%40 == 0 {
		signals[((c - 20) / 40)] = r * c
	}

	r += v

	return c, r, signals, ctr
}

func processNoop(cycle int, r int, signals []int, ctr []string) (int, int, []int, []string) {
	c := cycle + 1

	pointer := c - 1
	if pointer%40 > r+1 || pointer%40 < r-1 {
		ctr[pointer/40] += "."
	} else {
		ctr[pointer/40] += "#"
	}

	if c == 20 || (c-20)%40 == 0 {
		signals[((c - 20) / 40)] = r * c
	}

	return c, r, signals, ctr
}

func moveCycle(rawInstruction string, cycle int, r int, signals []int, ctr []string) (int, int, []int, []string) {
	instruction := strings.Split(rawInstruction, " ")

	if instruction[0] == "addx" {
		return processAdd(cycle, r, signals, toInt(instruction[1]), ctr)
	} else if instruction[0] == "noop" {
		return processNoop(cycle, r, signals, ctr)
	}

	panic("Unknown instruction")
}

func getStrengthSignals(filename string, maxCycle int) (int, int, []int, []string) {
	signals := make([]int, ((maxCycle-20)/40)+1)
	ctr := make([]string, ((maxCycle-20)/40)+1)

	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	r, cycle := 1, 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		cycle, r, signals, ctr = moveCycle(line, cycle, r, signals, ctr)

		if cycle >= maxCycle {
			return r, cycle, signals, ctr
		}

	}

	return r, cycle, signals, ctr
}

func main() {
	input := os.Args[1]
	r, cycle, signals, ctr := getStrengthSignals(input, 240)
	fmt.Println(r, cycle, signals, sum(signals))

	for _, line := range ctr {
		fmt.Println(line)
	}
}
