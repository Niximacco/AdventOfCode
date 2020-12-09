package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Num struct {
	Numbers    []int
	CheckIndex int
	Lookahead  int
}

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

	numbers := parseLines(input_file)
	// Your Code goes below!
	var resultPart1 int
	if *boolPtr {
		resultPart1 = part1(numbers, 5)
	} else {
		resultPart1 = part1(numbers, 25)
	}
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(numbers, resultPart1)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(numbers []int, lookahead int) (invalid int) {
	for i, number := range numbers {
		numCheck := Num{
			Numbers:    numbers,
			CheckIndex: i,
			Lookahead:  lookahead,
		}
		if !numCheck.isValid() {
			invalid = number
			return
		}
	}
	return
}

func (numCheck *Num) isValid() (valid bool) {
	if numCheck.CheckIndex < numCheck.Lookahead {
		return true
	}

	possibleSums := make(map[int]bool)

	low := numCheck.CheckIndex - numCheck.Lookahead
	high := numCheck.CheckIndex
	fmt.Printf("low: %d\thigh: %d\n", low, high)
	lookaheadNums := numCheck.Numbers[low:high]

	for i1, num1 := range lookaheadNums {
		for i2, num2 := range lookaheadNums {
			if i1 == i2 {
				continue
			}
			possibleSums[num1+num2] = true
		}
	}

	if _, present := possibleSums[numCheck.Numbers[numCheck.CheckIndex]]; present {
		valid = true
	}

	return
}

func part2(numbers []int, searchNum int) (weakness int) {
	for startIndex := 0; startIndex < len(numbers); startIndex++ {
		checkSum := 0
		for checkIndex := startIndex; checkSum < searchNum; checkIndex++ {
			checkSum += numbers[checkIndex]
			if checkSum == searchNum {
				lowest := numbers[startIndex]
				highest := numbers[startIndex]

				for _, number := range numbers[startIndex:checkIndex] {
					if number < lowest {
						lowest = number
					}

					if number > highest {
						highest = number
					}
				}

				return lowest + highest
			}
		}
	}
	return
}

func parseLines(lines []string) (numbers []int) {
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		check(err)

		numbers = append(numbers, number)
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
