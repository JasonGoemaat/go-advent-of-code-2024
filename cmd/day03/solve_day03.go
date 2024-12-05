package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
)

func calculate(text string) int {
	rx := regexp.MustCompile("mul\\((\\d\\d??\\d??),(\\d\\d??\\d??)\\)")
	results := rx.FindAllStringSubmatch(text, -1)
	total := 0
	for _, match := range results {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			util.MyLog("ERROR converting %s\n", match[1])
		}
		b, _ := strconv.Atoi(match[2])
		if err != nil {
			util.MyLog("ERROR converting %s\n", match[2])
		}
		product := a * b
		total += a * b
		util.MyLog("%s: %d * %d = %d  (running total %d)\n", match[0], a, b, product, total)
	}
	return total
}

func SolveDay03(filePath string) int {
	line := util.LoadString(filePath)
	return calculate(line)
}

func SolveDay03Part2(filePath string) int {
	line := util.LoadString(filePath)
	dos := strings.Split(line, "do()")
	util.MyLog("dos: %v\n", dos)
	total := 0
	for _, s := range dos {
		split := strings.Split(s, "don't()")
		util.MyLog("split: %v\n", split)
		total += calculate(split[0])
	}
	return total
}
