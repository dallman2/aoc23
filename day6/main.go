package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(scanner *bufio.Scanner) {

	lines := make([]string, 0, 2)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	times := lines[0]
	distances := lines[1]

	timesSplit := strings.Split(times, " ")
	distancesSplit := strings.Split(distances, " ")
	timesParsed := make([]int, 0, 10)
	timesFiltered := make([]string, 0, 10)
	distancesParsed := make([]int, 0, 10)
	distancesFiltered := make([]string, 0, 10)

	for _, tok := range timesSplit {
		if tok != "" {
			timesFiltered = append(timesFiltered, tok)
		}
	}
	for _, tok := range distancesSplit {
		if tok != "" {
			distancesFiltered = append(distancesFiltered, tok)
		}
	}
	_, timesFiltered = timesFiltered[0], timesFiltered[1:]
	_, distancesFiltered = distancesFiltered[0], distancesFiltered[1:]
	for _, t := range timesFiltered {
		time, _ := strconv.Atoi(t)
		timesParsed = append(timesParsed, time)
	}
	for _, d := range distancesFiltered {
		distance, _ := strconv.Atoi(d)
		distancesParsed = append(distancesParsed, distance)
	}

	fmt.Println(timesParsed)
	fmt.Println(distancesParsed)

	winsList := make([]int, 0, 50)
	for i := 0; i < len(timesParsed); i++ {
		time := timesParsed[i]
		distance := distancesParsed[i]
		fmt.Println(time, distance)

		wins := 0
		for t := 0; t < time; t++ {
			runtime := time - t
			d := runtime * t

			if d > distance {
				wins++
			}
		}
		winsList = append(winsList, wins)
	}

	product := 1
	for _, wins := range winsList {
		product *= wins
	}

	fmt.Println(winsList)
	fmt.Println(product)
}

func partTwo(scanner *bufio.Scanner) {

	lines := make([]string, 0, 2)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	times := lines[0]
	distances := lines[1]

	timesSplit := strings.Split(times, " ")
	distancesSplit := strings.Split(distances, " ")
	timesFiltered := make([]string, 0, 10)
	distancesFiltered := make([]string, 0, 10)

	for _, tok := range timesSplit {
		if tok != "" {
			timesFiltered = append(timesFiltered, tok)
		}
	}
	for _, tok := range distancesSplit {
		if tok != "" {
			distancesFiltered = append(distancesFiltered, tok)
		}
	}
	_, timesFiltered = timesFiltered[0], timesFiltered[1:]
	_, distancesFiltered = distancesFiltered[0], distancesFiltered[1:]

	fmt.Println(timesFiltered)
	fmt.Println(distancesFiltered)

	trueTime, _ := strconv.Atoi(strings.Join(timesFiltered, ""))
	trueDistance, _ := strconv.Atoi(strings.Join(distancesFiltered, ""))

	fmt.Println(trueTime)
	fmt.Println(trueDistance)

	wins := 0
	for t := 0; t < trueTime; t++ {
		runtime := trueTime - t
		d := runtime * t

		if d > trueDistance {
			wins++
		}
	}

	fmt.Println(wins)
}

func main() {
	infile, err := os.Open("input2.txt")
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
