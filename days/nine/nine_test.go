package days

import (
	base "aoc2023"
	"testing"
)

func solution(input []string) string {
	return "2"
}

/*

generateNextNumbers

allZeroes

getNumAndAddToLast

last -> all equal -> get last and sum last from previous


*/

func allZeroes(ints []int) bool {
	for _, num := range ints {
		if num != 0 {
			return false
		}
	}
	return true
}

func generateNextNumbers(ints []int) []int {
	var nextNumbers []int
	for i, _ := range ints[1:] {
		nextNumbers = append(nextNumbers, ints[i+1]-ints[i])
	}

	return nextNumbers
}

func TestGenerateNextNumbers(t *testing.T) {
	ints := []int{0, 3, 6, 9, 12, 15}

	var result = generateNextNumbers(ints)

	if result[0] != 3 {
		t.Errorf("input errado! want: 3, got %d", result)
	}

	if result[3] != 3 {
		t.Errorf("input errado! want: 3, got %d", result)
	}
}

func TestSolution(t *testing.T) {
	var input = []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}

	var result = solution(input)

	if result != "114" {
		t.Errorf("input errado! want: 114, got %s", result)
	}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "jjeee" {
		t.Errorf("input errado! want: jjeee, got %s", result)
	}
}
