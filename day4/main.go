package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id          int
	instances   int
	winningNums map[int]bool
	actualNums  map[int]bool
}

func GameFactory(gameNumStr string, winning string, actual string) Game {
	winningParsed := make(map[int]bool)
	actualParsed := make(map[int]bool)

	winningSingleSpaced := strings.ReplaceAll(winning, "  ", " ")
	winningTrimmed := strings.TrimSpace(winningSingleSpaced)
	for _, num := range strings.Split(winningTrimmed, " ") {
		trimmed := strings.TrimSpace(num)
		p, err := strconv.Atoi(trimmed)
		if err != nil {
			fmt.Println("Error parsing winning numbers")
			panic(err)
		}
		winningParsed[p] = true
	}

	actualSingleSpaced := strings.ReplaceAll(actual, "  ", " ")
	actualTrimmed := strings.TrimSpace(actualSingleSpaced)
	for _, num := range strings.Split(actualTrimmed, " ") {
		trimmed := strings.TrimSpace(num)
		p, err := strconv.Atoi(trimmed)
		if err != nil {
			fmt.Println("Error parsing winning numbers")
			panic(err)
		}
		actualParsed[p] = true
	}

	headerSplit := strings.Split(gameNumStr, " ")
	gameNum, err := strconv.Atoi(headerSplit[len(headerSplit)-1])
	if err != nil {
		fmt.Println("Error parsing game number")
		panic(err)
	}

	return Game{gameNum, 1, winningParsed, actualParsed}
}

func parseGame(line string) Game {
	headerSplit := strings.Split(line, ":")
	gameSplit := strings.Split(headerSplit[1], "|")
	game := GameFactory(headerSplit[0], gameSplit[0], gameSplit[1])

	return game
}

func partOne(scanner *bufio.Scanner) {
	games := make([]Game, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, parseGame(line))
	}

	sum := 0
	for _, game := range games {
		pow := 0
		for num, _ := range game.actualNums {
			// set membership
			if game.winningNums[num] {
				if pow == 0 {
					// base case
					pow = 1
				} else {
					// nth case
					pow *= 2
				}
			}
		}
		fmt.Println(game.id, pow)
		sum += pow
	}
	fmt.Println("final: ", sum)
}

func partTwo(scanner *bufio.Scanner) {
	games := make([]*Game, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		games = append(games, &game)
	}

	for idx, game := range games {
		wins := 0
		for num, _ := range game.actualNums {
			if game.winningNums[num] {
				wins++
			}
		}
		for _, targetGame := range games[idx+1 : idx+wins+1] {
			targetGame.instances += game.instances
		}
	}

	sum := 0
	for _, game := range games {
		sum += game.instances
	}

	fmt.Println("final: ", sum)
}

func main() {
	infile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer infile.Close()

	scanner := bufio.NewScanner(infile)
	scanner.Split(bufio.ScanLines)

	doOne := false

	if doOne {
		partOne(scanner)
	} else {
		partTwo(scanner)
	}

}
