package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

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
	values := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3, "W": 6, "D": 3, "L": 0}
	valueSlice := []int{1, 2, 3, 1, 2, 3}
	part1 := 0
	part2 := 0
	result := ""
	score := 0
	for _, line := range input_file {
		combo := strings.Split(line, " ")
		them := combo[0]

		// Part 1  Logic
		us := combo[1]
		fmt.Printf("%v\n", combo)
		result = ""

		if values[us] == values[them] {
			result = "D"
		} else if (values[us]-values[them]+3)%3 == 1 {
			result = "W"
		} else {
			result = "L"
		}
		score = values[result] + values[us]
		part1 += score
		fmt.Printf("pt 1: %d\n", score)

		// Part 2 Logic
		score = 0
		desiredResult := combo[1]
		if desiredResult == "Y" {
			// Draw
			score += values["D"] + values[them]
		} else if desiredResult == "X" {
			// Lose
			score += valueSlice[(values[them] + 1)]
		} else {
			// Win
			score += valueSlice[(values[them])%3]
			score += values["W"]
		}
		fmt.Printf("pt 2: %d\n", score)
		part2 += score
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
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
