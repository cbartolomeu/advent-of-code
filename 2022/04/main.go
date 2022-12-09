package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	min int
	max int
}

type Pair struct {
	a Assignment
	b Assignment
}

func createSection(raw string) int {
	section, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return section
}

func createAssignment(raw string) Assignment {
	split := strings.Split(raw, "-")

	return Assignment{min: createSection(split[0]), max: createSection(split[1])}
}

func getPairAssignment(filename string, c chan Pair) {
	readFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		pairSplit := strings.Split(line, ",")

		c <- Pair{a: createAssignment(pairSplit[0]), b: createAssignment(pairSplit[1])}
	}

	readFile.Close()
	close(c)
}

func isContainedBy(a Assignment, b Assignment) bool {
	return a.min >= b.min && a.max <= b.max
}

func isContainedAssignment(pair Pair) bool {
	return isContainedBy(pair.a, pair.b) || isContainedBy(pair.b, pair.a)
}

func overlapsWith(a Assignment, b Assignment) bool {
	return a.min >= b.min && a.min <= b.max || a.max >= b.min && a.max <= b.max
}

func isOverlapAssignment(pair Pair) bool {
	return overlapsWith(pair.a, pair.b) || overlapsWith(pair.b, pair.a)
}

func getCommonAssignments(filename string, isCommonFn func(pair Pair) bool) int {
	c := make(chan Pair)

	go getPairAssignment(filename, c)

	common := 0

	for pair := range c {
		if isCommonFn(pair) {
			common++
		}
	}

	return common
}

func main() {
	input := os.Args[1]

	common := getCommonAssignments(input, isContainedAssignment)
	fmt.Println(common)

	common = getCommonAssignments(input, isOverlapAssignment)
	fmt.Println(common)
}
