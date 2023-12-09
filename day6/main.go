package main

import (
	"fmt"
	"os"
	"strings"
)

func numWonRaces(time, distance int) int {
	if time == 0 || distance == 0 {
		return 0
	}

	var nRaces int
	for tPress := 0; tPress <= time; tPress++ {
		speed := tPress
		myDistance := speed * (time - tPress)
		if myDistance > distance {
			nRaces++
		}
	}
	return nRaces
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	times := make([]int, 4)
	distances := make([]int, 4)
	titles := make([]string, 2)

	fmt.Sscanf(lines[0], "%s %d %d %d %d", &titles[0], &times[0], &times[1], &times[2], &times[3])
	fmt.Sscanf(lines[1], "%s %d %d %d %d", &titles[1], &distances[0], &distances[1], &distances[2], &distances[3])

	fmt.Println(times)
	fmt.Println(distances)

	mult := 1
	for i := 0; i < len(times); i++ {
		fmt.Printf("Set %d won %d races\n", i, numWonRaces(times[i], distances[i]))
		mult = mult * numWonRaces(times[i], distances[i])
	}
	fmt.Println("Mult: ", mult)

	fmt.Println("Now for part 2")

	newTime := 0
	newDistance := 0
	fmt.Sscanf(fmt.Sprintf("%d%d%d%d", times[0], times[1], times[2], times[3]), "%d", &newTime)
	fmt.Sscanf(fmt.Sprintf("%d%d%d%d", distances[0], distances[1], distances[2], distances[3]), "%d", &newDistance)

	fmt.Printf("%d %T\n", newTime, newTime)
	fmt.Printf("%d %T\n", newDistance, newDistance)
	n := numWonRaces(newTime, newDistance)
	fmt.Println("Part 2: ", n)
}
