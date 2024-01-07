package days

import (
	base "aoc2023"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func solution(input []string) (sum int) {
	for _, line := range input {
		parts := strings.Split(line, " ")

		numStrs := strings.Split(parts[1], ",")

		var errors []int
		for _, numStr := range numStrs {
			num, _ := strconv.ParseInt(numStr, 10, 64)
			errors = append(errors, int(num))
		}

		iterSum := countWays(parts[0], errors)

		// fmt.Println(parts[0], errors, iterSum)
		sum += iterSum
	}

	return sum
}

func TestSolution(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}

	result := solution(input)

	if result != 21 {
		t.Errorf("input errado! want: 21, got %d", result)
	}

}
func solve(input []string) int {
	memory := SafeCounter{v: make(map[string]int)}
	var countWays func(str string, errors []int) int
	countWays = toMemoised(func(str string, errors []int) int {
		result := 0
		runes := []rune(str)

		if len(runes) == 0 {
			if len(errors) == 0 {
				return 1
			}

			return 0
		}

		if len(errors) == 0 {
			for _, myRune := range runes {
				if string(myRune) == "#" {
					return 0
				}
			}

			return 1
		}

		//fmt.Println(runes, errors, len(runes), sumArr(errors), sumArr(errors)+len(errors))
		if len(runes) < sumArr(errors)+len(errors)-1 {

			// if len(runes) == sumArr(errors) && len(errors) == 1 {
			// 	return 1
			// }

			return 0
		}

		if string(runes[0]) == "." {
			result = countWays(string(runes[1:]), errors)
			return result
		}

		if string(runes[0]) == "#" {
			// fmt.Println(runes, errors)
			if len(runes) == errors[0] {
				for i := 0; i < errors[0]; i++ {
					if string(runes[i]) == "." {
						return 0
					}
				}

				return 1
			}

			for i := 0; i < errors[0]; i++ {
				if string(runes[i]) == "." {
					return 0
				}
			}

			if string(runes[errors[0]]) == "#" {
				return 0
			}

			// if len(runes) == errors[0] && len(errors) == 1 {
			// 	for _, myRune := range runes {
			// 		if string(myRune) == "." {
			// 			//fmt.Println("result", 0, str, errors)
			// 			return 0
			// 		}
			// 	}
			// 	return 1
			// }

			result = countWays(string(runes[errors[0]+1:]), errors[1:])
			return result
		}

		result1 := countWays("#"+string([]rune(str)[1:]), errors)

		result2 := countWays("."+string([]rune(str)[1:]), errors)
		return result1 + result2

		//return countWays("#"+string([]rune(str)[1:]), errors) + countWays("."+string([]rune(str)[1:]), errors)
	}, &memory)

	calls := len(input)
	var results []int

	var wg sync.WaitGroup

	for _, line := range input {
		parts := strings.Split(line, " ")

		numStrs := strings.Split(parts[1], ",")

		var errors []int
		for _, numStr := range numStrs {
			num, _ := strconv.ParseInt(numStr, 10, 64)
			errors = append(errors, int(num))
		}

		countWays(parts[0], errors)

	}

}

// func TestCountWays(t *testing.T) {
// 	var result1 = countWays("???.###", []int{1, 1, 3})
// 	if result1 != 1 {
// 		t.Errorf("input errado! want: 1, got %d", result1)
// 	}

// 	var result2 = countWays(".??..??...?##", []int{1, 1, 3})
// 	if result2 != 4 {
// 		t.Errorf("input errado! want: 4, got %d", result2)
// 	}

// 	var result3 = countWays("?#?#?#?#?#?#?#?", []int{1, 3, 1, 6})
// 	if result3 != 1 {
// 		t.Errorf("input errado! want: 1, got %d", result3)
// 	}

// 	var result4 = countWays("????.#...#...", []int{4, 1, 1})
// 	if result4 != 1 {
// 		t.Errorf("input errado! want: 1, got %d", result4)
// 	}

// 	var result5 = countWays("????.######..#####.", []int{1, 6, 5})
// 	if result5 != 4 {
// 		t.Errorf("input errado! want: 4, got %d", result5)
// 	}

// 	var result6 = countWays("?###????????", []int{3, 2, 1})
// 	if result6 != 10 {
// 		t.Errorf("input errado! want: 10, got %d", result6)
// 	}

// 	var result7 = countWays(".???", []int{1})
// 	if result7 != 3 {
// 		t.Errorf("input errado! want: 3, got %d", result7)
// 	}

