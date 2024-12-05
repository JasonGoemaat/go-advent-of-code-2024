/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day05"
	"github.com/spf13/cobra"
)

// day05Cmd represents the day05 command
var day05Cmd = &cobra.Command{
	Use:   "day05",
	Short: "Day 5: Print Queue",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day05 called")
		fmt.Println("sample part1:", day05.SolveDay05("data/day05/sample.txt"))
		// fmt.Println("input part1 :", day05.SolveDay05("data/day05/input.txt"))
		// fmt.Println("sample part1:", day05.SolveDay05Part2("data/day05/sample.txt"))
		// fmt.Println("input part1 :", day05.SolveDay05Part2("data/day05/input.txt"))
	},
}

func init() {
	rootCmd.AddCommand(day05Cmd)
}
