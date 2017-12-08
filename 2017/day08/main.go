package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
	"strings"
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
	registers := make(map[string]int)
	maxPart2 := 0
	for _, line := range input_file {
		fields := strings.Fields(line)
		name := fields[0]
		operation := fields[1]
		amount, err := strconv.Atoi(fields[2])
		check(err)

		compLeft := fields[4]
		compSymbol := fields[5]
		compRight, err := strconv.Atoi(fields[6])
		check(err)

		//fmt.Printf("%sing %s by %d if %s %s %d\n", operation, name, amount, compLeft, compSymbol, compRight)
		if _, ok := registers[name]; !ok {
			registers[name] = 0
		}

		if _, ok := registers[compLeft]; !ok {
			registers[compLeft] = 0
		}

		eval := false
		switch compSymbol {
		case ">":
			eval = (registers[compLeft] > compRight)
		case ">=":
			eval = (registers[compLeft] >= compRight)
		case "<":
			eval = (registers[compLeft] < compRight)
		case "<=":
			eval = (registers[compLeft] <= compRight)
		case "==":
			eval = (registers[compLeft] == compRight)
		case "!=":
			eval = (registers[compLeft] != compRight)
		}

		if eval {
			//fmt.Printf("\tTRUE\n")
			switch operation {
			case "inc":
				registers[name] += amount
			case "dec":
				registers[name] -= amount
			}
		}
		if registers[name] > maxPart2 {
			maxPart2 = registers[name]
		}
		//
		//fmt.Println(registers)
		//fmt.Println("======")
	}

	max := 0
	maxKey := ""
	for k, val := range registers {
		if val > max {
			maxKey = k
			max = val
		}
	}

	fmt.Printf("Part 1 Answer: %d in %s\n", max, maxKey)
	fmt.Printf("Part 2 Answer: %d\n", maxPart2)
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
