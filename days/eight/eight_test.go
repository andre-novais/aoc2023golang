package days

import (
	base "aoc2023"
	"strconv"
	"strings"
	"testing"
)

func solution(input []string) string {
	commands := strings.Split(input[0], "")
	locMap := mapLineIntoDict(input[2:])

	positions := getAllStartingNodes(locMap)

	var stepsP []int
	for _, position := range positions {
		stepsP = append(stepsP, getNumSteps(position, locMap, commands))
	}

	if len(stepsP) == 1 {
		return strconv.FormatInt(int64(stepsP[0]), 10)
	}

	if len(stepsP) == 2 {
		return strconv.FormatInt(int64(stepsP[0]*stepsP[1]), 10)
	}
	return strconv.FormatInt(int64(LCM1Factor(stepsP)), 10)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func LCM1Factor(integersP []int) int {

	a := integersP[0]
	b := integersP[1]
	integers := integersP[2:]

	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		//fmt.Println(integers)
		result = LCM(result, integers[i])
	}

	return result
}

func getNumSteps(position string, locMap map[string][]string, commands []string) int {
	index := 0
	commandsLen := len(commands)

	steps := 0
	commandCurrStrIndex := 0

	for !allPositionsFinished([]string{position}) {
		if commands[index] == "R" {
			commandCurrStrIndex = 1
		} else {
			commandCurrStrIndex = 0
		}

		position = locMap[position][commandCurrStrIndex]

		index += 1

		if index == commandsLen {
			index = 0
		}

		steps += 1
	}
	return steps
}

func mapLineIntoDict(input []string) map[string][]string {
	locMap := make(map[string][]string)
	for _, line := range input {
		parts := strings.Split(line, "=")
		from := strings.TrimSpace(parts[0])
		toParts := strings.Split(
			strings.ReplaceAll(strings.ReplaceAll(
				strings.ReplaceAll(parts[1], " ", ""),
				"(", ""), ")", ""), ",")
		locMap[from] = toParts
	}

	return locMap
}

/*


map each line into a dict
iterate instruction to instruction
get next line from dict


*/

func getAllStartingNodes(locMap map[string][]string) []string {
	var positions []string
	for key, _ := range locMap {
		chars := strings.Split(key, "")

		if chars[len(chars)-1] == "A" {
			positions = append(positions, key)
		}
	}
	return positions
}

func allPositionsFinished(positions []string) bool {
	for _, position := range positions {
		chars := strings.Split(position, "")

		if chars[len(chars)-1] != "Z" {
			return false
		}
	}

	return true
}

func TestGetAllStartingNodes(t *testing.T) {
	var input = []string{
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}

	locMap := mapLineIntoDict(input)

	var result = getAllStartingNodes(locMap)

	if result[0] != "11A" {
		t.Errorf("input errado! want: 11A, got %s", result)
	}
	if result[1] != "22A" {
		t.Errorf("input errado! want: 22A, got %s", result)
	}
}

func TestAllPositionsFinished(t *testing.T) {
	positions1 := []string{"11Z", "LLL"}
	positions2 := []string{"11Z", "22Z"}

	var result1 = allPositionsFinished(positions1)
	var result2 = allPositionsFinished(positions2)

	if result1 != false {
		t.Errorf("input errado! want: false, got %t", result1)
	}
	if result2 != true {
		t.Errorf("input errado! want: true, got %t", result2)
	}
}

func TestMapLineIntoDict(t *testing.T) {
	var input = []string{
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	var result = mapLineIntoDict(input)

	if result["BBB"][1] != "EEE" {
		t.Errorf("input errado! want: EEE, got %s", result)
	}

	if result["GGG"][0] != "GGG" {
		t.Errorf("input errado! want: GGG, got %s", result)
	}
}

func TestSolution(t *testing.T) {
	var input = []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	var result = solution(input)

	if result != "2" {
		t.Errorf("input errado! want: 2, got %s", result)
	}
}

func TestSolution2(t *testing.T) {
	var input = []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	var result = solution(input)

	if result != "6" {
		t.Errorf("input errado! want: 6, got %s", result)
	}
}

func TestSolution3(t *testing.T) {
	var input = []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}

	var result = solution(input)

	if result != "6" {
		t.Errorf("input errado! want: 6, got %s", result)
	}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "12833235391111" {
		t.Errorf("input errado! want: 12833235391111, got %s", result)
	}
}
