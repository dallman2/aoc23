package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sample struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id      int
	samples []Sample
}

func SampleFactory(sample string) Sample {
	templ := Sample{-1, -1, -1}

	entries := strings.Split(sample, ", ")
	for _, entry := range entries {
		entryTuple := strings.Split(entry, " ")
		color := entryTuple[1]
		count, err := strconv.Atoi(entryTuple[0])
		if err != nil {
			fmt.Println("Error converting string to int")
		}

		switch color {
		case "blue":
			templ.blue = count
		case "green":
			templ.green = count
		case "red":
			templ.red = count
		}
	}

	return templ
}

func parseGame(game string) Game {
	// Split the game into the header and the list of samples
	topSplit := strings.Split(game, ": ")
	// split the list of samples into individual samples
	sampleSplit := strings.Split(topSplit[1], "; ")
	// grab the id out of the header
	id, err := strconv.Atoi(strings.Split(topSplit[0], " ")[1])
	if err != nil {
		fmt.Println("Error converting string to int")
	}

	g := Game{id, []Sample{}}

	for _, sample := range sampleSplit {
		g.samples = append(g.samples, SampleFactory(sample))
	}
	return g
}

func partOne(scanner *bufio.Scanner) {
	goodGames := make([]Game, 100)
	thresholdSample := Sample{12, 13, 14}

	for scanner.Scan() {
		g := parseGame(scanner.Text())

		good := true

		for _, sample := range g.samples {
			if thresholdSample.red >= sample.red && thresholdSample.green >= sample.green && thresholdSample.blue >= sample.blue {
				continue
			} else {
				good = false
				break
			}
		}

		if good {
			goodGames = append(goodGames, g)
		}
	}

	sum := 0

	for _, game := range goodGames {
		fmt.Println(game.id)
		sum += game.id
	}

	fmt.Println("part one")
	fmt.Println(sum)
}

func partTwo(scanner *bufio.Scanner) {
	sum := 0

	for scanner.Scan() {
		minCubes := Sample{0, 0, 0}
		g := parseGame(scanner.Text())

		for _, sample := range g.samples {
			if sample.red > minCubes.red {
				minCubes.red = sample.red
			}
			if sample.green > minCubes.green {
				minCubes.green = sample.green
			}
			if sample.blue > minCubes.blue {
				minCubes.blue = sample.blue
			}
		}

		sum += minCubes.red * minCubes.green * minCubes.blue

		fmt.Println(g.id, minCubes)
	}

	fmt.Println("part two")
	fmt.Println(sum)
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
