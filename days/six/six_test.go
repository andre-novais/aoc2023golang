package days

import (
	base "aoc2023"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func solution(input []string) string {

	var times []int
	var records []int

	for i, line := range input {
		if i == 0 {
			timeStr := strings.ReplaceAll(strings.TrimSpace(strings.Split(line, ":")[1]), " ", "")
			time64, _ := strconv.ParseInt(timeStr, 10, 64)

			if time64 != 0 {
				times = append(times, int(time64))
			}
			continue
		}
		if i == 1 {
			recordStr := strings.ReplaceAll(strings.TrimSpace(strings.Split(line, ":")[1]), " ", "")
			record64, _ := strconv.ParseInt(recordStr, 10, 64)

			if record64 != 0 {
				records = append(records, int(record64))
			}
			continue
		}
	}

	result := 1
	for i, time := range times {
		fmt.Println(times, i, len(times))
		race := Race{time: time, record: records[i]}
		result *= race.calcNumWays()
	}

	return strconv.FormatInt(int64(result), 10)
}

type Race struct {
	time   int
	record int
}

func (race *Race) calcNumWays() int {
	numWays := 0

	for i := 1; i < race.time; i++ {
		if i*(race.time-i) > race.record {
			numWays += 1
		}
	}

	return numWays
}

/*
times []int
recordDistance []int

race {
	time int
	recordDistance int
}

for (,,,) {
	if could win {
		push
	}
}


*/
func TestCalcNumWays(t *testing.T) {
	var race = Race{7, 9}

	var result = race.calcNumWays()

	if result != 4 {
		t.Errorf("input errado! want: 4, got %d", result)
	}
}
func TestSolution(t *testing.T) {
	var input = []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	var result = solution(input)

	if result != "71503" {
		t.Errorf("input errado! want: 71503, got %s", result)
	}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "38220708" {
		t.Errorf("input errado! want: 38220708, got %s", result)
	}
}
