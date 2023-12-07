package days

import (
	"fmt"
	utils "github.com/chriscamarillo/advent-of-code-2023/common"
	"strconv"
	"strings"
)

const bagRedAmount int = 12
const bagGreenAmount int = 13
const bagBlueAmount int = 14

type Game struct {
	id       int
	maxRed   int
	maxGreen int
	maxBlue  int
}

func Day2Part1Solver() string {
	input := utils.FetchAOCInput(2)
	games := collectGameData(input)

	idSum := 0
	// Figure out which games are possible
	for _, game := range games {
		if game.maxRed <= bagRedAmount &&
			game.maxGreen <= bagGreenAmount &&
			game.maxBlue <= bagBlueAmount {
			idSum += game.id
			fmt.Println("This game is possible: ", game)
		}
	}

	answer := strconv.Itoa(idSum)
	return answer
}

func Day2Part2Solver() string {
	input := utils.FetchAOCInput(2)
	games := collectGameData(input)

	idSum := 0
	// Figure out which games are possible
	for _, game := range games {
		power := game.maxRed * game.maxGreen * game.maxBlue
		idSum += power
		fmt.Printf("%s - power: %d\n", game, power)
	}

	answer := strconv.Itoa(idSum)
	return answer
}

func collectGameData(input string) []Game {
	var games []Game

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		game := parseGameLine(line)
		games = append(games, game)
	}

	return games
}

func parseGameLine(line string) Game {
	gameString := strings.Split(line, ":")[0]
	resultString := strings.Split(line, ":")[1]

	var game Game

	game.id, _ = strconv.Atoi(strings.Fields(gameString)[1])

	hands := strings.Split(resultString, ";")
	for _, hand := range hands {
		pairs := strings.Split(hand, ",")

		for _, pair := range pairs {
			parts := strings.Fields(pair)
			amount, _ := strconv.Atoi(parts[0])

			color := parts[1]
			switch color {
			case "red":
				game.maxRed = max(amount, game.maxRed)
			case "green":
				game.maxGreen = max(amount, game.maxGreen)
			case "blue":
				game.maxBlue = max(amount, game.maxBlue)
			}
		}
	}
	return game
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (g Game) String() string {
	return fmt.Sprintf("Game(Id: %d, Red: %d, Green: %d, Blue: %d)",
		g.id, g.maxRed, g.maxGreen, g.maxBlue)
}
