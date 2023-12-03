package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type take struct {
	red   int
	green int
	blue  int
}

type game struct {
	id    int
	takes []take
}

var reGameTakes = `^Game (\d*): (.*$)`

func lineToGame(line string) game {
	g := game{}
	re := regexp.MustCompile(reGameTakes)
	matches := re.FindStringSubmatch(line)
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal(err)
	}
	g.id = id
	newTake := take{}
	g.takes = []take{}
	takesStr := regexp.MustCompile(";").Split(matches[2], -1)
	for _, t := range takesStr {
		newTake = take{}
		cubeColors := regexp.MustCompile(",").Split(t, -1)
		for _, cColor := range cubeColors {
			cColor = strings.TrimSpace(cColor)
			cNumber := regexp.MustCompile(" ").Split(cColor, -1)
			switch cNumber[1] {
			case "red":
				newTake.red, err = strconv.Atoi(cNumber[0])
				if err != nil {
					log.Fatal(err)
				}
			case "blue":
				newTake.blue, err = strconv.Atoi(cNumber[0])
				if err != nil {
					log.Fatal(err)
				}
			case "green":
				newTake.green, err = strconv.Atoi(cNumber[0])
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		g.takes = append(g.takes, newTake)
	}
	return g
}

func main() {

	var sum int
	var possible bool
	games := []game{}

	// open a file called input.txt and read line by line
	// for each line, create a game object
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		g := lineToGame(line)
		games = append(games, g)
	}
	for _, g := range games {
		possible = true
		for _, t := range g.takes {
			if t.red > 12 || t.green > 13 || t.blue > 14 {
				possible = false
				//fmt.Println("Impossible: ", g)
				break
			}
		}
		if possible {
			//fmt.Println("Possible: ", g)
			sum += g.id
		}
	}
	fmt.Println(sum)

	var powersum int
	for _, g := range games {
		var minred, mingreen, minblue int
		for _, t := range g.takes {
			if t.red > minred {
				minred = t.red
			}
			if t.green > mingreen {
				mingreen = t.green
			}
			if t.blue > minblue {
				minblue = t.blue
			}
		}
		fmt.Println("Game ID: ", g.id, "takes: ", g.takes)
		fmt.Println("Min red, green, blue: ", minred, mingreen, minblue)
		powersum += minred * mingreen * minblue
	}

	fmt.Println(powersum)
}
