package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var values [][]int
	var sum, sum2 int

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
	for _, line := range lines {
		vSlice := strings.Split(line, " ")
		var vRow []int
		for _, v := range vSlice {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			vRow = append(vRow, vInt)
		}
		values = append(values, vRow)
	}

	for _, v := range values {
		var valToZeros [][]int
		valToZeros = append(valToZeros, v)
		allZeros := false
		for !allZeros {
			var row []int
			for i := 0; i < len(v)-1; i++ {
				row = append(row, v[i+1]-v[i])
			}
			allZeros = true
			for _, n := range row {
				if n != 0 {
					allZeros = false
				}
			}
			v = row
			valToZeros = append(valToZeros, row)
		}
		// filled the slice of slices down to all zeros
		// now let go back
		for i := len(valToZeros) - 1; i > 0; i-- {
			valToZeros[i-1] = append(valToZeros[i-1], valToZeros[i-1][len(valToZeros[i-1])-1]+valToZeros[i][len(valToZeros[i])-1])
		}
		sum += valToZeros[0][len(valToZeros[0])-1]
	}
	fmt.Println("sum: ", sum)

	fmt.Println("Part 2")
	for _, v := range values {
		var valToZeros [][]int
		valToZeros = append(valToZeros, v)
		allZeros := false
		for !allZeros {
			var row []int
			for i := 0; i < len(v)-1; i++ {
				row = append(row, v[i+1]-v[i])
			}
			allZeros = true
			for _, n := range row {
				if n != 0 {
					allZeros = false
				}
			}
			v = row
			valToZeros = append(valToZeros, row)
		}
		// filled the slice of slices down to all zeros
		// now let go back for Part 2 and add a number in the beginning
		for i := len(valToZeros) - 1; i > 0; i-- {
			valToZeros[i-1] = append([]int{valToZeros[i-1][0] - valToZeros[i][0]}, valToZeros[i-1]...)
		}
		sum2 += valToZeros[0][0]
	}

	fmt.Println("sum2: ", sum2)
}
