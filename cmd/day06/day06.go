/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package day06

import (
	"fmt"

	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
	"github.com/spf13/cobra"
)

// day06Cmd represents the day06 command
var Day06Cmd = &cobra.Command{
	Use:   "day06",
	Short: "Day 6: Guard Gallivant",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day06 called")
		content := util.GetContent()
		if content != "" {
			fmt.Println("Part 1:", SolvePart1(content))
			fmt.Println("Part 2:", SolvePart2(content))
			return
		}
		sample := util.LoadString("cmd/day06/data/sample.txt")
		input := util.LoadString("cmd/day06/data/input.txt")
		fmt.Println("Part 1 Sample:", SolvePart1(sample))
		fmt.Println("Part 1 Input :", SolvePart1(input))
		fmt.Println("Part 2 Sample:", SolvePart2(sample))
		fmt.Println("Part 2 Input :", SolvePart2(input))
	},
}
