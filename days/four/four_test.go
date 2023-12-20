package days

import (
	base "aoc2023"
	"strings"
	"testing"
)

func solution(input []string) string {

for _, line := range input {

	parts := strings.Split(line, ":")
	
}
}


func TestTest(t *testing.T) {
	var input = base.GetInput("./input.txt")

	var result = solution(input)

	if result != "jjeee" {
		t.Errorf("input errado! want: jjeee, got %s", result)
	}
}
