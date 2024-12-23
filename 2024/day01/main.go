package main

import (
	"flag"
	"fmt"
	utils "github.com/niximacco/aocutils"
	"sort"
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
	input_file, err := utils.ReadLines(filename)
	utils.Check(err)

	// Your Code goes below!
	part1, part2 := 0, 0
	var line_parts_left []int
	var line_parts_right []int

	for _, line := range input_file {
		line_parts := strings.Split(line, "   ")
		line_left, _ := strconv.Atoi(line_parts[0])
		line_right, _ := strconv.Atoi(line_parts[1])

		line_parts_left = append(line_parts_left, line_left)
		line_parts_right = append(line_parts_right, line_right)
	}

	sort.Ints(line_parts_left)
	sort.Ints(line_parts_right)

	// Make frequency lookup
	frequency := make(map[int]int)

	for i := 0; i < len(line_parts_left); i++ {
		diff := utils.AbsDiffInt(line_parts_left[i], line_parts_right[i])
		part1 += diff
		frequency[line_parts_right[i]] = frequency[line_parts_right[i]] + 1
	}

	for i := 0; i < len(line_parts_left); i++ {
		value := line_parts_left[i] * frequency[line_parts_left[i]]
		part2 += value
	}

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
