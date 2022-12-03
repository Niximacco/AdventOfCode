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
	part1 := 0
	for _, line := range input_file {
		comp1 := line[:(len(line) / 2)]
		comp2 := line[len(line)/2:]
		fmt.Printf("%s -> (%s) (%s)\n", line, comp1, comp2)
		for _, char := range []rune(comp1) {
			if strings.Contains(comp2, string(char)) {
				part1 += getPriority(char)
				break
			}
		}
	}

	part2 := 0
	for i := 0; i < len(input_file); i += 3 {
		var group []string
		group = input_file[i : i+3]
		for _, char := range []rune(group[0]) {
			if strings.Contains(group[1], string(char)) && strings.Contains(group[2], string(char)) {
				part2 += getPriority(char)
				break
			}
		}
	}

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)

}

func getPriority(char rune) (priority int) {

	if char >= 65 && char <= 90 {
		priority = int(char) - 38
	} else {
		priority = int(char) - 96
	}

	fmt.Printf("[%d] %s -> %d\n", char, string(char), priority)
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
