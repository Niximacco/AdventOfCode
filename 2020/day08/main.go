package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	resultPart1, _ := part1(input_file)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(input_file)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(lines []string) (accumulator int, completed bool) {
	i := 0
	visited := make(map[int]bool)

	for {
		if _, present := visited[i]; present {
			return
		}

		if i >= len(lines) {
			completed = true
			return
		}
		visited[i] = true

		lineParts := strings.Fields(lines[i])
		switch lineParts[0] {
		case "nop":
			i++
		case "acc":
			amount, err := strconv.Atoi(lineParts[1])
			check(err)

			accumulator += amount
			i++
		case "jmp":
			amount, err := strconv.Atoi(lineParts[1])
			check(err)

			i += amount
		}
	}
}

func part2(lines []string) (accumulator int) {
	var jmpLines []int
	for i, line := range lines {
		lineParts := strings.Fields(line)
		if lineParts[0] == "jmp" {
			jmpLines = append(jmpLines, i)
		}
	}

	for _, jmpLineIndex := range jmpLines {
		trialLines := make([]string, len(lines))
		copy(trialLines, lines)
		trialLines[jmpLineIndex] = "nop 0"
		accumulator, completed := part1(trialLines)
		if completed {
			return accumulator
		}
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
