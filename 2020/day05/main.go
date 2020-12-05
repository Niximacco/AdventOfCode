package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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

	resultPart1 := part1(input_file)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(input_file)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)
}

func part1(lines []string) (highestSeatId int) {
	for _, line := range lines {
		row := calcRow(line[:7])
		seat := calcSeatId(line[7:])
		value := (row * 8) + seat
		fmt.Printf("%s: %d (seat %d, row %d)\n", line, value, seat, row)
		if value > highestSeatId {
			highestSeatId = value
		}
	}
	return
}

func part2(lines []string) (seatId int) {
	var seatIds []int
	for _, line := range lines {
		row := calcRow(line[:7])
		seat := calcSeatId(line[7:])
		value := (row * 8) + seat
		seatIds = append(seatIds, value)
	}

	sort.Ints(seatIds)
	previous := 0
	for _, seatId := range seatIds {
		if seatId != previous+1 && previous != 0 {
			return previous + 1
		}
		previous = seatId
	}

	return -1
}

func calcRow(rowInfo string) (rowNum int) {
	charCount := len(rowInfo)
	binaryValue := RecursivePower(2, (charCount - 1))
	for _, char := range rowInfo {
		if strings.ToUpper(string(char)) == "B" {
			rowNum += binaryValue
		}
		binaryValue /= 2
	}
	return
}

func RecursivePower(base int, exponent int) int {
	if exponent != 0 {
		return base * RecursivePower(base, exponent-1)
	} else {
		return 1
	}
}

func calcSeatId(seatInfo string) (seatNum int) {
	charCount := len(seatInfo)
	binaryValue := RecursivePower(2, (charCount - 1))
	for _, char := range seatInfo {
		if strings.ToUpper(string(char)) == "R" {
			seatNum += binaryValue
		}
		binaryValue /= 2
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
