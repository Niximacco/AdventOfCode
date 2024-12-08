package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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
	part1, part2 := 0, 0
	fmt.Printf("%d\n", getCombineVal(15, 6))
	fmt.Printf("%d\n", getCombineVal(1, 4))
	fmt.Printf("%d\n", getCombineVal(1433, 2))
	fmt.Printf("%d\n", getCombineVal(122222222, 6))

	for _, line := range input_file {
		println()
		fmt.Printf("%s\n", line)
		parts := strings.Split(line, ":")
		testNumber, _ := strconv.Atoi(parts[0])
		numberStrings := strings.Split(parts[1], " ")
		var numbers []int
		for _, numberStr := range numberStrings {
			number, _ := strconv.Atoi(numberStr)
			numbers = append(numbers, number)
		}

		currentValues := make(map[int]bool)
		for i, number := range numbers {
			if i == 0 {
				currentValues[number] = true
				continue
			}

			newValues := make(map[int]bool)
			for value, _ := range currentValues {
				newValues[value+number] = true
				newValues[value*number] = true
			}

			currentValues = newValues

			if newValues[testNumber] {
				part1 += testNumber
				break
			}
		}
	}

	for _, line := range input_file {
		println()
		fmt.Printf("%s\n", line)
		parts := strings.Split(line, ":")
		testNumber, _ := strconv.Atoi(parts[0])
		numbersString := strings.TrimSpace(parts[1])
		numberStrings := strings.Split(numbersString, " ")
		var numbers []int
		for _, numberStr := range numberStrings {
			fmt.Printf("'%s'\t", numberStr)
			number, _ := strconv.Atoi(numberStr)
			numbers = append(numbers, number)
		}
		println()

		currentPart2 := make(map[int]bool)
		//fmt.Printf("%v\n", currentPart2)
		//fmt.Printf("numbers: %v\n", numbers)
		for i, number := range numbers {
			if i == 0 {
				//fmt.Printf("[%d] setting numbers[%d] = true\n", i, number)
				currentPart2[number] = true
				continue
			}

			newValues := make(map[int]bool)
			//fmt.Printf("Starting(%v)\n", newValues)
			//fmt.Printf("currentPart2: %v\n", currentPart2)
			for value, _ := range currentPart2 {
				//fmt.Printf("[%d] Testing %d %d\n", i, value, number)
				newValues[value+number] = true
				newValues[value*number] = true

				newValues[getCombineVal(value, number)] = true
			}

			currentPart2 = newValues

			//fmt.Printf("%d: %v\n", testNumber, newValues)
		}
		if currentPart2[testNumber] {
			part2 += testNumber
		}
	}

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
	fmt.Printf("Max Int: %d\n", math.MaxInt)
}

func getCombineVal(a, b int) (result int) {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	// combine the strings
	combine := fmt.Sprintf("%s%s", aStr, bStr)

	// turn it back to an int
	result, _ = strconv.Atoi(combine)
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
