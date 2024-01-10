package days

import (
	base "aoc2023"
	"fmt"
	"testing"
)

func solution(input []string) int {

	sum := 0
	var grid []string
	for _, line := range input {
		if line == "" {
			sum += getAboveRowsOrCollumns(separateLine(grid))
			grid = []string{}
			continue
		}

		grid = append(grid, line)
	}

	if len(grid) > 0 {
		sum += getAboveRowsOrCollumns(separateLine(grid))
	}

	return sum
}

func separateLine(lines []string) [][]rune {
	var result [][]rune

	for _, line := range lines {
		result = append(result, []rune(line))
	}

	return result
}

func getAboveRowsOrCollumns(grid [][]rune) int {

	colIndex := getCollumnReflexionIndex(grid)
	if colIndex != -1 {
		return colIndex + 1
	}

	rowIndex := getRowReflexionIndex(grid)
	if rowIndex != -1 {
		return (rowIndex + 1) * 100
	}

	fmt.Println(grid)
	return 0
}

func getCollumnReflexionIndex(grid [][]rune) int {
	for i := 0; i < len(grid[0])-1; i++ {
		if checkOtherCollumns(grid, i) {
			return i
		}
	}
	return -1
}

func checkOtherCollumns(grid [][]rune, leftCollumnIndex int) bool {
	k := leftCollumnIndex + 1
	foundASmudge := false
	for i := leftCollumnIndex; i >= 0; i-- {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] != grid[j][k] {
				if foundASmudge {
					return false
				}
				foundASmudge = true
			}
		}
		k += 1
		if k == len(grid[0]) {
			return foundASmudge
		}
	}
	return foundASmudge
}

func getRowReflexionIndex(grid [][]rune) int {
	for i := 0; i < len(grid)-1; i++ {
		if checkOtherRows(grid, i) {
			return i
		}
	}
	return -1
}

func checkOtherRows(grid [][]rune, upperRowIndex int) bool {
	foundASmudge := false
	k := upperRowIndex + 1
	for i := upperRowIndex; i >= 0; i-- {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != grid[k][j] {
				if foundASmudge {
					return false
				}
				foundASmudge = true
			}
		}
		k += 1
		if k == len(grid) {
			return foundASmudge
		}
	}
	return foundASmudge
}

func TestSolution(t *testing.T) {
	var input1 = []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}

	var result1 = getAboveRowsOrCollumns(separateLine(input1))

	if result1 != 300 {
		t.Errorf("input errado! want: 300, got %d", result1)
	}

	var input2 = []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}

	var result2 = getAboveRowsOrCollumns(separateLine(input2))

	if result2 != 100 {
		t.Errorf("input errado! want: 100, got %d", result2)
	}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != 23479 {
		t.Errorf("input errado! want: 23479, got %d", result)
	}
}
