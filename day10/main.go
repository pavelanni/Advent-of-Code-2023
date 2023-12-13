package main

import (
	"bufio"
	"fmt"
	"os"
)

var maze [][]byte

type cell struct {
	char byte
	x    int
	y    int
}

func main() {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		maze = append(maze, []byte(line))
	}

	var sx, sy int
	// find the starting position
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == 'S' {
				sx = j
				sy = i
			}
		}
	}
	sChar := cell{char: 'S', x: sx, y: sy}
	s, err := whatIsS(sChar)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s.char), s.x, s.y)

	var route []cell
	var steps int
	prevCell := cell{char: maze[sy-1][sx]} // imagine we came from the top
	curCell := s

	for {
		nextCell, err := curCell.nextStep(prevCell)
		//fmt.Println(string(nextCell.char), nextCell.x, nextCell.y)
		if err != nil {
			panic(err)
		}
		if nextCell.char == 'S' {
			curCell = nextCell // we came back to S
			fmt.Println(steps)
			break
		}
		route = append(route, curCell)
		prevCell = curCell
		curCell = nextCell
		steps++
	}

	steps = 0
	prevCell = cell{char: maze[sy][sx-1]} // imagine we came from the left
	curCell.char = 'J'                    // yes, it's a hack, I'll fix it later
	for {
		nextCell, err := curCell.nextStep(prevCell)
		//fmt.Println(string(nextCell.char), nextCell.x, nextCell.y)
		if err != nil {
			fmt.Println(string(curCell.char), curCell.x, curCell.y)
			panic(err)
		}
		if nextCell.char == 'S' {
			curCell = nextCell // we came back to S
			fmt.Println(steps)
			break
		}
		prevCell = curCell
		curCell = nextCell
		steps++
	}
	fmt.Println(steps/2 + steps%2)
	fmt.Println("Route length: ", len(route))

	// Print out the route
	for i := range maze {
		for j := range maze[i] {
			maze[i][j] = '.'
		}
	}
	for i := range route {
		maze[route[i].y][route[i].x] = route[i].char
	}
	for i := range maze {
		fmt.Println(string(maze[i]))
	}
}
