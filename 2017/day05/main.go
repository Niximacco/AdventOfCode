package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	input_file, err := readLines("input.txt")
	check(err)

	part1(input_file)
	part2(input_file)
}

func part1(input_file []string) {
	var nums []int
	for i := 0; i<len(input_file); i++ {
		num, err := strconv.Atoi(strings.TrimSuffix(input_file[i], "\n"))
		check(err)

		nums = append(nums, num)
	}
	index := 0
	steps1 := 0

	for index < len(nums) {
		steps1++
		instruction := nums[index]
		nums[index]++
		index += instruction
		//fmt.Printf("Steps: %d\tInstruction: %d\tNext Index: %d\n", steps, instruction, index)
	}
	fmt.Printf("Num Steps Part1: %d\n", steps1)
}

func part2(input_file []string) {
	var nums []int
	for i := 0; i<len(input_file); i++ {
		num, err := strconv.Atoi(strings.TrimSuffix(input_file[i], "\n"))
		check(err)

		nums = append(nums, num)
	}
	index := 0
	steps2 := 0

	for index < len(nums) {
		steps2++
		instruction := nums[index]
		if instruction >= 3 {
			nums[index]--
		} else {
			nums[index]++
		}
		index += instruction
		//fmt.Printf("Steps: %d\tInstruction: %d\tNext Index: %d\n", steps, instruction, index)
	}
	fmt.Printf("Num Steps Part2: %d\n", steps2)
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
