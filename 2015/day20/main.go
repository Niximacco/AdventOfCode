package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var deliveries map[int]int
var excludeFactors map[int]bool

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
	part1, part2 := 0, 0
	var goal int
	for _, line := range input_file {
		goal, _ = strconv.Atoi(line)
	}
	fmt.Printf("\n")

	var result int
	part1 = 831000
	deliveries = make(map[int]int)
	excludeFactors = make(map[int]bool)

	for result <= goal {
		part1 += 2
		factors := getFactorsForNumber(part1)
		result = 0
		for _, factor := range factors {
			result += factor * 10
		}
		if result > 25000000 {
			fmt.Printf("%d: %d\n", part1, result)
		}
	}

	result = 0
	part2 = 0
	for result <= goal {
		part2 += 1
		factors := getFactorsForNumber(part2)
		result = 0
		// exclude exhausted factors
		// decrement deliveries remaining

		for _, factor := range factors {
			if !excludeFactors[factor] {
				result += factor * 11
				deliveries[factor]++

				if deliveries[factor] == 50 {
					excludeFactors[factor] = true
				}
			}
		}
		if result > 25000000 {
			fmt.Printf("%d: %d\n", part2, result)
		}
	}

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
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

func getFactorsForNumber(number int) (factors []int) {
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}
