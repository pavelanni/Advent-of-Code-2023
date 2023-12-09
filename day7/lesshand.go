package main

import (
	"strings"
)

func lessHand(hand1, hand2 string) bool {
	//fmt.Println(hand1, hand2)
	hand1 = strings.Replace(hand1, "T", "V", -1)
	hand1 = strings.Replace(hand1, "J", "1", -1) // changed for part 2
	hand1 = strings.Replace(hand1, "Q", "X", -1)
	hand1 = strings.Replace(hand1, "K", "Y", -1)
	hand1 = strings.Replace(hand1, "A", "Z", -1)
	hand2 = strings.Replace(hand2, "T", "V", -1)
	hand2 = strings.Replace(hand2, "J", "1", -1) // changed for part 2
	hand2 = strings.Replace(hand2, "Q", "X", -1)
	hand2 = strings.Replace(hand2, "K", "Y", -1)
	hand2 = strings.Replace(hand2, "A", "Z", -1)
	//fmt.Println(hand1, hand2)
	return hand1 < hand2
}
