package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"golang.org/x/text/message"
)

type node struct {
	left  string
	right string
}

func main() {

	p := message.NewPrinter(message.MatchLanguage("en"))
	nodes := map[string]node{}

	// Read input.txt line by line the recommended way
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println("Read in", len(lines), "lines")
	turns := lines[0] // turn insturctions are in the first line

	for _, line := range lines[2:] {
		re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			log.Fatal("bad line: ", line)
		}
		title := matches[1]
		left := matches[2]
		right := matches[3]
		nodes[title] = node{left, right}
	}

	currentNode := "AAA"
	var steps int
	for i := 0; ; i++ {
		if i >= len(turns) {
			i = 0
		}
		// fmt.Printf("Current node: %s, node: %s ", currentNode, nodes[currentNode])
		if turns[i] == 'L' {
			currentNode = nodes[currentNode].left
		} else if turns[i] == 'R' {
			currentNode = nodes[currentNode].right
		}
		steps++
		if currentNode == "ZZZ" {
			break
		}
	}
	fmt.Println("steps: ", steps)

	fmt.Println("Part 2")
	var currentNodes []string
	// how many keys end with "A"?
	for k, _ := range nodes {
		if k[2] == 'A' {
			currentNodes = append(currentNodes, k)
		}
	}
	var steps2 int
	for i := 0; ; i++ {
		if i >= len(turns) {
			i = 0
		}
		//fmt.Printf("Current nodes: %v ", currentNodes)
		for i, cn := range currentNodes {
			if turns[i] == 'L' {
				currentNodes[i] = nodes[cn].left
			} else if turns[i] == 'R' {
				currentNodes[i] = nodes[cn].right
			}
		}
		//fmt.Printf("Next nodes: %v\n", currentNodes)
		steps2++
		if steps2%10000000 == 0 {
			p.Println("Steps so far: ", steps2)
		}
		allZ := true
		for _, cn := range currentNodes {
			if cn[2] != byte('Z') {
				allZ = false
			}
		}
		if allZ {
			break
		}
	}
	fmt.Println("steps: ", steps2)
}
