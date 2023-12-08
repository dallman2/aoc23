package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HandRank int

const (
	FIVE_KIND HandRank = iota
	FOUR_KIND
	FULL_HOUSE
	THREE_KIND
	TWO_PAIR
	PAIR
	HIGH_CARD
)

type Hand struct {
	bid    int
	cards  []rune
	counts map[rune]int
	rank   HandRank
}

var CardRank = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func rankHand(hand map[rune]int) HandRank {

	bestRank := HIGH_CARD
	// fmt.Println(hand)
	for _, v := range hand {
		// fmt.Println(v, bestRank)
		if v == 5 {
			bestRank = FIVE_KIND
		} else if v == 4 && bestRank > FOUR_KIND {
			bestRank = FOUR_KIND
		} else if v == 2 && bestRank == PAIR {
			bestRank = TWO_PAIR
		} else if v == 3 && bestRank == PAIR {
			bestRank = FULL_HOUSE
		} else if v == 2 && bestRank == THREE_KIND {
			bestRank = FULL_HOUSE
		} else if v == 3 && bestRank > THREE_KIND {
			bestRank = THREE_KIND
		} else if v == 2 && bestRank > PAIR {
			bestRank = PAIR
		}
	}

	// fmt.Println(bestRank)

	return bestRank
}

func parseHand(line string) Hand {
	split := strings.Split(line, " ")
	bid, _ := strconv.Atoi(split[1])

	cards := make([]rune, 0, 5)
	counts := make(map[rune]int)
	for _, c := range split[0] {
		cards = append(cards, c)
		counts[c]++
	}

	rank := rankHand(counts)

	return Hand{bid, cards, counts, rank}
}

func orderHand(a, b Hand) int {
	if a.rank == b.rank {
		for i := 0; i < 5; i++ {
			if a.cards[i] != b.cards[i] {
				return CardRank[a.cards[i]] - CardRank[b.cards[i]]
			}
		}
	}
	return int(b.rank - a.rank)
}

func partOne(scanner *bufio.Scanner) {

	hands := make([]Hand, 0, 100)

	fmt.Println("unsorted")
	for scanner.Scan() {
		line := scanner.Text()
		hands = append(hands, parseHand(line))
	}

	slices.SortFunc(hands, orderHand)

	fmt.Println("sorted")
	sum := 0
	for idx, hand := range hands {
		fmt.Println(string(hand.cards), hand.rank, idx+1, hand.bid, sum, "\t", hand.counts)
		sum += hand.bid * (idx + 1)
	}

	fmt.Println("sum", sum)

}

func partTwo(scanner *bufio.Scanner) {

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

	doOne := true

	if doOne {
		partOne(scanner)
	} else {
		partTwo(scanner)
	}

}
