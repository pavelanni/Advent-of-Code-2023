package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type part struct {
	number int
	line   int
	start  int
	end    int
	isPart bool
}

type gear struct {
	line     int
	pos      int
	adjParts []int
	isGear   bool
}

var reDigits = regexp.MustCompile(`\d+`)
var reNotDot = regexp.MustCompile(`[^.0-9]`)
var reStar = regexp.MustCompile(`[*]`)

func main() {
	var parts []part

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	// remove the last line because it's empty
	lines = lines[:len(lines)-1]

	// collect all parts
	for i, line := range lines {
		loc := reDigits.FindAllStringIndex(line, -1)
		if loc != nil {
			for j := 0; j < len(loc); j++ {
				n, err := strconv.Atoi(line[loc[j][0]:loc[j][1]])
				if err != nil {
					log.Fatal(err)
				}
				parts = append(parts, part{
					number: n,
					line:   i,
					start:  loc[j][0],
					end:    loc[j][1],
				})
			}
		}
	}

	// find adjacent bytes to each part
	for i, p := range parts {
		var adj []byte
		if p.line > 0 { // add from the line above
			adj = append(adj, lines[p.line-1][p.start:p.end]...) // right above
			if p.start > 0 {
				adj = append(adj, lines[p.line-1][p.start-1:p.start]...) // left
			}
			if p.end < len(lines[p.line-1]) {
				adj = append(adj, lines[p.line-1][p.end:p.end+1]...) // right
			}
		}
		if p.line < len(lines)-1 { //add from the line below
			adj = append(adj, lines[p.line+1][p.start:p.end]...)
			if p.start > 0 {
				adj = append(adj, lines[p.line+1][p.start-1:p.start]...)
			}
			if p.end < len(lines[p.line+1]) {
				adj = append(adj, lines[p.line+1][p.end:p.end+1]...)
			}
		}
		// add from the same line
		if p.start > 0 {
			adj = append(adj, lines[p.line][p.start-1:p.start]...)
		}
		if p.end < len(lines[p.line]) {
			adj = append(adj, lines[p.line][p.end:p.end+1]...)
		}
		// check if part
		if len(reNotDot.FindAll(adj, -1)) > 0 {
			p.isPart = true
		}
		//fmt.Println(p.number, " ", string(adj), " ", p.isPart)
		parts[i] = p
	}

	var sum int
	for _, p := range parts {
		if p.isPart {
			sum += p.number
		}
	}
	fmt.Println(sum)

	var gears []gear
	// collect all gears
	for i, line := range lines {
		loc := reStar.FindAllStringIndex(line, -1)
		if loc != nil {
			for j := 0; j < len(loc); j++ {
				gears = append(gears, gear{
					line: i,
					pos:  loc[j][0],
				})
			}
		}
	}
	for g := range gears {
		fmt.Println(gears[g].line, gears[g].pos)
		var adjParts []int
		// at the same line
		loc := reDigits.FindAllStringIndex(lines[gears[g].line], -1)
		if loc != nil {
			for j := 0; j < len(loc); j++ {
				if loc[j][0] == gears[g].pos+1 || loc[j][1] == gears[g].pos { // beginning or end of the part is adjacent
					partNo, err := strconv.Atoi(lines[gears[g].line][loc[j][0]:loc[j][1]])
					if err != nil {
						log.Fatal(err)
					}
					adjParts = append(adjParts, partNo)
				}
			}
		}
		// at the line above
		if gears[g].line > 0 {
			loc = reDigits.FindAllStringIndex(lines[gears[g].line-1], -1)
			if loc != nil {
				for j := 0; j < len(loc); j++ {
					if gears[g].pos >= loc[j][0]-1 && gears[g].pos <= loc[j][1] {
						partNo, err := strconv.Atoi(lines[gears[g].line-1][loc[j][0]:loc[j][1]])
						if err != nil {
							log.Fatal(err)
						}
						adjParts = append(adjParts, partNo)
					}
				}
			}
		}
		// at the line below
		if gears[g].line < len(lines)-1 {
			loc = reDigits.FindAllStringIndex(lines[gears[g].line+1], -1)
			if loc != nil {
				for j := 0; j < len(loc); j++ {
					if gears[g].pos >= loc[j][0]-1 && gears[g].pos <= loc[j][1] {
						partNo, err := strconv.Atoi(lines[gears[g].line+1][loc[j][0]:loc[j][1]])
						if err != nil {
							log.Fatal(err)
						}
						adjParts = append(adjParts, partNo)
					}
				}
			}
		}
		gears[g].adjParts = adjParts
		if len(adjParts) == 2 {
			gears[g].isGear = true
		}
		fmt.Println(gears[g].line, gears[g].pos, gears[g].adjParts, gears[g].isGear)
	}

	var sumRatios int
	for _, g := range gears {
		if g.isGear {
			sumRatios += g.adjParts[1] * g.adjParts[0]
		}
	}
	fmt.Println(sumRatios)

}
