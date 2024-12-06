/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package day03

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day03Cmd represents the day03 command
var Day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "Day 3: Mull It Over",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day03 called")
		// fmt.Println("sample part1:", day03.SolveDay03("cmd/day03/data/sample.txt"))
		// fmt.Println("input part1 :", day03.SolveDay03("cmd/day03/data/input.txt"))
		fmt.Println("sample part2:", SolveDay03Part2("cmd/day03/data/sample_part2.txt"))
		fmt.Println("input part2 :", SolveDay03Part2("cmd/day03/data/input.txt"))
	},
}
