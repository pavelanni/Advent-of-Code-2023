package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	cardNo  int
	matches int
}

func main() {
	var cards []card

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := regexp.MustCompile(`\r?\n`).Split(string(data), -1)
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		if line == "" {
			continue
		}
		loc := regexp.MustCompile(`[:|]`).Split(line, -1)
		winNums := regexp.MustCompile(`[ ]+`).Split(strings.TrimSpace(loc[1]), -1)

		winNumSet := make(map[string]struct{})
		for _, num := range winNums {
			winNumSet[num] = struct{}{}
		}
		myNums := regexp.MustCompile(`[ ]+`).Split(strings.TrimSpace(loc[2]), -1)
		matches := 0
		for _, num := range myNums {
			if _, ok := winNumSet[num]; ok {
				matches++
			}
		}
		n, err := strconv.Atoi(regexp.MustCompile(`[ ]+`).Split(strings.TrimSpace(loc[0]), -1)[1])
		if err != nil {
			panic(err)
		}
		cards = append(cards, card{
			cardNo:  n,
			matches: matches,
		})
	}

	sum := uint32(0)
	for _, card := range cards {
		//fmt.Println(card.name, card.matches, uint32(1)<<(card.matches-1), sum)
		if card.matches > 0 {
			sum += uint32(1) << (card.matches - 1)
		}
	}
	fmt.Println(sum)
	for i, c := range cards {
		if i < 10 {
			fmt.Println(c.cardNo, c.matches)
		}
	}

	scards := make([][]card, len(cards))
	// first pass: fill just one card in each slot
	for i := 0; i < len(cards); i++ {
		scards[i] = append(scards[i], cards[i])
	}
	for i, sc := range scards {
		for k := 0; k < len(sc); k++ { // for each scratchcards in this slot
			for j := 1; j <= sc[k].matches; j++ { // add cards to other slots
				scards[i+j] = append(scards[i+j], cards[i+j])
			}

		}
	}

	sumSC := 0
	for _, sc := range scards {
		sumSC += len(sc)
	}
	fmt.Println(sumSC)
}
