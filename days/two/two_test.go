package days

import (
	base "aoc2023"
	"strconv"
	"strings"
	"testing"
)

func solution(input []string) string {
	sum := 0

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	for _, line := range input {
		game := MakeGameMaxBalls(line)
		if !game.IsPossible(maxRed, maxBlue, maxGreen) {
			sum += game.id
		}

	}

	return strconv.FormatInt(int64(sum), 10)
}

type GamesMaxBalls struct {
	red   int
	blue  int
	green int
	id    int
}

func (Game *GamesMaxBalls) IsPossible(maxRed int, maxBlue int, maxGreen int) bool {
	return maxRed <= Game.red && maxBlue <= Game.blue && maxGreen <= Game.green
}

func MakeGameMaxBalls(line string) GamesMaxBalls {
	red := 0
	blue := 0
	green := 0

	data := strings.Split(line, ":")

	id, _ := strconv.ParseInt(strings.Split(strings.TrimSpace(data[0]), " ")[1], 10, 32)

	idInt := int(id)

	turns := strings.Split(data[1], ";")
	for _, turn := range turns {
		for _, ballSet := range strings.Split(turn, ",") {
			ballType := 0
			if strings.Contains(ballSet, "red") {
				ballType = 0
			} else if strings.Contains(ballSet, "blue") {
				ballType = 1
			} else {
				ballType = 2
			}

			num64, _ := strconv.ParseInt(strings.Split(strings.TrimSpace(ballSet), " ")[0], 10, 32)

			num := int(num64)

			switch ballType {
			case 0:
				if num > red {
					red = num
				}
			case 1:
				if num > blue {
					blue = num
				}
			case 2:
				if num > green {
					green = num
				}
			}
		}
	}

	return GamesMaxBalls{red: red, blue: blue, green: green, id: idInt}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "jjeee" {
		t.Errorf("input errado! want: jjeee, got %s", result)
	}
}
