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

type Stack map[int][]string

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
	part1, part2 := "", ""
	stackPart1 := getInitialStacks(input_file)
	stackPart2 := getInitialStacks(input_file)
	for _, line := range input_file {
		if !strings.Contains(line, "move") {
			continue
		}

		re := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
		parts := re.FindStringSubmatch(line)
		count, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[2])
		to, _ := strconv.Atoi(parts[3])

		stackPart1.MovePart1(count, from, to)
		stackPart2.MovePart2(count, from, to)
	}

	part1 = stackPart1.GetOutput()
	part2 = stackPart2.GetOutput()

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func (stack Stack) GetOutput() (output string) {
	for i := 1; i <= len(stack); i++ {
		output = fmt.Sprintf("%s%s", output, stack[i][len(stack[i])-1])
	}

	return
}

func (stack Stack) MovePart1(count int, from int, to int) {
	onCrane := ""
	for i := 0; i < count; i++ {
		onCrane, stack[from] = stack[from][len(stack[from])-1], stack[from][:len(stack[from])-1]
		stack[to] = append(stack[to], onCrane)
	}
}

func (stack Stack) MovePart2(count int, from int, to int) {
	var onCrane []string
	onCrane, stack[from] = stack[from][len(stack[from])-count:], stack[from][:len(stack[from])-count]
	stack[to] = append(stack[to], onCrane...)
}

func getInitialStacks(input []string) (stacks Stack) {
	stacks = make(map[int][]string)
	for i, line := range input {
		if !strings.Contains(line, "[") {
			createStacks := strings.Fields(line)
			for col, stack := range createStacks {
				stackNum, _ := strconv.Atoi(stack)
				stacks[stackNum] = make([]string, 0)
				for row := i; row >= 0; row-- {
					rowLine := input[row]
					if (col*4 < len(rowLine)) && string(rowLine[col*4]) == "[" {
						stacks[stackNum] = append(stacks[stackNum], string(rowLine[(col*4)+1]))
					}
				}
			}
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
