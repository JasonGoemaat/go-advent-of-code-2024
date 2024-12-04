/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"slices"

	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
	"github.com/spf13/cobra"
)

// day01Cmd represents the day01 command
var day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: "Day 01: Historian Hysteria",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day01:")
		fmt.Println("Solution (part 1 sample):", solve(filepath.Join("data/day01/sample.txt")))
		fmt.Println("Solution (part 1 input ):", solve(filepath.Join("data/day01/input.txt")))
		fmt.Println("Solution (part 2 sample):", solve2(filepath.Join("data/day01/sample.txt")))
		fmt.Println("Solution (part 2 input ):", solve2(filepath.Join("data/day01/input.txt")))
	},
}

func init() {
	rootCmd.AddCommand(day01Cmd)
}

// like 'solve', but shows the parsing functions working
func info(filePath string) string {
	fmt.Println("LoadString:")
	contents := util.LoadString(filePath)
	fmt.Println(contents)

	fmt.Println("")
	fmt.Println("LoadLines:")
	lines := util.LoadLines(filePath)
	for i, line := range lines {
		fmt.Printf("%d: %s\n", i, line)
	}

	fmt.Println("")
	fmt.Println("LoadNumbers:")
	numbers := util.LoadNumbers(filePath)
	for i, arr := range numbers {
		fmt.Printf("%d: %v\n", i, arr)
	}

	return "answer"
}

func solve2(filePath string) int {
	numbers := util.LoadNumbers(filePath)
	a := takeElement(numbers, 0)
	b := takeElement(numbers, 1)
	m := make(map[int]int)
	for _, v := range b {
		m[v]++
	}

	total := 0
	for _, v := range a {
		total += v * m[v]
	}

	return total
}

func solve(filePath string) int {
	numbers := util.LoadNumbers(filePath)
	a := takeElement(numbers, 0)
	b := takeElement(numbers, 1)

	// each method, use sort.Ints which is built for it
	// sort.Ints(a)
	// sort.Ints(b)

	// sort.Slice(a, func(i int, j int) bool {
	// 	return a[i] < a[j]
	// })
	// sort.Slice(b, func(i int, j int) bool {
	// 	return b[i] < b[j]
	// })

	// myCompare := func(a, b int) int {
	// 	return cmp.Compare(a, b)
	// }
	// slices.SortFunc(a, myCompare)
	// slices.SortFunc(b, myCompare)

	// this is easy, use 'slice' package
	slices.Sort(a)
	slices.Sort(b)
	total := 0
	distance := func(i, j int) int {
		if i < j {
			return j - i
		}
		return i - j
	}
	for i, v := range a {
		total += distance(v, b[i])
	}
	return total
}

func takeElement(numbers [][]int, index int) []int {
	// I could start with an empty slice and use append, but
	// if you know the size ahead of time I think there's a
	// function called make(), looking that up...
	// didn't find a good explanation quickly, but let's try it
	var a = make([]int, len(numbers))
	for i, n := range numbers {
		a[i] = n[index]
	}
	return a
}
