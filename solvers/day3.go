package days

import (
	"math"
	"strconv"
	"strings"
	"unicode"

	utils "github.com/chriscamarillo/advent-of-code-2023/common"
)

type Number struct {
	StartCol int
	EndCol   int
	Row      int
	Value    int
}

type Symbol struct {
	Col   int
	Row   int
	Value rune
}

func Day3Part1Solver() string {
	input := utils.FetchAOCInput(3)

	numbers, symbols := parseInput(input)

	sum := 0
	for _, n := range numbers {
		for _, s := range symbols {
			if areNeighbors(n, s) {
				sum += n.Value
			}
		}
	}

	answer := strconv.Itoa(sum)
	return answer
}

func Day3Part2Solver() string {
	input := utils.FetchAOCInput(3)

	numbers, symbols := parseInput(input)

	neighborMap := make(map[Symbol][]Number)

	for _, n := range numbers {
		for _, s := range symbols {
			if areNeighbors(n, s) && s.Value == '*' {
				neighborMap[s] = append(neighborMap[s], n)
			}
		}
	}

	sum := 0

	for _, parts := range neighborMap {
		if len(parts) == 2 {
			gearRatio := parts[0].Value * parts[1].Value
			sum += gearRatio
		}
	}

	answer := strconv.Itoa(sum)
	return answer
}

func parseInput(input string) ([]Number, []Symbol) {
	row := 0
	col := 0

	startCol := 0
	endCol := 0

	var numbers []Number
	var symbols []Symbol
	var digits strings.Builder
	var inDigit bool

	for _, r := range input {
		if unicode.IsDigit(r) {
			if !inDigit {
				startCol = col
				endCol = col
			} else {
				endCol++
			}

			inDigit = true
			digits.WriteRune(r)
		} else {
			if inDigit {
				value, _ := strconv.Atoi(digits.String())
				number := Number{
					startCol,
					endCol,
					row,
					value,
				}
				numbers = append(numbers, number)
				digits.Reset()
			}

			inDigit = false
			if r == '\n' {
				row++
				col = 0
			}
			if r != '.' && r != '\n' {
				symbol := Symbol{col, row, r}
				symbols = append(symbols, symbol)
			}
		}
		col++
	}
	return numbers, symbols
}

func areNeighbors(n Number, s Symbol) bool {
	if math.Abs(float64(n.Row-s.Row)) > 1 {
		return false
	}
	return s.Col >= n.StartCol-1 && s.Col <= n.EndCol+1
}
