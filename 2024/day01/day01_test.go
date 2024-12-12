package day01

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	want := 11
	got := part1("input-test.txt")

	if got != want {
		t.Errorf("[SHORT] Expected %v, but got %v", want, got)
	}

	want = 2970687
	got = part1("input.txt")

	if got != want {
		t.Errorf("[LARGE] Expected %v, but got %v", want, got)
	}
}

func Test_Part2(t *testing.T) {
	want := 31
	got := part2("input-test.txt")

	if got != want {
		t.Errorf("[SHORT] Expected %v, but got %v", want, got)
	}

	want = 23963899
	got = part2("input.txt")

	if got != want {
		t.Errorf("[LARGE] Expected %v, but got %v", want, got)
	}
}
