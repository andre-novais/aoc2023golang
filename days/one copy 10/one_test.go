package days

import (
	base "aoc2023"
	"testing"
)

func solution(input []string) string {

}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "jjeee" {
		t.Errorf("input errado! want: jjeee, got %d", result)
	}
}
