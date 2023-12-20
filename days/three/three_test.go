package days

import (
	base "aoc2023"
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func solution(input []string) string {
	sum := 0
	var parNumbers []PartNumber

	numbersRe := regexp.MustCompile("[0-9]+")
	for iRow, row := range input {
		numsIndexes := numbersRe.FindAllStringIndex(row, -1)

		for _, indexes := range numsIndexes {
			num, _ := strconv.ParseInt(row[indexes[0]:indexes[1]], 10, 32)
			partNumber := PartNumber{row: int64(iRow), start: int64(indexes[0]), end: int64(indexes[1]) - 1, num: num}

			parNumbers = append(parNumbers, partNumber)
		}
	}

	// fmt.Println(FindStarIndexes(input))

	for _, star := range FindStarIndexes(input) {

		starPartNums := GetAdgencents(input, parNumbers, star)
		if len(starPartNums) == 2 {
			sum += int(starPartNums[0].num * starPartNums[1].num)
		}
		if len(starPartNums) != 2 {

			fmt.Printf("star %d has %d partNums: %d\n", star, len(starPartNums), starPartNums)
		}
	}

	return strconv.FormatInt(int64(sum), 10)
}

type PartNumber struct {
	row   int64
	start int64
	end   int64
	num   int64
}

//[row, col]
func FindStarIndexes(grid []string) [][]int {
	var indexes [][]int
	starRe := regexp.MustCompile("\\*")

	for iRow, row := range grid {
		stars := starRe.FindAllStringIndex(row, -1)

		for _, star := range stars {
			indexes = append(indexes, []int{iRow, star[0]})
		}
	}

	return indexes
}

func GetAdgencents(grid []string, partNumbers []PartNumber, star []int) []PartNumber {
	var adgencents []PartNumber

	for _, partNumber := range partNumbers {
		one := int(partNumber.row) >= star[0]-1
		two := int(partNumber.row) <= star[0]+1
		three := int(partNumber.start)-1 >= star[1]
		four := int(partNumber.end)+1 <= star[1]
		if star[0] == 8 && star[1] == 23 && partNumber.num == 582 {
			fmt.Println(one, two, three, four)
		}

		if int(partNumber.row) >= star[0]-1 && int(partNumber.row) <= star[0]+1 && int(partNumber.start)-1 <= star[1] && int(partNumber.end)+1 >= star[1] {

			adgencents = append(adgencents, partNumber)
		}
	}

	return adgencents
}

func IsAdgencent(grid []string, partNumber PartNumber) bool {
	specialRe := regexp.MustCompile("[^0-9 | ^.]")
	start := partNumber.start - 1
	if partNumber.start == 0 {
		start = 0
	}

	end := partNumber.end + 1
	if partNumber.end >= int64(len(grid[0])) {
		end = partNumber.end
	}

	if partNumber.row > 0 {
		fmt.Println("upper := grid[partNumber.row-1]")
		upper := grid[partNumber.row-1]
		uInter := upper[start:end]

		isAdgencent := specialRe.Match([]byte(uInter))
		if isAdgencent {
			return true
		}
	}

	if partNumber.row < int64(len(grid)-1) {
		fmt.Println("lower := grid[partNumber.row+1]")
		lower := grid[partNumber.row+1]
		uLower := lower[start:end]

		isAdgencent := specialRe.Match([]byte(uLower))
		if isAdgencent {
			return true
		}
	}

	fmt.Println("middle := grid[partNumber.row]")
	fmt.Println(partNumber)
	middle := grid[partNumber.row]
	uMiddle := middle[start:end]

	isAdgencent := specialRe.Match([]byte(uMiddle))
	return isAdgencent
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "jjeee" {
		t.Errorf("input errado! want: jjeee, got %s", result)
	}
}
