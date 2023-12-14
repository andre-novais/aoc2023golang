package base

import (
	"os"
	"strings"
	"testing"
)

func GetInput(path string) []string {
	var data, err = os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func TestTest(t *testing.T) {
	var input = GetInput("./test-input.txt")

	if input[0] != "uno" {
		t.Errorf("input errado!")
	}
}
