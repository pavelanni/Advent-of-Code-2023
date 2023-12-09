package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type seed struct {
	seed int
	maps map[string]int
}

var mapIndex = []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	mapStrings := regexp.MustCompile(`\n\n`).Split(string(data), -1)
	if err != nil {
		panic(err)
	}

	var seedStrings []string

	maps := make([][][]int, len(mapStrings)-1)
	for i, m := range mapStrings {
		loc := regexp.MustCompile(`:`).Split(m, -1)
		if err != nil {
			panic(err)
		}
		mapData := loc[1]
		if i == 0 { // seeds
			seedStrings = regexp.MustCompile(`\s+`).Split(strings.TrimSpace(mapData), -1)
			if err != nil {
				panic(err)
			}
		} else { // other maps
			lines := regexp.MustCompile(`\n`).Split(strings.TrimSpace(mapData), -1)
			if err != nil {
				panic(err)
			}
			mapSlice := make([][]int, 0)
			for _, l := range lines {
				m := make([]int, 3)
				fmt.Sscanf(l, "%d %d %d", &m[0], &m[1], &m[2])
				mapSlice = append(mapSlice, m)
			}
			maps[i-1] = mapSlice
		}
	}

	const maxUint = ^uint(0)
	const maxInt = int(maxUint >> 1) // from here https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
	// or use them from math: https://pkg.go.dev/math#pkg-constants
	minLocation := maxInt
	for i := 0; i < len(seedStrings); i += 2 {
		var sStart, sLength int
		fmt.Sscanf(seedStrings[i], "%d", &sStart)
		fmt.Sscanf(seedStrings[i+1], "%d", &sLength)
		fmt.Println(sStart, sLength)
		for j := 0; j < sLength; j++ {
			s := seed{
				seed: sStart + j,
				maps: make(map[string]int),
			}
			source := s.seed
			for i, m := range maps {
				//fmt.Println("Looking for ", mapIndex[i])
				found := false
				for _, l := range m {
					if source >= l[1] && source <= l[1]+l[2] {
						s.maps[mapIndex[i]] = l[0] + source - l[1] // found in maps
						found = true
						break
					}
				}
				// if not found in maps
				if !found {
					s.maps[mapIndex[i]] = source
				}
				source = s.maps[mapIndex[i]]
			}
			if s.maps["location"] < minLocation {
				minLocation = s.maps["location"]
			}

		}
	}

	fmt.Println("Min location: ", minLocation)
}
