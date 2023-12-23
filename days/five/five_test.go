package days

import (
	base "aoc2023"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

//segunda iteração o numero é maior que deveria pode ser
func solution(input []string) string {
	seeds, sourceDestMaps := getSeedAndMapSections(input)

	var seedLocations []int
	//seed := seeds[0]

	// fmt.Println(seed)
	for _, seed := range seeds {

		// sdm := findSourceDestMap("seed", sourceDestMaps)

		// temp := sdm.translate(seed)

		// for sdm.dest != "location" {
		// 	// fmt.Println(temp)
		// 	sdm = findSourceDestMap(sdm.dest, sourceDestMaps)
		// 	temp = sdm.translate(temp)
		// }

		seedLocations = append(seedLocations, divideByElevenRecur(seed[0], seed[1], sourceDestMaps))
	}

	min := math.MaxInt
	for _, location := range seedLocations {
		if location < min {
			min = location
		}
	}

	return strconv.FormatInt(int64(min), 10)
}

func getLocationFromSeed(seed int, maps []SourceDestMap) int {
	sdm := findSourceDestMap("seed", maps)

	temp := sdm.translate(seed)

	for sdm.dest != "location" {
		sdm = findSourceDestMap(sdm.dest, maps)
		temp = sdm.translate(temp)
	}

	return temp
}

func divideByElevenRecur(start int, rangeNum int, maps []SourceDestMap) int {
	maxRange := 1000
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(start, rangeNum)

	if rangeNum < maxRange {
		var ranges []int
		for i := 0; i < rangeNum; i++ {
			ranges = append(ranges, getLocationFromSeed(start+i, maps))
		}

		min := math.MaxInt
		for _, i := range ranges {
			if i < min {
				min = i
			}
		}

		return min
	}

	var ranges []int
	for i := 0; i <= maxRange; i++ {
		ranges = append(ranges, getLocationFromSeed(start+((rangeNum*i)/maxRange), maps))
	}

	min := math.MaxInt
	indexMin := 0
	for index, i := range ranges {
		if i < min {
			min = i
			indexMin = index
		}
	}

	fmt.Println(ranges, indexMin)

	if indexMin == 0 {
		return ranges[0]
	}

	return divideByElevenRecur(start+((rangeNum*(indexMin-1))/maxRange), rangeNum/maxRange, maps)
}

func findSourceDestMap(source string, maps []SourceDestMap) SourceDestMap {
	var found SourceDestMap
	for _, sdm := range maps {
		if sdm.source == source {
			return sdm
		}
	}

	return found
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "20191102" {
		t.Errorf("input errado! want: 20191102, got %s", result)
	}
}

type SourceDestMap struct {
	source string
	dest   string
	mapp   []numberMap
}

type numberMap struct {
	startingSource int
	startingDest   int
	valueRange     int
}

func (sdm *SourceDestMap) translate(sourceInt int) int {
	for _, mapp := range sdm.mapp {
		if sourceInt >= mapp.startingSource && sourceInt <= mapp.startingSource+mapp.valueRange-1 {
			return mapp.startingDest + (sourceInt - mapp.startingSource)
		}
	}

	return sourceInt
}

func makeSourceDestMap(source string, dest string) SourceDestMap {
	return SourceDestMap{source: source, dest: dest, mapp: make([]numberMap, 3)}
}

func getSeedAndMapSections(input []string) ([][2]int, []SourceDestMap) {
	var seeds [][2]int
	var sourceDestMaps []SourceDestMap

	mapSecStrAgg := ""
	for i, line := range input {
		if i == 0 {
			fmt.Println(line)
			seedsStrs := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")

			seed := 0
			for _, seedStr := range seedsStrs {
				if seed == 0 {
					seed64, _ := strconv.ParseInt(seedStr, 10, 64)
					seed = int(seed64)
					continue
				}

				seedRange64, _ := strconv.ParseInt(seedStr, 10, 64)

				seeds = append(seeds, [2]int{seed, int(seedRange64)})
				seed = 0
			}
			continue
		}

		if line == "" && mapSecStrAgg != "" {
			sourceDestMaps = append(sourceDestMaps, createSourceDestMapFromSection(mapSecStrAgg))
			mapSecStrAgg = ""
			continue
		}

		mapSecStrAgg += line + "\n"
	}

	if mapSecStrAgg != "" {
		sourceDestMaps = append(sourceDestMaps, createSourceDestMapFromSection(mapSecStrAgg))
	}

	return seeds, sourceDestMaps
}

func createSourceDestMapFromSection(mapSection string) SourceDestMap {
	parts := strings.Split(mapSection, ":")

	sourceDest := strings.Split(parts[0], " ")[0]
	splitedSourceDest := strings.Split(strings.ReplaceAll(sourceDest, "\n", ""), "-")

	source := splitedSourceDest[0]
	dest := splitedSourceDest[2]

	sourceDestMap := makeSourceDestMap(source, dest)

	for _, line := range strings.Split(parts[1], "\n") {
		if line == "" {
			continue
		}

		splitedLine := strings.Split(line, " ")

		sourceStartingRange, _ := strconv.ParseInt(splitedLine[1], 10, 64)
		destStartingRange, _ := strconv.ParseInt(splitedLine[0], 10, 64)
		valueRange, _ := strconv.ParseInt(splitedLine[2], 10, 64)

		sourceDestMap.mapp = append(
			sourceDestMap.mapp,
			numberMap{
				startingSource: int(sourceStartingRange),
				startingDest:   int(destStartingRange),
				valueRange:     int(valueRange),
			},
		) //  [int(sourceStartingRange)] = int(destStartingRange)
	}

	return sourceDestMap
}

//test for first parser, as the elfs were wrong this test stoped being usefull
// func TestGetSeedAndMapSections(t *testing.T) {
// 	var input []string
// 	input = append(input, "seeds: 1 2 3 4")
// 	input = append(input, "")
// 	input = append(input, "seed-to-soil map:")
// 	input = append(input, "1 5 2")

// 	seeds, mapSections := getSeedAndMapSections(input)

// 	if !reflect.DeepEqual(seeds, []int{1, 2, 3, 4, 5, 6}) {
// 		t.Errorf("seeds erradas! want: 1, 2, 3, 4, 5, 6, got %d", seeds)
// 	}

// 	expected := makeSourceDestMap("seed", "soil")
// 	expected.mapp = append(expected.mapp, numberMap{
// 		startingSource: int(5),
// 		startingDest:   int(1),
// 		valueRange:     int(2),
// 	})

// 	if !reflect.DeepEqual(mapSections[0], expected) {
// 		t.Errorf("mapSections erradas! want: %+v, got  %+v", expected.source, mapSections[0].source)
// 	}

// }

func TestCreateSourceDestMapFromSection(t *testing.T) {
	//var input = base.GetInput("./input.txt")
	mapSectionStr := "seed-to-soil map:\n1 5 2\n10 20 1"

	sourceDestmap := createSourceDestMapFromSection(mapSectionStr)

	expected := makeSourceDestMap("seed", "soil")
	expected.mapp = append(expected.mapp, numberMap{
		startingSource: int(5),
		startingDest:   int(1),
		valueRange:     int(2),
	})
	expected.mapp = append(expected.mapp, numberMap{
		startingSource: int(20),
		startingDest:   int(10),
		valueRange:     int(1),
	})

	if !reflect.DeepEqual(sourceDestmap, expected) {
		t.Errorf(" want: %+v, got  %+v", expected, sourceDestmap)
	}
}

func TestTranslate(t *testing.T) {
	//var input = base.GetInput("./input.txt")
	sdm := makeSourceDestMap("seed", "soil")
	sdm.mapp = append(sdm.mapp, numberMap{
		startingSource: int(5),
		startingDest:   int(1),
		valueRange:     int(2),
	})
	sdm.mapp = append(sdm.mapp, numberMap{
		startingSource: int(20),
		startingDest:   int(10),
		valueRange:     int(1),
	})

	if sdm.translate(5) != 1 {
		t.Errorf(" want: %d got  %d", 1, sdm.translate(5))
	}

	if sdm.translate(6) != 2 {
		t.Errorf(" want: %d got  %d", 2, sdm.translate(6))
	}

	if sdm.translate(7) != 7 {
		t.Errorf(" want: %d got  %d", 7, sdm.translate(7))
	}

	if sdm.translate(10) != 10 {
		t.Errorf(" want: %d got  %d", 10, sdm.translate(10))
	}

	if sdm.translate(20) != 10 {
		t.Errorf(" want: %d got  %d", 10, sdm.translate(20))
	}

	if sdm.translate(21) != 21 {
		t.Errorf(" want: %d got  %d", 21, sdm.translate(21))
	}
}

func TestGetLocationFromSeed(t *testing.T) {
	var input []string
	input = append(input, "seeds: 1 2 3 4")
	input = append(input, "")
	input = append(input, "seed-to-soil map:")
	input = append(input, "1 5 2")
	input = append(input, "")
	input = append(input, "soil-to-location map:")
	input = append(input, "10 20 4")

	_, mapSections := getSeedAndMapSections(input)

	seed5Loc := getLocationFromSeed(5, mapSections)
	seed20Loc := getLocationFromSeed(20, mapSections)
	seed6Loc := getLocationFromSeed(6, mapSections)

	if seed5Loc != 1 {
		t.Errorf("seeds erradas! want: 1, got %d", seed5Loc)
	}
	if seed20Loc != 10 {
		t.Errorf("seeds erradas! want: 10, got %d", seed20Loc)
	}
	if seed6Loc != 2 {
		t.Errorf("seeds erradas! want: 2, got %d", seed6Loc)
	}
}

func TestLowestLocationFromSeedAndRange(t *testing.T) {
	var input []string
	input = append(input, "seeds: 1 2 3 4")
	input = append(input, "")
	input = append(input, "seed-to-soil map:")
	input = append(input, "150 5 2")
	input = append(input, "")
	input = append(input, "soil-to-location map:")
	input = append(input, "1 500 4")

	_, mapSections := getSeedAndMapSections(input)

	seed3_40Loc := divideByElevenRecur(5, 4000, mapSections)

	if seed3_40Loc != 1 {
		t.Errorf("seeds erradas! want: 10, got %d", seed3_40Loc)
	}
}

/*
get seed numbers
get map sections



for each seed ->
	get map for seed -> get map for dest recursive -> get location


find lowest location

seeds: 929142010 467769747 2497466808 210166838 3768123711 33216796 1609270159 86969850 199555506 378609832 1840685500 314009711 1740069852 36868255 2161129344 170490105 2869967743 265455365 3984276455 31190888

seed-to-soil map:
3788621315 24578909 268976974
3633843608 2672619957 154777707
1562003446 2827397664 767899879
2618130896 293555883 1015712712
178572254 3595297543 462300746
640873000 1553961386 921130446
2373438105 1435027522 118933864
2492371969 1309268595 125758927
2329903325 2629085177 43534780
24578909 2475091832 153993345

*/
