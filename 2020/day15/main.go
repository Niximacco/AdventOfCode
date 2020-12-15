package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
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

	resultPart1 := part1(input_file, 2020)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)
	//
	//resultPart2 := part2(instructions)
	//fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(lines []string, stop int) (value int) {
	lookup := make(map[int]int)
	timesSeen := make(map[int]int)

	turn := 0
	lastNum := 0

	initialNums := strings.Split(lines[0], ",")

	for _, numString := range initialNums {
		num, _ := strconv.Atoi(numString)

		lookup[num] = turn
		lastNum = num
		timesSeen[num]++
		turn++
	}

	for turn < stop {
		if timesSeen[lastNum] == 1 {
			lastNum = 0
			lookup[0] = turn
			timesSeen[0]++
		} else {

		}

		turn++
	}
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
