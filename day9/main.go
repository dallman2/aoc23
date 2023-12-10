package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sequence struct {
	sequence []int
	diffs    []int
	zeros    bool
	child    *Sequence
}

func SequenceFactory(nums []int) Sequence {

	diffs := make([]int, 0, len(nums)-1)

	for i := 1; i < len(nums); i++ {
		diffs = append(diffs, nums[i]-nums[i-1])
	}

	nonZero := false
	for _, diff := range diffs {
		if diff != 0 {
			nonZero = true
			break
		}
	}

	if nonZero {
		child := SequenceFactory(diffs)
		return Sequence{nums, diffs, !nonZero, &child}
	} else {
		return Sequence{nums, diffs, !nonZero, nil}
	}
}

func parseSequence(line string) Sequence {
	nums := make([]int, 0, 20)
	split := strings.Split(strings.TrimSpace(line), " ")
	for _, num := range split {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}
	return SequenceFactory(nums)
}

func partOne(scanner *bufio.Scanner) {

	sequences := make([]Sequence, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		sequences = append(sequences, parseSequence(line))
	}

	projected := make([]int, 0, 20)

	for _, seq := range sequences {
		nexts := make([]int, 0, 20)
		nexts = append(nexts, seq.sequence[len(seq.sequence)-1])
		for child := seq.child; child != nil; child = child.child {
			nexts = append(nexts, child.sequence[len(child.sequence)-1])
		}

		sum := 0
		for _, next := range nexts {
			sum += next
		}
		projected = append(projected, sum)
	}

	fmt.Println(projected)

	sum := 0
	for _, num := range projected {
		sum += num
	}
	fmt.Println(sum)

}

func partTwo(scanner *bufio.Scanner) {
	sequences := make([]Sequence, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		sequences = append(sequences, parseSequence(line))
	}

	projected := make([]int, 0, 20)

	for _, seq := range sequences {
		nexts := make([]int, 0, 20)
		nexts = append(nexts, seq.sequence[0])
		for child := seq.child; child != nil; child = child.child {
			nexts = append(nexts, child.sequence[0])
		}

		sum := 0
		// slices.Reverse(nexts)
		fmt.Println(nexts)
		for _, next := range nexts {
			sum += next
		}
		projected = append(projected, sum)
	}

	fmt.Println(projected)

	sum := 0
	for _, num := range projected {
		sum += num
	}
	fmt.Println(sum)
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
