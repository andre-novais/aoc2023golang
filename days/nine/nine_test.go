package days

import (
	base "aoc2023"
	"strconv"
	"strings"
	"testing"
)

func solution(input []string) string {

	sum := 0

	for _, line := range input {
		sum += solutionLine(line)
	}

	return strconv.FormatInt(int64(sum), 10)
}

func solutionLine(line string) int {
	numsStrs := strings.Split(line, " ")

	var startNums []int
	for _, numStr := range numsStrs {
		num, _ := strconv.ParseInt(numStr, 10, 64)
		startNums = append(startNums, int(num))
	}

	stack := [][]int{startNums}

	for !allZeroes(stack[len(stack)-1]) {
		stack = append(stack, generateNextNumbers(stack[len(stack)-1]))
	}

	acc1 := 0

	for i := range stack {
		k := len(stack) - 1 - i

		if i == 0 {
			continue
		}

		acc1 = stack[k][0] - acc1
	}

	return acc1
}

/*

generateNextNumbers

allZeroes

getNumAndAddToLast

last -> all equal -> get last and sum last from previous




stack

pega linha e gera proximo numero, inclui no stack
continua até ultimo item sõ tiver zeros
vai pegando ultimo valor e somando com o previo

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

func TestSolutionLine(t *testing.T) {
	inputLine := "0 3 6 9 12 15"

	var result = solutionLine(inputLine)

	if result != -3 {
		t.Errorf("input errado! want: -3, got %d", result)
	}
}

func TestSolution(t *testing.T) {
	var input = []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}

	var result = solution(input)

	if result != "2" {
		t.Errorf("input errado! want: 2, got %s", result)
	}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "1016" {
		t.Errorf("input errado! want: 1016, got %s", result)
	}
}
