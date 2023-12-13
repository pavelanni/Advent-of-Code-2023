package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var space [][]byte

type galaxy struct {
	x, y int
}

var galaxies []galaxy

func intAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// transpose transposes a 2D slice of bytes
func transpose(a [][]byte) [][]byte {
	b := make([][]byte, len(a[0]))
	for i := 0; i < len(a[0]); i++ {
		b[i] = make([]byte, len(a))
		for j := 0; j < len(a); j++ {
			b[i][j] = a[j][i]
		}
	}
	return b
}

func expandRows(space [][]byte) [][]byte {
	var newSpace [][]byte
	for _, r := range space {
		newSpace = append(newSpace, r)
		if regexp.MustCompile(`^\.+$`).Match(r) {
			newSpace = append(newSpace, r) // if no galaxies in this row
		}
	}
	return newSpace
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
		space = append(space, []byte(line))
	}

	fmt.Printf("Space size: %d x %d\n", len(space[0]), len(space))
	// expand rows
	space = expandRows(space)
	fmt.Printf("Space size: %d x %d\n", len(space[0]), len(space))
	// transpose
	space = transpose(space)
	fmt.Printf("Space size: %d x %d\n", len(space[0]), len(space))
	space = expandRows(space)
	fmt.Printf("Space size: %d x %d\n", len(space[0]), len(space))
	// transpose again
	space = transpose(space)
	fmt.Printf("Space size: %d x %d\n", len(space[0]), len(space))

	// fill the galaxies slice
	for y := 0; y < len(space); y++ {
		for x := 0; x < len(space[0]); x++ {
			if space[y][x] == '#' {
				galaxies = append(galaxies, galaxy{x, y})
			}
		}
	}

	fmt.Printf("Galaxies: %d\n", len(galaxies))
	var sum int

	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			g1 := galaxies[i]
			g2 := galaxies[j]
			distance := intAbs(g1.x-g2.x) + intAbs(g1.y-g2.y)
			sum += distance
		}
	}
	fmt.Println(sum)
}
