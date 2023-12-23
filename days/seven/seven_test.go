package days

import (
	base "aoc2023"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func solution(input []string) string {
	sum := 0
	var handBids []HandBid

	for _, line := range input {
		parts := strings.Split(line, " ")

		bid, _ := strconv.ParseInt(parts[1], 10, 64)
		handBid := HandBid{parts[0], int(bid)}
		handBids = append(handBids, handBid)
	}

	sort.SliceStable(handBids, func(i int, j int) bool {
		return compare(handBids[i].hand, handBids[j].hand) != 1
	})

	fmt.Println(handBids)

	for i, handBid := range handBids {
		sum += handBid.bid * (i + 1)
	}
	return strconv.FormatInt(int64(sum), 10)
}

func compare(hand1 string, hand2 string) int {
	highCardMap := make(map[string]int)
	highCardMap["2"] = 1
	highCardMap["3"] = 2
	highCardMap["4"] = 3
	highCardMap["5"] = 4
	highCardMap["6"] = 5
	highCardMap["7"] = 6
	highCardMap["8"] = 7
	highCardMap["9"] = 8
	highCardMap["T"] = 9
	highCardMap["J"] = 0
	highCardMap["Q"] = 11
	highCardMap["K"] = 12
	highCardMap["A"] = 13

	hand1Type := getHandType(hand1)
	hand2Type := getHandType(hand2)

	if hand1Type != hand2Type {
		if hand1Type > hand2Type {
			return 1
		}

		return -1
	}

	for i, card := range strings.Split(hand1, "") {
		if []rune(card)[0] != []rune(hand2)[i] {
			if highCardMap[card] > highCardMap[string([]rune(hand2)[i])] {
				return 1
			}
			return -1
		}
	}
	return 1
}

type HandBid struct {
	hand string
	bid  int
}

func getHandType(hand string) int {
	handMap := make(map[string]int)
	js := 0

	for _, card := range strings.Split(hand, "") {
		val, ok := handMap[card]

		if card == "J" {
			js += 1
		}

		if !ok {
			handMap[card] = 1
			continue
		}
		if card == "J" {
			continue
		}

		handMap[card] = 1 + val
	}

	var pairs []int
	for _, value := range handMap {
		if value > 0 {
			pairs = append(pairs, value)
		}
	}

	sort.Ints(pairs)

	pairs[len(pairs)-1] = js + pairs[len(pairs)-1]

	fmt.Println(hand, pairs)

	if pairs[len(pairs)-1] == 5 || js == 5 || js == 4 {
		return 7
	}

	if pairs[len(pairs)-1] == 4 || js == 3 {
		return 6
	}

	if pairs[len(pairs)-1] == 3 && pairs[len(pairs)-2] == 2 {
		return 5
	}

	if pairs[len(pairs)-1] == 3 {
		return 4
	}

	if pairs[len(pairs)-1] == 2 && pairs[len(pairs)-2] == 2 {
		return 3
	}

	if pairs[len(pairs)-1] == 2 {
		return 2
	}

	return 1
}

func TestGetType(t *testing.T) {
	compare5 := getHandType("QQQJA")
	compare4 := getHandType("32T3K")
	compare3 := getHandType("KK677")

	if compare5 != 6 {
		t.Errorf("input errado! want: 6, got %d", compare5)
	}

	if compare4 != 2 {
		t.Errorf("input errado! want: 2, got %d", compare4)
	}

	if compare3 != 3 {
		t.Errorf("input errado! want: 3, got %d", compare3)
	}
}

func TestCompare(t *testing.T) {
	compare5 := compare("QQQJA", "T55J5")
	compare4 := compare("T55J5", "QQQJA")
	compare3 := compare("KK677", "32T3K")
	compare2 := compare("T55J5", "33322")

	if compare5 != 1 {
		t.Errorf("input errado! want: 1, got %d", compare5)
	}

	if compare4 != -1 {
		t.Errorf("input errado! want: -1, got %d", compare4)
	}

	if compare3 != 1 {
		t.Errorf("input errado! want: 1, got %d", compare3)
	}

	if compare2 != 1 {
		t.Errorf("input errado! want: 1, got %d", compare2)
	}
}

func TestSolution(t *testing.T) {
	var input = []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	var result = solution(input)

	if result != "5905" {
		t.Errorf("input errado! want: 5905, got %s", result)
	}
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "251421071" {
		t.Errorf("input errado! want: 251421071, got %s", result)
	}
}
