package day03

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	want := 161
	got := part1("input-test.txt")

	if got != want {
		t.Errorf("[SHORT] Expected %v, but got %v", want, got)
	}

	want = 180233229
	got = part1("input.txt")

	if got != want {
		t.Errorf("[LARGE] Expected %v, but got %v", want, got)
	}
}

func Test_Part2(t *testing.T) {
	want := 48
	got := part2("input-test.txt")

	if got != want {
		t.Errorf("[SHORT] Expected %v, but got %v", want, got)
	}

	want = 95411583
	got = part2("input.txt")

	if got != want {
		t.Errorf("[LARGE] Expected %v, but got %v", want, got)
	}
}
