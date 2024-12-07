/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package day07

import (
	"fmt"

	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
	"github.com/spf13/cobra"
)

// day07Cmd represents the day07 command
var Day07Cmd = &cobra.Command{
	Use:   "day07",
	Short: "Day 7: Bridge Repair",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day07 called")
		content := util.GetContent()
		if content != "" {
			fmt.Println("Part 1:", SolvePart1(content))
			fmt.Println("Part 2:", SolvePart2(content))
			return
		}
		sample := util.LoadString("cmd/day07/data/sample.txt")
		input := util.LoadString("cmd/day07/data/input.txt")
		fmt.Println("Part 1 Sample:", SolvePart1(sample))
		fmt.Println("Part 1 Input :", SolvePart1(input))
		fmt.Println("Part 2 Sample:", SolvePart2(sample))
		fmt.Println("Part 2 Input :", SolvePart2(input))
	},
}
