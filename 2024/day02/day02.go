package day02

import (
	"cbartolomeu/advent-of-code-2024/utils"
	"slices"
	"sort"
	"strings"
)

func isDescending(a []int) bool {
	slices.Reverse(a)
	return sort.IntsAreSorted(a)
}

func isAscending(a []int) bool {
	return sort.IntsAreSorted(a)
}

func isSafe(a []int) bool {
	if !isAscending(a) && !isDescending(a) {
		return false
	}

	for i := 1; i < len(a); i++ {
		lj := a[i-1]
		li := a[i]
		diff := utils.IntAbs(li - lj)

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func getPermutations(a []int) [][]int {
	var permutations [][]int

	permutations = append(permutations, slices.Clone(a))

	for i := 0; i < len(a); i++ {
		b := slices.Clone(a)
		perm := append(b[:i], b[i+1:]...)
		permutations = append(permutations, perm)
	}

	return permutations
}

func part1(filename string) int {
	input := utils.ReadInput(filename)

	safe_reports := 0
	for _, line := range input {
		levels := utils.ToIntSlice(strings.Split(line, " "))

		if isSafe(levels) {
			safe_reports++
		}
	}

	return safe_reports
}

func part2(filename string) int {
	input := utils.ReadInput(filename)

	safe_reports := 0
	for _, line := range input {
		levels := utils.ToIntSlice(strings.Split(line, " "))

		levels_perms := getPermutations(levels)

		isSafeLvl := slices.ContainsFunc(levels_perms, func(a []int) bool {
			return isSafe(a)
		})

		if isSafeLvl {
			safe_reports++
		}
	}

	return safe_reports
}
