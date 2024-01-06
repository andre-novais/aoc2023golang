package days

import (
	base "aoc2023"
	"testing"
)

func solution(input []string) int {
	return 3
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != 21 {
		t.Errorf("input errado! want: 12, got %d", result)
	}
}
