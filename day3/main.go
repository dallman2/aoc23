package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

type Symbol struct {
	row       int
	col       int
	gear      bool
	neighbors []Number
}

type Number struct {
	row    int
	col    int
	length int
	value  int
}

func parseLine(line string, lineIdx int) ([]Symbol, []Number, int) {
	i := 0
	lineLen := len(line)

	symbols := make([]Symbol, 0, 10)
	numbers := make([]Number, 0, 10)

	for i < lineLen {
		char := line[i]
		if char == '.' {

			i++
		} else if unicode.IsDigit(rune(char)) {
			numLength := 1
			for i+numLength < lineLen && unicode.IsDigit(rune(line[i+numLength])) {
				numLength++
			}
			val, err := strconv.Atoi(line[i : i+numLength])
			if err != nil {
				fmt.Println("Error converting string to int")
			}

			num := Number{lineIdx, i, numLength, val}
			i += numLength
			numbers = append(numbers, num)
		} else if !unicode.IsDigit(rune(char)) {

			sym := Symbol{lineIdx, i, char == '*', make([]Number, 0, 4)}
			i++
			symbols = append(symbols, sym)
		}
	}

	return symbols, numbers, lineLen
}

func partOne(scanner *bufio.Scanner) {

	partSum := 0

	symbols := make([]Symbol, 0, 100)
	numbers := make([]Number, 0, 100)

	symbolLookup := make(map[int]map[int]Symbol, 0)

	lineIdx := 0
	lineLen := 0
	for scanner.Scan() {
		lineSymbols, lineNumbers, len := parseLine(scanner.Text(), lineIdx)

		symbols = append(symbols, lineSymbols...)
		numbers = append(numbers, lineNumbers...)

		lineLen = len
		lineIdx++
	}

	totalRows := lineIdx

	for _, sym := range symbols {
		if _, ok := symbolLookup[sym.row]; !ok {
			symbolLookup[sym.row] = make(map[int]Symbol, 0)
		}
		symbolLookup[sym.row][sym.col] = sym
	}

	for _, num := range numbers {
		posToCheck := make([][2]int, 0, 10)

		// start at neg one and continue until length + 1 to catch diagonals
		for i := -1; i <= num.length; i++ {
			// make sure our column is in bounds
			if !(num.col+i >= 0) || !(num.col+i < lineLen) {
				continue
			}
			if num.row > 0 {
				// make sure we can check the previous row
				posToCheck = append(posToCheck, [2]int{num.row - 1, num.col + i})
			}
			if num.row <= totalRows-2 {
				// make sure we can check the next row
				posToCheck = append(posToCheck, [2]int{num.row + 1, num.col + i})
			}
			posToCheck = append(posToCheck, [2]int{num.row, num.col + i})

		}

		found := false
		for _, pos := range posToCheck {
			if _, rowOk := symbolLookup[pos[0]]; !rowOk {
			} else if _, ok := symbolLookup[pos[0]][pos[1]]; !ok {
			} else {
				fmt.Println("found symbol at: ", pos[0], pos[1], symbolLookup[pos[0]][pos[1]], num)
				found = true
			}
		}
		if found {
			partSum += num.value
		}

	}

	fmt.Println("Part 1: ", partSum)

}

func partTwo(scanner *bufio.Scanner) {
	partSum := 0

	symbols := make([]Symbol, 0, 100)
	numbers := make([]Number, 0, 100)

	symbolLookup := make(map[int]map[int]Symbol, 0)

	lineIdx := 0
	lineLen := 0
	for scanner.Scan() {
		lineSymbols, lineNumbers, len := parseLine(scanner.Text(), lineIdx)

		symbols = append(symbols, lineSymbols...)
		numbers = append(numbers, lineNumbers...)

		lineLen = len
		lineIdx++
	}

	gears := make([]Symbol, 0, 100)

	for _, sym := range symbols {
		if sym.gear {
			gears = append(gears, sym)
		}
	}

	totalRows := lineIdx

	for _, sym := range gears {
		if _, ok := symbolLookup[sym.row]; !ok {
			symbolLookup[sym.row] = make(map[int]Symbol, 0)
		}
		symbolLookup[sym.row][sym.col] = sym
	}

	for _, num := range numbers {
		posToCheck := make([][2]int, 0, 10)

		// start at neg one and continue until length + 1 to catch diagonals
		for i := -1; i <= num.length; i++ {
			// make sure our column is in bounds
			if !(num.col+i >= 0) || !(num.col+i < lineLen) {
				continue
			}
			if num.row > 0 {
				// make sure we can check the previous row
				posToCheck = append(posToCheck, [2]int{num.row - 1, num.col + i})
			}
			if num.row <= totalRows-2 {
				// make sure we can check the next row
				posToCheck = append(posToCheck, [2]int{num.row + 1, num.col + i})
			}
			posToCheck = append(posToCheck, [2]int{num.row, num.col + i})

		}

		for _, pos := range posToCheck {
			if _, rowOk := symbolLookup[pos[0]]; !rowOk {
			} else if _, ok := symbolLookup[pos[0]][pos[1]]; !ok {
			} else {
				// fmt.Println("found symbol at: ", pos[0], pos[1], symbolLookup[pos[0]][pos[1]], num)
				sym := symbolLookup[pos[0]][pos[1]]
				if !slices.ContainsFunc(sym.neighbors, func(n Number) bool { return n.col == num.col && n.row == num.row }) {
					sym.neighbors = append(sym.neighbors, num)
				}
				symbolLookup[pos[0]][pos[1]] = sym
			}
		}
	}

	for _, symMap := range symbolLookup {
		for _, sym := range symMap {
			if len(sym.neighbors) == 2 {
				fmt.Println("found gear at: ", sym.row, sym.col, sym.neighbors[0], sym.neighbors[1])
				partSum += sym.neighbors[0].value * sym.neighbors[1].value
			}
		}
	}

	fmt.Println("Part 2: ", partSum)
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
