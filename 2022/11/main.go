package main

import (
	"fmt"
	"strconv"
)

type Monkey struct {
	items     []int
	operation func(old int) int
	div       int
	divT      int
	divF      int
	n         int
}

func toInt(raw string) int {
	res, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return res
}

func toIntSlice(raw []string) []int {
	slice := make([]int, len(raw))

	for i, item := range raw {
		slice[i] = toInt(item)
	}

	return slice
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

func dot(arr []int) int {
	res := 1
	for _, it := range arr {
		res *= it
	}
	return res
}

func loadMonkeysTest() []Monkey {
	return []Monkey{
		{items: []int{79, 98}, operation: func(old int) int { return old * 19 }, div: 23, divT: 2, divF: 3},
		{items: []int{54, 65, 75, 74}, operation: func(old int) int { return old + 6 }, div: 19, divT: 2, divF: 0},
		{items: []int{79, 60, 97}, operation: func(old int) int { return old * old }, div: 13, divT: 1, divF: 3},
		{items: []int{74}, operation: func(old int) int { return old + 3 }, div: 17, divT: 0, divF: 1},
	}
}

func loadMonkeys() []Monkey {
	return []Monkey{
		{items: []int{83, 88, 96, 79, 86, 88, 70}, operation: func(old int) int { return old * 5 }, div: 11, divT: 2, divF: 3},
		{items: []int{59, 63, 98, 85, 68, 72}, operation: func(old int) int { return old * 11 }, div: 5, divT: 4, divF: 0},
		{items: []int{90, 79, 97, 52, 90, 94, 71, 70}, operation: func(old int) int { return old + 2 }, div: 19, divT: 5, divF: 6},
		{items: []int{97, 55, 62}, operation: func(old int) int { return old + 5 }, div: 13, divT: 2, divF: 6},
		{items: []int{74, 54, 94, 76}, operation: func(old int) int { return old * old }, div: 7, divT: 0, divF: 3},
		{items: []int{58}, operation: func(old int) int { return old + 4 }, div: 17, divT: 7, divF: 1},
		{items: []int{66, 63}, operation: func(old int) int { return old + 6 }, div: 2, divT: 7, divF: 5},
		{items: []int{56, 56, 90, 96, 68}, operation: func(old int) int { return old + 7 }, div: 3, divT: 4, divF: 1},
	}
}

func getNMoreActiveMonkeys(monkeys []Monkey, n int) int {
	active := make([]int, n)

	for _, monkey := range monkeys {
		j, min := findMin(active)
		if monkey.n > min {
			active[j] = monkey.n
		}
	}

	return dot(active)
}

func inspectDuringNRounds(monkeys []Monkey, worryLevelUpdate func(old int) int, rounds int) ([]Monkey, int) {
	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			for _, item := range monkeys[j].items {
				worryLevel := monkeys[j].operation(item)
				worryLevel = worryLevelUpdate(worryLevel)

				nextMonkey := monkeys[j].divF
				if worryLevel%monkeys[j].div == 0 {
					nextMonkey = monkeys[j].divT
				}
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, worryLevel)
			}
			monkeys[j].n += len(monkeys[j].items)
			monkeys[j].items = []int{}
		}
	}

	return monkeys, getNMoreActiveMonkeys(monkeys, 2)
}

func getCommonDivisor(monkeys []Monkey) int {
	common := 1
	for _, monkey := range monkeys {
		common *= monkey.div
	}
	return common
}

func main() {
	monkeys, monkeyBusiness := loadMonkeysTest(), 0
	monkeys, monkeyBusiness = inspectDuringNRounds(monkeys, func(old int) int { return old / 3 }, 20)
	fmt.Println("(test) 1:", monkeyBusiness)

	monkeys, monkeyBusiness = loadMonkeys(), 0
	monkeys, monkeyBusiness = inspectDuringNRounds(monkeys, func(old int) int { return old / 3 }, 20)
	fmt.Println("1:", monkeyBusiness)

	monkeys, monkeyBusiness = loadMonkeysTest(), 0
	commonDivisor := getCommonDivisor(monkeys)
	monkeys, monkeyBusiness = inspectDuringNRounds(monkeys, func(old int) int { return old % commonDivisor }, 10000)
	fmt.Println("(test) 2:", monkeyBusiness)

	monkeys, monkeyBusiness = loadMonkeys(), 0
	commonDivisor = getCommonDivisor(monkeys)
	monkeys, monkeyBusiness = inspectDuringNRounds(monkeys, func(old int) int { return old % commonDivisor }, 10000)
	fmt.Println("2:", monkeyBusiness)
}
