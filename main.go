package main

import (
	"fmt"

	days "github.com/chriscamarillo/advent-of-code-2023/solvers"
	"github.com/spf13/cobra"
)

type Solver func() string

var Solvers = map[string]Solver{
	"day1part1": days.Day1Part1Solver,
	"day1part2": days.Day1Part2Solver,
	"day2part1": days.Day2Part1Solver,
	"day2part2": days.Day2Part2Solver,
	"day3part1": days.Day3Part1Solver,
	"day3part2": days.Day3Part2Solver,
	"day4part1": days.Day4Part1Solver,
	"day4part2": days.Day4Part2Solver,
}

func main() {
	var problem string

	var rootCmd = &cobra.Command{
		Use:   "advent-of-code-2023",
		Short: "a CLI tool to solve advent of code problems",
	}

	var solveCmd = &cobra.Command{
		Use:   "solve",
		Short: "Solves a given problem based on the day and part ex: day1part1",
		Run: func(cmd *cobra.Command, args []string) {
			answer := Solvers[problem]()
			fmt.Println(answer)
		},
	}

	solveCmd.Flags().StringVar(&problem, "problem", "day1part1", "which day/part you'd like solved")
	solveCmd.MarkFlagRequired("problem")
	rootCmd.AddCommand(solveCmd)
	rootCmd.Execute()
}
