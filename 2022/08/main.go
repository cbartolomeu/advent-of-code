package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Grid []([]int)

func toInt(raw string) int {
	section, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return section
}

func getGrid(filename string) Grid {
	grid := Grid{}

	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		row := []int{}
		for _, v := range line {
			row = append(row, toInt(string(v)))
		}
		grid = append(grid, row)
	}

	return grid
}

func getCol(grid Grid, col int) []int {
	arr := []int{}
	for i := 0; i < len(grid); i++ {
		arr = append(arr, grid[i][col])
	}
	return arr
}

func max(arr []int) int {
	res := -1
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}

func getVisibleTrees(grid Grid) int {
	visible := len(grid[0])*2 + (len(grid)-2)*2

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			tree := grid[i][j]
			// left + right + up + down
			col := getCol(grid, j)
			if tree > max(grid[i][:j]) || tree > max(grid[i][j+1:]) || tree > max(col[:i]) || tree > max(col[i+1:]) {
				visible++
			}
		}
	}

	return visible
}

func getTreeScenicScore(grid Grid, row int, col int) int {
	currentTree := grid[row][col]

	scoreRight := 1
	for k := col + 1; k < len(grid[row])-1; k++ {
		if grid[row][k] >= currentTree {
			break
		}
		scoreRight++
	}

	scoreLeft := 1
	for k := col - 1; k > 0; k-- {
		if grid[row][k] >= currentTree {
			break
		}
		scoreLeft++
	}

	scoreBottom := 1
	for k := row + 1; k < len(grid[row])-1; k++ {
		if grid[k][col] >= currentTree {
			break
		}
		scoreBottom++
	}

	scoreTop := 1
	for k := row - 1; k > 0; k-- {
		if grid[k][col] >= currentTree {
			break
		}
		scoreTop++
	}

	return scoreBottom * scoreLeft * scoreRight * scoreTop
}

func getHighestScenicScore(grid Grid) int {
	maxScore := 0

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			score := getTreeScenicScore(grid, i, j)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func main() {
	input := os.Args[1]

	grid := getGrid(input)
	fmt.Println(getVisibleTrees(grid))
	fmt.Println(getHighestScenicScore(grid))
}
