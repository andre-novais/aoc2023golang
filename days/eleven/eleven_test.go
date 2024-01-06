package days

import (
	base "aoc2023"
	"math"
	"strings"
	"testing"
)

/*


1- expand the universe:
	1- find the row gaps
	2- expand rows
	3 - find the collumn gaps
	4 - expand collumns

2 - get shortest distance of each


3 - sum distances


*/

func expand(rawInput []string) ([][]string, []int, []int) {
	transformed := transformToArrs(rawInput)
	rows, cols := findRowAndColGaps(transformed)

	// for count, rowIndex := range rows {
	// 	transformed = includeRowGap(transformed, rowIndex+count)
	// }

	// for count, colIndex := range cols {
	// 	transformed = includeCollumnGap(transformed, colIndex+count)
	// }

	return transformed, rows, cols
}

// func TestExpansion(t *testing.T) {
// 	var input = []string{
// 		"...#......",
// 		".......#..",
// 		"#.........",
// 		"..........",
// 		"......#...",
// 		".#........",
// 		".........#",
// 		"..........",
// 		".......#..",
// 		"#...#.....",
// 	}

// 	var result = expand(input)

// 	var expected = []string{
// 		"....#........",
// 		".........#...",
// 		"#............",
// 		".............",
// 		".............",
// 		"........#....",
// 		".#...........",
// 		"............#",
// 		".............",
// 		".............",
// 		".........#...",
// 		"#....#.......",
// 	}

// 	for i, line := range expected {
// 		if line != strings.Join(result[i][:], "") {
// 			t.Errorf("input errado! want: line, got %s", strings.Join(result[i][:], ""))
// 		}
// 	}
// }

func transformToArrs(input []string) [][]string {
	var result [][]string
	for _, line := range input {
		result = append(result, strings.Split(line, ""))
	}

	return result
}

func findRowAndColGaps(input [][]string) (rows []int, cols []int) {
	for i := 0; i < len(input); i++ {
		found := true
		for _, spot := range input[i] {
			if spot == "#" {
				found = false
			}
		}

		if found {
			rows = append(rows, i)
		}
	}

	for i := 0; i < len(input[0]); i++ {
		found := true
		for j := 0; j < len(input); j++ {
			if input[j][i] == "#" {
				found = false
			}
		}

		if found {
			cols = append(cols, i)
		}
	}

	return rows, cols

}

func TestFindRowAndColGaps(t *testing.T) {
	var input = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	transformed := transformToArrs(input)

	rows, cols := findRowAndColGaps(transformed)

	if rows[0] != 3 {
		t.Errorf("input errado! want: 3, got %d", rows[0])
	}
	if rows[1] != 7 {
		t.Errorf("input errado! want: 7, got %d", rows[1])
	}

	if cols[0] != 2 {
		t.Errorf("input errado! want: 2, got %d", cols[0])
	}

	if cols[1] != 5 {
		t.Errorf("input errado! want: 5, got %d", cols[1])
	}

	if cols[2] != 8 {
		t.Errorf("input errado! want: 8, got %d", cols[2])
	}
}

func includeRowGap(input [][]string, index int) [][]string {
	rowLength := len(input[0])

	var newRow []string
	for i := 0; i < rowLength; i++ {
		newRow = append(newRow, ".")
	}

	input = append(input[:index+1], input[index:]...)
	input[index] = newRow

	return input
}

func TestIncludeRowGap(t *testing.T) {
	var input = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	transformed := transformToArrs(input)

	var result = includeRowGap(transformed, 3)

	var expected = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	for i, line := range expected {
		if line != strings.Join(result[i][:], "") {
			t.Errorf("input errado! want: line, got %s", strings.Join(result[i][:], ""))
		}
	}
}

func includeCollumnGap(input [][]string, index int) [][]string {
	colLength := len(input)

	for i := 0; i < colLength; i++ {
		input[i] = append(input[i][:index+1], input[i][index:]...)
		input[i][index] = "."
	}

	return input
}

