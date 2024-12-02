package main

import (
	"flag"
	"fmt"
	utils "github.com/niximacco/aocutils"
	"sort"
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
	input_file, err := utils.ReadLines(filename)
	utils.Check(err)

	// Your Code goes below!
	part1, part2 := 0, 0
	for _, line := range input_file {
		ints := utils.GetAllIntegersFromString(line)
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
			diff := utils.AbsDiffInt(last, num)
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

func remove(slice []int, s int) []int {
	var newSlice []int
	newSlice = append(newSlice, slice[:s]...)
	newSlice = append(newSlice, slice[s+1:]...)
	return newSlice
}
