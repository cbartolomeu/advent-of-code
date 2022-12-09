package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(s []int) int {
	total := 0

	for _, element := range s {
		total += element
	}

	return total
}

func findMin(s []int) (int, int) {
	minIdx, min := 0, s[0]
	for idx, element := range s {
		if element < min {
			minIdx, min = idx, element
		}
	}

	return minIdx, min
}

func getFoodCaloriesByElf(filename string, c chan int) {
	readFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	calories := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line != "" {
			foodCalories, err := strconv.Atoi(line)

			if err != nil {
				panic(err)
			}

			calories += foodCalories
		} else {
			c <- calories
			calories = 0
		}
	}

	readFile.Close()
	close(c)
}

func getTopNElfsWithMaxCalories(filename string, n int) int {
	c := make(chan int)

	go getFoodCaloriesByElf(filename, c)

	maxCalories := make([]int, n)

	for foodCalories := range c {
		minElfIdx, minElfCalories := findMin(maxCalories)
		if foodCalories > minElfCalories {
			maxCalories[minElfIdx] = foodCalories
		}
	}

	return sum(maxCalories)
}

func main() {
	input := os.Args[1]

	calories := getTopNElfsWithMaxCalories(input, 1)

	fmt.Printf("The Elf with most food calories has %d calories!\n", calories)

	calories = getTopNElfsWithMaxCalories(input, 3)

	fmt.Printf("The top 3 Elfs with most food calories have %d calories!\n", calories)
}
