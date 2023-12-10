package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func partOne(scanner *bufio.Scanner) {
	scanner.Scan()
	instLine := scanner.Text()
	scanner.Scan()
	scanner.Text()

	nodeMapping := make(map[string][2]string)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "=")
		node := strings.TrimSpace(split[0])
		children := strings.Split(split[1], ",")
		left := strings.TrimSpace(strings.TrimLeft(strings.TrimSpace(children[0]), "("))
		right := strings.TrimSpace(strings.TrimRight(strings.TrimSpace(children[1]), ")"))

		nodeMapping[node] = [2]string{left, right}
	}

	fmt.Println(instLine)
	fmt.Println(nodeMapping)

	startNode := "AAA"
	pos := 0
	inst := instLine[pos]
	mod := len(instLine)
	for startNode != "ZZZ" {
		inst = instLine[pos%mod]
		if inst == 'L' {
			startNode = nodeMapping[startNode][0]
		} else {
			startNode = nodeMapping[startNode][1]
		}
		pos++
	}

	fmt.Println(startNode)
	fmt.Println(pos)

}

func partTwo(scanner *bufio.Scanner) {
	scanner.Scan()
	instLine := scanner.Text()
	scanner.Scan()
	scanner.Text()

	nodeMapping := make(map[string][2]string)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "=")
		node := strings.TrimSpace(split[0])
		children := strings.Split(split[1], ",")
		left := strings.TrimSpace(strings.TrimLeft(strings.TrimSpace(children[0]), "("))
		right := strings.TrimSpace(strings.TrimRight(strings.TrimSpace(children[1]), ")"))

		nodeMapping[node] = [2]string{left, right}
	}

	fmt.Println(instLine)
	fmt.Println(nodeMapping)

	startNodes := make([]string, 0, len(nodeMapping))

	for k, _ := range nodeMapping {
		if strings.HasSuffix(k, "A") {
			startNodes = append(startNodes, k)
		}
	}

	fmt.Println(startNodes)

	stepList := make([]int, 0, len(startNodes))

	for _, v := range startNodes {
		startNode := v
		pos := 0
		inst := instLine[pos]
		mod := len(instLine)
		for !strings.HasSuffix(startNode, "Z") {
			inst = instLine[pos%mod]
			if inst == 'L' {
				startNode = nodeMapping[startNode][0]
			} else {
				startNode = nodeMapping[startNode][1]
			}
			pos++
		}
		stepList = append(stepList, pos)
	}

	fmt.Println(startNodes)
	fmt.Println(stepList)

	fmt.Println(LCM(stepList[0], stepList[1], stepList...))

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
