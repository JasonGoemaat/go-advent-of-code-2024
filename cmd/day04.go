/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day04"
	"github.com/spf13/cobra"
)

// day04Cmd represents the day04 command
var day04Cmd = &cobra.Command{
	Use:   "day04",
	Short: "Day 4: Ceres Search",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day04 called")
		// fmt.Println("sample part1:", day04.SolveDay04("data/day04/sample.txt"))
		// fmt.Println("input part1 :", day04.SolveDay04("data/day04/input.txt"))
		fmt.Println("sample part1:", day04.SolveDay04Part2("data/day04/sample.txt"))
		fmt.Println("input part1 :", day04.SolveDay04Part2("data/day04/input.txt"))
	},
}

func init() {
	rootCmd.AddCommand(day04Cmd)
}
