package days

import (
	"math"
	"strconv"
	"strings"

	utils "github.com/chriscamarillo/advent-of-code-2023/common"
	"github.com/scylladb/go-set"
	"github.com/scylladb/go-set/strset"
)

func Day4Part1Solver() string {
	input := utils.FetchAOCInput(4)
	input = strings.Trim(input, "\n")

	scratchOffs := parseCardRows(input)
	sum := 0.0

	for _, scratchOff := range scratchOffs {
		common := strset.Intersection(scratchOff.playerNumbers, scratchOff.winningNumbers)
		finds := common.Size()

		if finds > 0 {
			sum += math.Pow(2, float64(finds-1))
		}
	}

	answer := strconv.Itoa(int(sum))
	return answer
}

func Day4Part2Solver() string {
	input := utils.FetchAOCInput(4)
	input = strings.Trim(input, "\n")

	scratchOffs := parseCardRows(input)
	sum := 0

	var cardCopies []int
	var cardFinds []int

	for _, scratchOff := range scratchOffs {
		common := strset.Intersection(scratchOff.playerNumbers, scratchOff.winningNumbers)
		cardFinds = append(cardFinds, common.Size())
		cardCopies = append(cardCopies, 1)
	}

	for i := range scratchOffs {
		finds := cardFinds[i]
		amount := cardCopies[i]
		for j := 0; j < finds; j++ {
			cardCopies[i+j+1] += amount
		}
	}

	for _, cardCount := range cardCopies {
		sum += cardCount
	}

	answer := strconv.Itoa(int(sum))
	return answer
}

func parseCardRows(input string) []ScratchOff {
	lines := strings.Split(input, "\n")

	var cards []ScratchOff

	for _, line := range lines {
		playerNumbers := set.NewStringSet()
		winningNumbers := set.NewStringSet()

		scratchNumbers := strings.Split(line, ":")[1]

		playerRow := strings.Split(scratchNumbers, "|")[0]
		winningRow := strings.Split(scratchNumbers, "|")[1]

		for _, n := range strings.Fields(playerRow) {
			playerNumbers.Add(n)
		}

		for _, n := range strings.Fields(winningRow) {
			winningNumbers.Add(n)
		}

		card := ScratchOff{
			playerNumbers,
			winningNumbers,
		}
		cards = append(cards, card)
	}

	return cards
}

type ScratchOff struct {
	playerNumbers  *strset.Set
	winningNumbers *strset.Set
}
