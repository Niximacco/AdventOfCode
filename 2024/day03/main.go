package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	part1, part2 := 0, 0
	enabled := true
	for _, line := range input_file {
		r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don?'?t?\(\)`)
		matches := r.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			fmt.Printf("%v\n", match)
			if strings.Contains(match[0], "mul") {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				fmt.Printf("%v => %d * %d\n", match[0], num1, num2)
				part1 += num1 * num2
				if enabled {
					part2 += num1 * num2
				}
			} else {
				if strings.Contains(match[0], "don't") {
					enabled = false
				} else if strings.Contains(match[0], "do") {
					enabled = true
				}
			}
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
