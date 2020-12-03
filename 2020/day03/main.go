package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Slope struct {
	Right int
	Down  int
}

func main() {
	boolPtr := flag.Bool("test", false, "test mode")
	flag.Parse()

	var filename string
	if *boolPtr {
		filename = "input-test.txt"
	} else {
		filename = "input.txt"
	}

	input_file, err := readLines(filename)
	check(err)

	// Your Code goes below!
	part1Slope := Slope{3, 1}
	resultPart1 := part1(input_file, part1Slope)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(input_file)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)
}

func part1(lines []string, slope Slope) (numTrees int) {
	currentPosition := 0
	width := len(lines[0])

	for i := slope.Down; i < len(lines); i += slope.Down {
		currentPosition = (currentPosition + slope.Right) % width
		if string(lines[i][currentPosition]) == "#" {
			numTrees++
		}
	}

	return numTrees
}

func part2(lines []string) (totalTrees int) {
	checkSlopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	// Explicitly set it to 1 because the named return is 0 by default and multiplying by 0 is 0
	totalTrees = 1
	for _, slope := range checkSlopes {
		totalTrees *= part1(lines, slope)
	}

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
