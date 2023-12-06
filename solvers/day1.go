package day1

import (
	"fmt"
	utils "github.com/chriscamarillo/advent-of-code-2023/common"
	"strconv"
	"strings"
	"unicode"
)

var digitStringRuneMap = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func Part1Solver() string {
	input := utils.FetchAOCInput(1)
	answer := solve(input, false)

	return answer
}

func Part2Solver() string {
	input := utils.FetchAOCInput(1)
	answer := solve(input, true)

	return answer
}

func solve(input string, digitSpellCheckEnabled bool) string {
	sum := 0

	firstDigit := ' '
	lastDigit := ' '

	currentLineIndex := 0
	lastDigitIndex := -1

	var alphabetSoup strings.Builder

	for _, runeValue := range input {
		if runeValue == '\n' {
			if digitSpellCheckEnabled {
				digitFound, digitIndex := findSpelledDigit(alphabetSoup.String())
				if digitIndex > lastDigitIndex {
					lastDigit = digitFound
				}
			}

			if firstDigit != ' ' && lastDigit != ' ' {
				number, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
				sum += number
				fmt.Printf("first digit [%c] last digit [%c] = [%d], original line: [%s]\n",
					firstDigit,
					lastDigit,
					number,
					strings.TrimSpace(alphabetSoup.String()))
			}
			firstDigit = ' '
			lastDigit = ' '
			lastDigitIndex = -1
			currentLineIndex = 0

			// start recording a new line
			alphabetSoup.Reset()
		}

		isRuneDigit := unicode.IsDigit(runeValue)
		if digitSpellCheckEnabled {
			digitFound, _ := findSpelledDigit(alphabetSoup.String())
			if digitFound != ' ' && firstDigit == ' ' {
				firstDigit = digitFound
			}
		}

		if isRuneDigit {
			if firstDigit == ' ' {
				firstDigit = runeValue
			}
			lastDigit = runeValue
			lastDigitIndex = currentLineIndex
		}

		alphabetSoup.WriteRune(runeValue)
		currentLineIndex++
	}
	answer := strconv.Itoa(sum)
	return answer
}

func findSpelledDigit(line string) (rune, int) {
	lastDigit := ' '
	lastIndex := -1

	for key := range digitStringRuneMap {
		foundIndex := strings.LastIndex(line, key)
		if foundIndex > lastIndex {
			lastDigit = digitStringRuneMap[key]
			lastIndex = foundIndex
		}
	}

	return lastDigit, lastIndex
}
