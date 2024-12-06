package day06

import (
	"testing"

	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
)

func TestPart1Sample(t *testing.T) {
	content := util.LoadString("c:/git/go/go-advent-of-code-2024/cmd/day06/data/sample.txt")
	t.Log(SolvePart1(content))
}
