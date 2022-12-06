package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	line := input_file[0]
	check(err)

	// Your Code goes below!
	part1 := FindUniqueCharsInARowIndex(line, 4)
	part2 := FindUniqueCharsInARowIndex(line, 14)

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func FindUniqueCharsInARowIndex(input string, uniqueChars int) (numChars int) {
	currentChars := make(map[string]int)

	for i := 0; i < len(input); i++ {
		char := string(input[i])

		if _, ok := currentChars[char]; ok {
			currentChars[char]++
		} else {
			currentChars[char] = 1
		}

		if i-uniqueChars >= 0 {
			oldestChar := string(input[i-uniqueChars])
			currentChars[oldestChar]--
			if currentChars[oldestChar] == 0 {
				delete(currentChars, oldestChar)
			}
		}

		if len(currentChars) == uniqueChars {
			numChars = i + 1
			break
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
