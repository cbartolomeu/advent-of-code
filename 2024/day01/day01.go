package day01

import (
	"cbartolomeu/advent-of-code-2024/utils"
	"sort"
	"strings"
)

func GetLists(input []string) ([]int, []int) {
	var l_list []int
	var r_list []int

	for _, line := range input {
		list_entry := strings.Split(line, "   ")
		l_list = append(l_list, utils.ToInt(list_entry[0]))
		r_list = append(r_list, utils.ToInt(list_entry[1]))
	}

	return l_list, r_list
}

func part1(filename string) int {
	input := utils.ReadInput(filename)

	l_list, r_list := GetLists(input)

	sort.Ints(l_list)
	sort.Ints(r_list)

	distance := 0
	for i := range l_list {
		l_entry := l_list[i]
		r_entry := r_list[i]

		distance += utils.IntAbs(r_entry - l_entry)
	}

	return distance
}

func part2(filename string) int {
	input := utils.ReadInput(filename)

	l_list, r_list := GetLists(input)
	r_map := make(map[int]int)

	for i := range r_list {
		r_map[r_list[i]] += 1
	}

	score := 0
	for i := range l_list {
		location := l_list[i]
		if r_map[location] > 0 {
			score += r_map[location] * location
		}
	}

	return score
}
