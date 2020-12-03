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
	resultPart1 := part1(input_file, 3, 1)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(input_file)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)
}

func part1(lines []string, right int, down int) (numTrees int) {
	currentPosition := 0
	width := len(lines[0])

	for i := down; i < len(lines); i += down {
		currentPosition = (currentPosition + right) % width
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

	totalTrees = 1
	for i, slope := range checkSlopes {
		slopeTrees := part1(lines, slope.Right, slope.Down)
		fmt.Printf("slope %d (%d r, %d d): %d\n", i, slope.Right, slope.Down, slopeTrees)
		totalTrees *= part1(lines, slope.Right, slope.Down)
	}

	return totalTrees
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
