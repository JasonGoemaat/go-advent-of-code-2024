/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package day04

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day04Cmd represents the day04 command
var Day04Cmd = &cobra.Command{
	Use:   "day04",
	Short: "Day 4: Ceres Search",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day04 called")
		// fmt.Println("sample part1:", day04.SolveDay04("cmd/day04/data/sample.txt"))
		// fmt.Println("input part1 :", day04.SolveDay04("cmd/day04/data/input.txt"))
		fmt.Println("sample part1:", SolveDay04Part2("cmd/day04/data/sample.txt"))
		fmt.Println("input part1 :", SolveDay04Part2("cmd/day04/data/input.txt"))
	},
}
