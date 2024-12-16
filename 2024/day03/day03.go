package day03

import (
	"cbartolomeu/advent-of-code-2024/utils"
	"regexp"
)

func calculateMul(str string) int {
	single_mul_regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := single_mul_regex.FindStringSubmatch(str)
	a := utils.ToInt(matches[1])
	b := utils.ToInt(matches[2])

	return a * b
}

func part1(filename string) int {
	input := utils.ReadInput(filename)
	mul_regex := regexp.MustCompile(`(mul\(\d+,\d+\)){1}`)

	result := 0
	for _, line := range input {
		matches := mul_regex.FindAllString(line, -1)
		for _, match := range matches {
			result += calculateMul(match)
		}
	}

	return result
}

func part2(filename string) int {
	input := utils.ReadInput(filename)
	mul_regex := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\)){1}`)

	result := 0
	enable := true
	for _, line := range input {
		matches := mul_regex.FindAllString(line, -1)
		for _, match := range matches {
			if match == "don't()" {
				enable = false
			} else if match == "do()" {
				enable = true
			} else if enable {
				result += calculateMul(match)
			}
		}
	}

	return result
}
