package days

import (
	base "aoc2023"
	"strconv"
	"strings"
	"testing"
)

func solution(input []string) string {
	inputClone := make([]string, len(input))

	copy(inputClone, input)

	gSum := 0

	var numNextCopies []int

	for _, line := range input {

		gSum += 1
		currNumCopies := 0

		if len(numNextCopies) > 0 {
			currNumCopies = numNextCopies[0]
			gSum += currNumCopies

			numNextCopies = numNextCopies[1:]
		}

		parts := strings.Split(strings.Split(line, ":")[1], "|")

		var winningNumbers []int
		for _, numStr := range strings.Split(strings.TrimSpace(parts[0]), " ") {
			if numStr == "" || numStr == " " {
				continue
			}

			i64, _ := strconv.ParseInt(numStr, 10, 32)
			winningNumbers = append(winningNumbers, int(i64))
		}

		ourNumbers := make(map[int]bool)
		for _, numStr := range strings.Split(strings.TrimSpace(parts[1]), " ") {
			if numStr == "" || numStr == " " {
				continue
			}
			i64, _ := strconv.ParseInt(numStr, 10, 32)
			ourNumbers[int(i64)] = true
		}

		matches := 0

		for _, num := range winningNumbers {
			if ourNumbers[num] {
				matches += 1
			}
		}

		counter := 0

		for matches != 0 {
			if len(numNextCopies)-counter <= 0 {
				numNextCopies = append(numNextCopies, 1+currNumCopies)
				matches -= 1
				counter += 1
				continue
			}

			numNextCopies[counter] += 1 + currNumCopies
			matches -= 1
			counter += 1
		}
	}
	return strconv.FormatInt(int64(gSum), 10)
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "10378710" {
		t.Errorf("input errado! want: 10378710, got %s", result)
	}
}
