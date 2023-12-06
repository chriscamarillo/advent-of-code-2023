package main

import (
	"fmt"
	day1 "github.com/chriscamarillo/advent-of-code-2023/solvers"
	"github.com/spf13/cobra"
)

type Solver func() string

var Solvers = map[string]Solver{
	"day1part1": day1.Part1Solver,
	"day1part2": day1.Part2Solver,
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