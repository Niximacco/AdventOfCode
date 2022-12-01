package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	var totals []int
	current := 0
	for _, line := range input_file {
		if line != "" {
			value, err := strconv.Atoi(line)
			check(err)

			current += value
		} else {
			totals = append(totals, current)
			current = 0
		}
	}
	totals = append(totals, current)

	sort.Ints(totals)

	part1 := totals[len(totals)-1]
	part2 := 0
	for _, v := range totals[len(totals)-3:] {
		part2 += v
	}

	fmt.Printf("%v\n", totals)

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