func TestIncludeCollumnGap(t *testing.T) {
	var input = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	transformed := transformToArrs(input)

	var result = includeCollumnGap(transformed, 5)

	var expected = []string{
		"...#.......",
		"........#..",
		"#..........",
		"...........",
		".......#...",
		".#.........",
		"..........#",
		"...........",
		"........#..",
		"#...#......",
	}

	for i, line := range expected {
		if line != strings.Join(result[i][:], "") {
			t.Errorf("input errado! want: %s, got %s", line, strings.Join(result[i][:], ""))
		}
	}
}

func findGalaxies(mapp [][]string) (galaxies [][]int) {

	for i := 0; i < len(mapp[0]); i++ {
		for j := 0; j < len(mapp); j++ {
			if mapp[j][i] == "#" {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	return galaxies
}

func TestFindGalaxies(t *testing.T) {
	var input = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	var transformed = transformToArrs(input)
	galaxies := findGalaxies(transformed)

	if galaxies[0][0] != 0 {
		t.Errorf("input errado! want: 0, got %d", galaxies[0][0])
	}
	if galaxies[0][1] != 2 {
		t.Errorf("input errado! want: 2, got %d", galaxies[0][1])
	}
}

func findShrotestDistance(x1 int, y1 int, x2 int, y2 int, expandedRows []int, expandedCols []int, expandedNum int) int {

	numIntersected := 0
	for _, row := range expandedCols {
		if (x1 > row && x2 < row) || (x1 < row && x2 > row) {
			numIntersected += 1
		}
	}

	for _, col := range expandedRows {
		if (y1 > col && y2 < col) || (y1 < col && y2 > col) {
			numIntersected += 1
		}
	}

	//fmt.Println(x1, y1, x2, y2, (int(math.Abs(float64(x2-x1))+math.Abs(float64(y2-y1))) + numIntersected*expandedNum - numIntersected), numIntersected, numIntersected*expandedNum-numIntersected)

	//fmt.Printf("{ x: %d, y: %d }, { x: %d, y: %d }, raw: %d, extra: %d, total: %d, intersects: %d\n\n", x1, y1, x2, y2, int(math.Abs(float64(x2-x1))+math.Abs(float64(y2-y1))), numIntersected*expandedNum-numIntersected, (int(math.Abs(float64(x2-x1))+math.Abs(float64(y2-y1))) + numIntersected*expandedNum - numIntersected), numIntersected)
	num := (int(math.Abs(float64(x2-x1))+math.Abs(float64(y2-y1))) + numIntersected*expandedNum - numIntersected)

	return num
}

func TestFindShrotestDistance(t *testing.T) {
	var result1 = findShrotestDistance(0, 2, 12, 7, []int{}, []int{}, 1)

	if result1 != 17 {
		t.Errorf("input errado! want: 17, got %d", result1)
	}
	var result2 = findShrotestDistance(0, 2, 12, 7, []int{}, []int{}, 1)

	if result2 != 17 {
		t.Errorf("input errado! want: 17, got %d", result1)
	}
}

func findDistances(posis [][]int, rows []int, cols []int, expandedNum int) (dist []int) {

	for i := 0; i < len(posis); i++ {
		for j := i; j < len(posis); j++ {
			if i == j {
				continue
			}

			dist = append(dist, findShrotestDistance(posis[i][0], posis[i][1], posis[j][0], posis[j][1], rows, cols, expandedNum))

		}
	}

	return dist
}

func sumDistances(dist []int) (sum int) {
	for _, i := range dist {
		sum += i
	}

	return sum
}

func solution(input []string) int {
	actualMap, rows, cols := expand(input)
	galaxies := findGalaxies(actualMap)
	distances := findDistances(galaxies, rows, cols, 2)
	return sumDistances(distances)
}

func TestSolution(t *testing.T) {
	var input = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	var result = solution(input)

	if result != 374 {
		t.Errorf("input errado! want: 374, got %d", result)
	}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != 9605127 {
		t.Errorf("input errado! want: 9605127, got %d", result)
	}
}

func solution2(input []string) int {
	actualMap, rows, cols := expand(input)
	galaxies := findGalaxies(actualMap)
	distances := findDistances(galaxies, rows, cols, 1000000)
	return sumDistances(distances)
}

func TestTest2(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution2(input)

	if result != 458191688761 {
		t.Errorf("input errado! want: 458191688761, got %d", result)
	}
}
