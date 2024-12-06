/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package day05

import (
	"fmt"

	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
	"github.com/spf13/cobra"
)

// day05Cmd represents the day05 command
var Day05Cmd = &cobra.Command{
	Use:   "day05",
	Short: "Day 5: Print Queue",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day05 called")
		content := util.GetContent()
		if content != "" {
			fmt.Println("Part 1:", SolveDay05(content))
			fmt.Println("Part 2:", SolveDay05Part2(content))
			return
		}
		sample := util.LoadString("cmd/day05/data/sample.txt")
		input := util.LoadString("cmd/day05/data/input.txt")
		fmt.Println("Part 1 Sample:", SolveDay05(sample))
		fmt.Println("Part 1 Input :", SolveDay05(input))
		fmt.Println("Part 2 Sample:", SolveDay05Part2(sample))
		fmt.Println("Part 2 Input :", SolveDay05Part2(input))
	},
}
