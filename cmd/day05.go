/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day05"
	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
	"github.com/spf13/cobra"
)

// day05Cmd represents the day05 command
var day05Cmd = &cobra.Command{
	Use:   "day05",
	Short: "Day 5: Print Queue",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day05 called")
		sample := util.LoadString("data/day05/sample.txt")
		input := util.LoadString("data/day05/input.txt")
		fmt.Println("sample part1:", day05.SolveDay05(sample))
		fmt.Println("input  part1:", day05.SolveDay05(input))
		fmt.Println("sample part2:", day05.SolveDay05Part2(sample))
		fmt.Println("input  part2:", day05.SolveDay05Part2(input))
	},
}

func init() {
	rootCmd.AddCommand(day05Cmd)
}
