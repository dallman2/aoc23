package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Mapping struct {
	sourceStart int
	destStart   int
	length      int
}
type Block struct {
	mappings []Mapping
	label    string
}

func calculateLocations(seedRange [2]int, blockList *[]Block, channel chan int, wg *sync.WaitGroup) {
	locations := make([]int, 0, seedRange[1])

	fmt.Println("starting: ", seedRange)

	for i := seedRange[0]; i < seedRange[0]+seedRange[1]; i++ {
		tformed := i
		for _, block := range *blockList {
			found := false
			for _, mapping := range block.mappings {
				if !found && tformed >= mapping.sourceStart && tformed < (mapping.sourceStart+mapping.length) {
					tformed = mapping.destStart + (tformed - mapping.sourceStart)
					found = true
				}
			}
		}

		locations = append(locations, tformed)
	}
	min := slices.Min(locations)

	fmt.Println("done: ", seedRange, min)
	channel <- min

	wg.Done()
}

func partOne(scanner *bufio.Scanner) {
	block := 0
	blocks := make([][]string, 0, 10)
	blocks = append(blocks, make([]string, 0, 30))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			block++
			blocks = append(blocks, make([]string, 0, 30))
		} else {
			blocks[block] = append(blocks[block], line)
		}
	}
	// fmt.Println(blocks)

	seedsStr, blocks := blocks[0], blocks[1:]

	seedStrSplit := strings.Split(seedsStr[0], " ")
	_, seedStrSplit = seedStrSplit[0], seedStrSplit[1:]
	seeds := make([]int, 0, 20)

	for _, seed := range seedStrSplit {
		num, err := strconv.Atoi(strings.TrimSpace(seed))
		if err != nil {
			fmt.Println("Error converting seed to int")
			panic(err)
		}
		seeds = append(seeds, num)
	}

	blockList := make([]Block, 0, 10)

	for _, block := range blocks {
		mappings := make([]Mapping, 0, 30)
		header, headless := block[0], block[1:]
		label := strings.Split(header, " ")[0]

		for _, line := range headless {
			entries := strings.Split(line, " ")
			dst, dstErr := strconv.Atoi(strings.TrimSpace(entries[0]))
			src, srcErr := strconv.Atoi(strings.TrimSpace(entries[1]))
			r, rErr := strconv.Atoi(strings.TrimSpace(entries[2]))
			if srcErr != nil || dstErr != nil || rErr != nil {
				fmt.Println("Error converting mapping to int")
				err := fmt.Errorf("srcErr: %v, dstErr: %v, rErr: %v", srcErr, dstErr, rErr)
				panic(err)
			}
			mappings = append(mappings, Mapping{src, dst, r})
		}
		blockList = append(blockList, Block{mappings, label})
	}
	// fmt.Println(blockList)

	locations := make([]int, 0, 20)

	for _, seed := range seeds {
		tformed := seed
		for _, block := range blockList {
			found := false
			for _, mapping := range block.mappings {
				if !found && tformed >= mapping.sourceStart && tformed < (mapping.sourceStart+mapping.length) {
					tformed = mapping.destStart + (tformed - mapping.sourceStart)
					found = true
				}
			}
		}

		locations = append(locations, tformed)
	}

	fmt.Println(locations)
	fmt.Println(slices.Min(locations))

}

func partTwo(scanner *bufio.Scanner) {
	block := 0
	blocks := make([][]string, 0, 10)
	blocks = append(blocks, make([]string, 0, 30))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			block++
			blocks = append(blocks, make([]string, 0, 30))
		} else {
			blocks[block] = append(blocks[block], line)
		}
	}
	// fmt.Println(blocks)

	seedsStr, blocks := blocks[0], blocks[1:]

	seedStrSplit := strings.Split(seedsStr[0], " ")
	_, seedStrSplit = seedStrSplit[0], seedStrSplit[1:]
	seeds := make([][2]int, 0, 20)

	for idx, seed := range seedStrSplit {
		if idx%2 != 0 {
			continue
		}
		base, err := strconv.Atoi(strings.TrimSpace(seed))
		length, err := strconv.Atoi(strings.TrimSpace(seedStrSplit[idx+1]))
		if err != nil {
			fmt.Println("Error converting seed to int")
			panic(err)
		}
		seeds = append(seeds, [2]int{base, length})
	}

	blockList := make([]Block, 0, 10)

	for _, block := range blocks {
		mappings := make([]Mapping, 0, 30)
		header, headless := block[0], block[1:]
		label := strings.Split(header, " ")[0]

		for _, line := range headless {
			entries := strings.Split(line, " ")
			dst, dstErr := strconv.Atoi(strings.TrimSpace(entries[0]))
			src, srcErr := strconv.Atoi(strings.TrimSpace(entries[1]))
			r, rErr := strconv.Atoi(strings.TrimSpace(entries[2]))
			if srcErr != nil || dstErr != nil || rErr != nil {
				fmt.Println("Error converting mapping to int")
				err := fmt.Errorf("srcErr: %v, dstErr: %v, rErr: %v", srcErr, dstErr, rErr)
				panic(err)
			}
			mappings = append(mappings, Mapping{src, dst, r})
		}
		blockList = append(blockList, Block{mappings, label})
	}

	locations := make([]int, 0, 20)

	var wg *sync.WaitGroup = &sync.WaitGroup{}
	channel := make(chan int, len(seeds))
	for _, seedRange := range seeds {
		wg.Add(1)
		go calculateLocations(seedRange, &blockList, channel, wg)

	}

	wg.Wait()
	close(channel)

	for res := range channel {
		locations = append(locations, res)
	}

	fmt.Println(locations)
	fmt.Println(len(locations))
	fmt.Println(slices.Min(locations))
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
