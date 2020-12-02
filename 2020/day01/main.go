package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
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

	numbers := map[int]bool{}
	for _, line := range(input_file) {
		fmt.Printf("%s\n", line)
		number, _ := strconv.Atoi(line)
		numbers[number] = true
	}

	resultPart1 := part1(numbers, 2020)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(numbers, 2020)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(numbers map[int]bool, searchValue int) (result int) {
	// Search through the numbers for 2 that add up to 2020, then multiply them together and return the result.
	for key, _ := range numbers {
		if _, present := numbers[searchValue - key]; present {
			return key * (searchValue - key)
		}
	}

	return -1
}

func part2(numbers map[int]bool, searchValue int) (result int) {
	for key, _ := range numbers {
		if result := part1(numbers, searchValue - key); result > 0 {
			return key * result
		}
	}

	return -1
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
