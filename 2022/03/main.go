package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getRucksacks(filename string, c chan []string, n int) {
	readFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	rucksacks := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if n == 1 {
			c1, c2 := line[:len(line)/2], line[len(line)/2:]

			c <- []string{c1, c2}
		} else {
			rucksacks = append(rucksacks, line)

			if len(rucksacks) == n {
				c <- rucksacks
				rucksacks = []string{}
			}
		}
	}

	readFile.Close()
	close(c)
}

func isPriorityItem(item rune, rucksack []string, idx int) bool {
	if idx >= len(rucksack) {
		return true
	}

	if strings.ContainsRune(rucksack[idx], item) {
		return isPriorityItem(item, rucksack, idx+1)
	}

	return false
}

func getPriority(rucksack []string) int {
	for _, item := range rucksack[0] {
		if isPriorityItem(item, rucksack, 1) {
			if item >= 'a' {
				return int(item-'a') + 1
			} else {
				return int(item-'A') + 1 + 26
			}
		}
	}

	return 0
}

func getItemsPriority(filename string, n int) int {
	c := make(chan []string)

	go getRucksacks(filename, c, n)

	priority := 0

	for rucksack := range c {
		priority += getPriority(rucksack)
	}

	return priority
}

func main() {
	input := os.Args[1]

	priority := getItemsPriority(input, 1)
	fmt.Println(priority)

	priority = getItemsPriority(input, 3)
	fmt.Println(priority)
}
