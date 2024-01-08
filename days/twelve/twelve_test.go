package days

import (
	base "aoc2023"
	"strconv"
	"strings"
	"sync"
	"testing"
)

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

		if len(runes) < sumArr(errors)+len(errors)-1 {
			return 0
		}

		if string(runes[0]) == "." {
			result = countWays(string(runes[1:]), errors)
			return result
		}

		if string(runes[0]) == "#" {
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

			result = countWays(string(runes[errors[0]+1:]), errors[1:])
			return result
		}

		result1 := countWays("#"+string([]rune(str)[1:]), errors)

		result2 := countWays("."+string([]rune(str)[1:]), errors)
		return result1 + result2

	}, &memory)

	result := 0

	var wg sync.WaitGroup
	output := make(chan int, 3)

	for _, line := range input {
		wg.Add(1)
		parts := strings.Split(line, " ")

		expandedStr := parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0]

		expandedErrors := parts[1] + "," + parts[1] + "," + parts[1] + "," + parts[1] + "," + parts[1]

		numStrs := strings.Split(expandedErrors, ",")

		var errors []int
		for _, numStr := range numStrs {
			num, _ := strconv.ParseInt(numStr, 10, 64)
			errors = append(errors, int(num))
		}

		go func() {
			num := countWays(expandedStr, errors)
			output <- num
		}()
	}

	go func() {
		for data := range output {
			result += data

			wg.Done()
		}
	}()

	wg.Wait()
	close(output)

	return result
}

func sumArr(ints []int) int {
	sum := 0
	for _, num := range ints {
		sum += num
	}
	return sum
}

func TestTest2(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solve(input)

	if result != 1088006519007 {
		t.Errorf("input errado! want: 1088006519007, got %d", result)
	}
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
