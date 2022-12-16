package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	row int
	col int
}

type cave map[int](map[int]bool)

func toInt(raw string) int {
	res, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return res
}

func loadRocks(filename string) cave {
	cave := cave{}

	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineSplit := strings.Split(line, "->")

		start := position{-1, -1}
		for _, rock := range lineSplit {
			rockSplit := strings.Split(rock, ",")

			pos := position{toInt(strings.TrimSpace(rockSplit[0])), toInt(strings.TrimSpace(rockSplit[1]))}

			if start.col == -1 {
				start = pos
			} else if start.col == pos.col {
				n := start.row - pos.row

				if n > 0 {
					for i := 0; i <= n; i++ {
						if cave[pos.col] == nil {
							cave[pos.col] = map[int]bool{pos.row + i: true}
						} else {
							cave[pos.col][pos.row+i] = true
						}
					}
				} else {
					for i := 0; i >= n; i-- {
						if cave[pos.col] == nil {
							cave[pos.col] = map[int]bool{pos.row + i: true}
						} else {
							cave[pos.col][pos.row+i] = true
						}
					}
				}
			} else {
				n := start.col - pos.col
				if n > 0 {
					for i := 0; i <= n; i++ {
						if cave[pos.col+i] == nil {
							cave[pos.col+i] = map[int]bool{pos.row: true}
						} else {
							cave[pos.col+i][pos.row] = true
						}
					}
				} else {
					for i := 0; i >= n; i-- {
						if cave[pos.col+i] == nil {
							cave[pos.col+i] = map[int]bool{pos.row: true}
						} else {
							cave[pos.col+i][pos.row] = true
						}
					}
				}
			}

			if cave[pos.col] == nil {
				cave[pos.col] = map[int]bool{pos.row: true}
			} else {
				cave[pos.col][pos.row] = true
			}

			start = pos
		}
	}

	return cave
}

func getMaxKey(cave cave) int {
	max := -1
	for k := range cave {
		if k > max {
			max = k
		}
	}
	return max
}

func getUnitsOfSand(cave cave, drop position, floor int) int {
	units := 0
	highestPoint := getMaxKey(cave)

	for {
		sand := position{drop.row, drop.col}

		for sand.col < floor {
			// Blocked down
			if cave[sand.col+1][sand.row] {
				if !cave[sand.col+1][sand.row-1] {
					// Move Left+Down
					sand.row--
				} else if !cave[sand.col+1][sand.row+1] {
					// Move Right+Down
					sand.row++
				} else {
					break
				}
			}

			sand.col++
		}

		if sand.col == drop.col && sand.row == drop.row {
			return units + 1
		}

		if sand.col == floor && highestPoint < floor {
			sand.col--
		}

		if sand.col >= floor {
			return units
		}

		if cave[sand.col] == nil {
			cave[sand.col] = map[int]bool{sand.row: true}
		} else {
			cave[sand.col][sand.row] = true
		}

		units++
	}
}

func main() {
	input := os.Args[1]

	cave := loadRocks(input)
	units := getUnitsOfSand(cave, position{500, 0}, getMaxKey(cave))
	fmt.Println("1:", units)

	cave = loadRocks(input)
	units = getUnitsOfSand(cave, position{500, 0}, getMaxKey(cave)+2)
	fmt.Println("2:", units)
}
