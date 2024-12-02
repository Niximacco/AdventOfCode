package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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
	part1, part2 := 0, 0
	for _, line := range input_file {
		re := regexp.MustCompile("[0-9]+")
		strings := re.FindAllString(line, -1)
		var ints []int
		for _, stringint := range strings {
			val, _ := strconv.Atoi(stringint)
			ints = append(ints, val)
		}
		safe := isSafe(ints)
		if safe {
			part1++
			part2++
		} else {
			if recursiveSafe(ints) {
				part2++
			}
		}
	}

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func recursiveSafe(input []int) bool {
	fmt.Printf("testing\n%v\n", input)
	for i, _ := range input {
		newInput := remove(input, i)
		fmt.Printf("\t%d%v", i, newInput)
		if isSafe(newInput) {
			fmt.Print("\tSAFE\n")
			return true
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Still Unsafe\n")
	return false
}

func isSafe(inputs []int) bool {
	last := -1
	for _, num := range inputs {
		if last != -1 {
			diff := absDiffInt(last, num)
			if diff < 1 || diff > 3 {
				return false
			}
		}

		last = num
	}

	// is sorted?
	sortedAsc := sort.SliceIsSorted(inputs, func(p, q int) bool {
		return inputs[p] < inputs[q]
	})

	if sortedAsc {
		return true
	}

	sortedDesc := sort.SliceIsSorted(inputs, func(p, q int) bool {
		return inputs[p] > inputs[q]
	})

	return sortedDesc
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

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func remove(slice []int, s int) []int {
	var newSlice []int
	newSlice = append(newSlice, slice[:s]...)
	newSlice = append(newSlice, slice[s+1:]...)
	return newSlice
}
