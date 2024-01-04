package days

import (
	base "aoc2023"
	"fmt"
	"strconv"
	"testing"
)

func solution(input []string) string {
	memory := make(map[string]Cano)

	makeS("15", "54", input, memory)

	max := 0

	for _, value := range memory {
		if value.Distance > max {
			max = value.Distance
		}
	}

	return strconv.FormatInt(int64(max), 10)
}

/*


a partir do start, ir de cano em cano



*/

type Cano struct {
	Symbol     string
	X          string
	Y          string
	Distance   int
	conections []Cano
}

func makeOrGetCano(symbol string, x string, y string, distance int, memory map[string]Cano, origin Cano, grid []string) Cano {
	oldCano, ok := memory[x+"-"+y]
	if ok {
		if oldCano.Distance > distance {
			oldCano.Distance = distance
		}
		return oldCano
	}

	xInt, _ := strconv.ParseInt(x, 10, 64)
	yInt, _ := strconv.ParseInt(y, 10, 64)

	xOriginInt, _ := strconv.ParseInt(origin.X, 10, 64)
	yOriginInt, _ := strconv.ParseInt(origin.Y, 10, 64)

	var nextYInt int64
	var nextXInt int64

	cano := Cano{X: x, Y: y, Symbol: symbol, Distance: distance, conections: []Cano{origin}}

	switch symbol {
	case "|":
		{
			if yOriginInt == yInt+1 {
				nextYInt = yInt - 1
				nextXInt = xInt
				break
			}

			nextYInt = yInt + 1
			nextXInt = xInt
			break
		}
	case "-":
		{
			if xOriginInt == xInt+1 {
				nextYInt = yInt
				nextXInt = xInt - 1
				break
			}

			nextYInt = yInt
			nextXInt = xInt + 1
			break
		}

	/*
		L is a 90-degree bend connecting north and east.
		J is a 90-degree bend connecting north and west.
		7 is a 90-degree bend connecting south and west.
		F is a 90-degree bend connecting south and east.
	*/
	case "L":
		{
			if yOriginInt == yInt-1 {
				nextYInt = yInt
				nextXInt = xInt + 1
				break
			}

			nextYInt = yInt - 1
			nextXInt = xInt
			break
		}
	case "J":
		{
			{
				if yOriginInt == yInt-1 {
					nextYInt = yInt
					nextXInt = xInt - 1
					break
				}

				nextYInt = yInt - 1
				nextXInt = xInt
				break
			}
		}
	case "7":
		{
			{
				if yOriginInt == yInt-1 {
					nextYInt = yInt
					nextXInt = xInt - 1
					break
				}

				nextYInt = yInt - 1
				nextXInt = xInt
				break
			}
		}
	case "F":
		{
			{
				if yOriginInt == yInt-1 {
					nextYInt = yInt
					nextXInt = xInt + 1
					break
				}

				nextYInt = yInt - 1
				nextXInt = xInt
				break
			}
		}
	case "S":
		{
			north := string([]rune(grid[yInt+1])[xInt])
			if north == "|" {
				nextNorth := makeOrGetCano(north, strconv.FormatInt(xInt, 10), strconv.FormatInt(yInt+1, 10), distance+1, memory, cano, grid)
				cano.conections = append(cano.conections, nextNorth)
			}

			if xInt-1 > 0 {
				west := string([]rune(grid[yInt])[xInt-1])
				nextwest := makeOrGetCano(west, strconv.FormatInt(xInt-1, 10), strconv.FormatInt(yInt, 10), distance+1, memory, cano, grid)
				cano.conections = append(cano.conections, nextwest)
			}

			east := string([]rune(grid[yInt])[xInt+1])
			if east == "J" {
				nexteast := makeOrGetCano(east, strconv.FormatInt(xInt+1, 10), strconv.FormatInt(yInt, 10), distance+1, memory, cano, grid)
				cano.conections = append(cano.conections, nexteast)
			}

			south := string([]rune(grid[yInt-1])[xInt])
			if south == "|" {
				nextsouth := makeOrGetCano(south, strconv.FormatInt(xInt, 10), strconv.FormatInt(yInt-1, 10), distance+1, memory, cano, grid)
				cano.conections = append(cano.conections, nextsouth)
			}

			return cano
		}
	default:
		return cano
	}

	nextSymbol := string([]rune(grid[nextYInt])[nextXInt])
	next := makeOrGetCano(nextSymbol, strconv.FormatInt(nextXInt, 10), strconv.FormatInt(nextYInt, 10), distance+1, memory, cano, grid)
	memory[next.X+"-"+next.Y] = next
	cano.conections = append(cano.conections, next)
	memory[cano.X+"-"+cano.Y] = cano
	return cano
}

func makeS(x string, y string, grid []string, memory map[string]Cano) Cano {
	cano := Cano{Symbol: "S", X: x, Y: y, Distance: 0}

	xInt, _ := strconv.ParseInt(x, 10, 64)
	yInt, _ := strconv.ParseInt(y, 10, 64)

	north := string([]rune(grid[yInt+1])[xInt])
	if north == "|" {
		nextNorth := makeOrGetCano(north, strconv.FormatInt(xInt, 10), strconv.FormatInt(yInt+1, 10), 1, memory, cano, grid)
		cano.conections = append(cano.conections, nextNorth)
	}

	if xInt-1 > 0 {
		west := string([]rune(grid[yInt])[xInt-1])
		nextwest := makeOrGetCano(west, strconv.FormatInt(xInt-1, 10), strconv.FormatInt(yInt, 10), 1, memory, cano, grid)
		cano.conections = append(cano.conections, nextwest)
	}

	east := string([]rune(grid[yInt])[xInt+1])
	if east == "J" {
		nexteast := makeOrGetCano(east, strconv.FormatInt(xInt+1, 10), strconv.FormatInt(yInt, 10), 1, memory, cano, grid)
		cano.conections = append(cano.conections, nexteast)
	}

	south := string([]rune(grid[yInt-1])[xInt])
	if south == "|" {
		nextsouth := makeOrGetCano(south, strconv.FormatInt(xInt, 10), strconv.FormatInt(yInt-1, 10), 1, memory, cano, grid)
		cano.conections = append(cano.conections, nextsouth)
	}

	return cano

}

func TestMakeS(t *testing.T) {
	var input = []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}

	cano := makeS("0", "2", input, make(map[string]Cano))

	if cano.Distance != 0 {
		t.Errorf("input errado! want: 0, got %d", cano.Distance)
	}

	if cano.conections[0].Distance != 1 {
		t.Errorf("input errado! want: 1, got %d", cano.Distance)
	}

	fmt.Println(cano)
}

func TestSolution(t *testing.T) {
	var input = []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}

	var result = solution(input)

	if result != "8" {
		t.Errorf("input errado! want: 8, got %s", result)
	}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "jjeee" {
		t.Errorf("input errado! want: jjeee, got %s", result)
	}
}
