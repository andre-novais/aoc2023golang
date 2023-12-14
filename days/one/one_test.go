package days

import (
	base "aoc2023"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func solution(input []string) string {
	sum := 0

	for _, line := range input {
		line = strings.ReplaceAll(line, "oneight", "oneeight")
		line = strings.ReplaceAll(line, "twone", "twoone")
		line = strings.ReplaceAll(line, "sevenine", "sevennine")
		line = strings.ReplaceAll(line, "eighthree", "eightthree")
		line = strings.ReplaceAll(line, "eightwo", "eighttwo")

		re := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")

		var nums = re.FindAllString(line, -1)
		var numStr = convToNum(nums[0]) + convToNum(nums[len(nums)-1])

		var holeNum, err = strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			panic(err)
		}

		sum += int(holeNum)
	}
	return strconv.FormatInt(int64(sum), 10)
}

func convToNum(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return s
	}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result2 = solution(input)

	if result2 != "jjeee" {
		t.Errorf("input errado! want: jjeee, got %s", result2)
	}
}
