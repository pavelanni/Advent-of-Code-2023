package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type hand struct {
	handCards string
	handType  int
	handBid   int
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func getType(cards string) int {
	type labelpair struct {
		label string
		count int
	}

	labels := map[string]int{
		"A": 0,
		"K": 0,
		"Q": 0,
		"J": 0,
		"9": 0,
		"8": 0,
		"7": 0,
		"6": 0,
		"5": 0,
		"4": 0,
		"3": 0,
		"2": 0,
	}

	for _, c := range cards {
		labels[string(c)]++
	}

	labelpairs := []labelpair{}
	for k, v := range labels {
		labelpairs = append(labelpairs, labelpair{k, v})
	}
	sort.Slice(labelpairs, func(i, j int) bool {
		return labelpairs[i].count > labelpairs[j].count
	})
	// added for part 2
	var k int
	if labels["J"] > 0 {
		if labelpairs[0].label == "J" {
			k = 1
		} else {
			k = 0
		}
		labelpairs[k].count += labels["J"]
		if labelpairs[k].count > 5 {
			labelpairs[k].count = 5
		}
		for i := range labelpairs {
			if labelpairs[i].label == "J" {
				labelpairs[i].count = 0
				break
			}
		}
	}
	for i, lp := range labelpairs {
		if lp.count == 5 {
			return FiveOfAKind
		}
		if lp.count == 4 {
			return FourOfAKind
		}
		if lp.count == 3 {
			if labelpairs[i+1].count == 2 {
				return FullHouse
			} else {
				return ThreeOfAKind
			}
		}
		if lp.count == 2 {
			if labelpairs[i+1].count == 2 {
				return TwoPair
			} else {
				return OnePair
			}
		}
	}
	return HighCard
}

func main() {
	var hands []hand
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		hand := hand{}
		fmt.Sscanf(line, "%s %d", &hand.handCards, &hand.handBid)
		hands = append(hands, hand)
	}
	for i, hand := range hands {
		hand.handType = getType(hand.handCards)
		hands[i] = hand
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		} else {
			return lessHand(hands[i].handCards, hands[j].handCards)
		}
	})
	for _, h := range hands {
		fmt.Println(h)
	}
	var total int
	for i, h := range hands {
		total += h.handBid * (i + 1)
	}
	fmt.Println(total)
}
