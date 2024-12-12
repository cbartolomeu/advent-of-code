package day02

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	want := 2
	got := part1("input-test.txt")

	if got != want {
		t.Errorf("[SHORT] Expected %v, but got %v", want, got)
	}

	want = 279
	got = part1("input.txt")

	if got != want {
		t.Errorf("[LARGE] Expected %v, but got %v", want, got)
	}
}

func Test_Part2(t *testing.T) {
	want := 4
	got := part2("input-test.txt")

	if got != want {
		t.Errorf("[SHORT] Expected %v, but got %v", want, got)
	}

	want = 343
	got = part2("input.txt")

	if got != want {
		t.Errorf("[LARGE] Expected %v, but got %v", want, got)
	}
}
