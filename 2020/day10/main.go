package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
	adapters := parseAdapters(input_file)

	resultPart1 := part1(adapters)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(adapters)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(adapters map[int]bool) (result int) {
	currentJoltage := 0
	usedAdapters := 0

	diffs := map[int]int{
		1: 0,
		2: 0,
		3: 1,
	}

	for usedAdapters < len(adapters) {
		if used, present := adapters[currentJoltage+1]; present && !used {
			currentJoltage += 1
			usedAdapters++
			diffs[1]++
		} else if used, present := adapters[currentJoltage+2]; present && !used {
			currentJoltage += 2
			usedAdapters++
			diffs[2]++
		} else if used, present := adapters[currentJoltage+3]; present && !used {

			currentJoltage += 3
			usedAdapters++
			diffs[3]++
		}
	}

	result = diffs[1] * diffs[3]

	return
}

func part2(adapters map[int]bool) (result int) {
	maxJoltage := 0
	ways := make(map[int]int)
	var adapterNums []int

	for k := range adapters {
		adapterNums = append(adapterNums, k)
		if k > maxJoltage {
			maxJoltage = k
		}
	}

	ways[0] = 1
	sort.Ints(adapterNums)
	for _, num := range adapterNums {
		ways[num] = ways[num-1] + ways[num-2] + ways[num-3]
	}

	result = ways[adapterNums[len(adapterNums)-1]]

	return
}

func countUniqueCombos(adapters map[int]bool, currentJoltage int, maxJoltage int) (numCombos int) {
	childAdapters := make(map[int]bool)
	for k := range adapters {
		if k != currentJoltage && k >= currentJoltage {
			childAdapters[k] = false
		}
	}

	for i := 1; i <= 3; i++ {
		if currentJoltage+i == maxJoltage {
			numCombos++
		}

		if used, present := adapters[currentJoltage+i]; present && !used {
			numCombos += countUniqueCombos(childAdapters, currentJoltage+i, maxJoltage)
		}
	}

	fmt.Printf("%d: %d\n", len(adapters), numCombos)
	return
}

func parseAdapters(input []string) (adapters map[int]bool) {
	adapters = make(map[int]bool)
	for _, line := range input {
		number, err := strconv.Atoi(line)
		check(err)

		adapters[number] = false
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