// 	/*?????

// 	  ##.#.
// 	  ##..#
// 	  .##.#

// 	*/

// 	var result8 = countWays("?????", []int{2, 1})
// 	if result8 != 3 {
// 		t.Errorf("input errado! want: 3, got %d", result8)
// 	}

// 	var result9 = countWays(".????", []int{2, 1})
// 	if result9 != 1 {
// 		t.Errorf("input errado! want: 1, got %d", result9)
// 	}

// 	var result10 = countWays("#?#????????.?#.", []int{4, 1, 2, 1})
// 	if result10 != 6 {
// 		t.Errorf("input errado! want: 6, got %d", result10)
// 	}
// }

func sumArr(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}

	return sum
}

func TestTestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	expected := []int{
		6, 15, 3, 2, 8, 5, 4, 4, 27, 1, 2, 1,
		6, 4, 4, 3, 2, 1, 3, 3, 4, 3, 2, 3,
		2, 13, 7, 2, 4, 9, 2, 4, 1, 2, 2, 25,
		2, 4, 5, 2, 86, 4, 1, 15, 42, 4, 8, 5,
		17, 2, 4, 10, 6, 6, 2, 3, 35, 6, 13, 3,
		4, 3, 10, 3, 2, 1, 5, 7, 20, 2, 1, 6,
		1, 14, 2, 6, 4, 4, 13, 6, 2, 2, 1, 35,
		3, 5, 16, 20, 1, 3, 2, 3, 22, 5, 8, 6,
		11, 7, 5, 4,
	}

	var nums []int
	for _, line := range input[0:100] {
		parts := strings.Split(line, " ")

		numStrs := strings.Split(parts[1], ",")

		var errors []int
		for _, numStr := range numStrs {
			num, _ := strconv.ParseInt(numStr, 10, 64)
			errors = append(errors, int(num))
		}

		iterSum := countWays(parts[0], errors)

		nums = append(nums, iterSum)
		// fmt.Println(parts[0], errors, iterSum)
		//sum += iterSum
	}

	for i, ways := range nums {
		if ways != expected[i] {
			t.Errorf("input errado! want: %d, got %d, index %d", expected[i], ways, i)
		}
	}

	// var result = solution(input)

	// if result != 7344 {
	// 	t.Errorf("input errado! want: 7344, got %d", result)
	// }
}

func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != 7344 {
		t.Errorf("input errado! want: 7344, got %d", result)
	}
}

func TestTest2(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution2(input)

	if result != 7344 {
		t.Errorf("input errado! want: 7344, got %d", result)
	}
}

func solution2(input []string) (sum int) {

	//waitGroup len lines
	//chan receiving results
	//mutex with memoised code

	memoised := make(map[string]int)

	for _, line := range input {
		parts := strings.Split(line, " ")

		strExpanded := parts[0] + parts[0] + parts[0] + parts[0] + parts[0]
		numsSExpanded := parts[1] + "," + parts[1] + "," + parts[1] + "," + parts[1] + "," + parts[1]

		numStrs := strings.Split(numsSExpanded, ",")

		var errors []int
		for _, numStr := range numStrs {
			num, _ := strconv.ParseInt(numStr, 10, 64)
			errors = append(errors, int(num))
		}

		iterSum := countWays(strExpanded, errors)

		// fmt.Println(parts[0], errors, iterSum)
		sum += iterSum
	}

	return sum
}

func generateKey(str string, nums []int) string {

	newStr := ""
	newStr += str
	for _, num := range nums {
		newStr += strconv.FormatInt(int64(num), 10) + ","
	}

	return newStr
}

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func toMemoised(fun func(str string, nums []int) int, sc *SafeCounter) func(str string, nums []int) int {
	fun2 := func(str string, nums []int) int {
		key := generateKey(str, nums)
		sc.mu.Lock()
		res, ok := sc.v[key]
		sc.mu.Unlock()
		if ok {
			return res
		}

		res2 := fun(str, nums)
		sc.mu.Lock()
		defer sc.mu.Unlock()
		sc.v[key] = res2

		return res2
	}

	return fun2
}

/* #?#????????.?#. [4 1 2 1] 10

####????????.?#. [4 1 2 1]
???????.?#. [1 2 1]
??????? [1 2]

#.##...
#..##..
#...##.
#....##
.#.##..
.#..##.
.#...##
..#.##.
..#..##
...#.##

#?#????????.?#. 4,1,2,1

*/
