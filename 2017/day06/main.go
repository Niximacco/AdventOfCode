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

	var states []string
	var nums []int
	stringFields := strings.Fields(input_file[0])
	for _, field := range stringFields {
		number, err := strconv.Atoi(field)
		check(err)

		nums = append(nums, number)
	}

	stateString := getStateString(nums)
	numCycles := 0

	for stateStringNotSeenBefore(stateString, states) {
		states = append(states, stateString)
		numCycles++

		index := getHighestIndex(nums)
		nums = distributeFromIndex(index, nums)

		stateString = getStateString(nums)
		//fmt.Printf("State: %s\nStates: %s\n\n", stateString, states)
	}
	fmt.Printf("Num Cycles: %d\n", numCycles)
	fmt.Printf("Num Since: %d\n", numCycles-indexSeenBefore(stateString, states))
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

func getStateString(nums []int) string {
	stateString := ""
	for _, number := range nums {
		stateString = stateString + fmt.Sprintf("%d", number)
	}

	return stateString
}

func stateStringNotSeenBefore(state string, states []string) bool {
	for _, checkState := range states {
		if state == checkState {
			return false
		}
	}
	return true
}

func indexSeenBefore(state string, states []string) int {
	for i, checkState := range states {
		if state == checkState {
			return i
		}
	}
	return -1
}

func getHighestIndex(nums []int) int{
	max := 0
	index := -1
	for i, num := range nums {
		if num > max {
			max=num
			index=i
		}
	}
	return index
}

func distributeFromIndex(index int, nums []int) []int {
	n := nums[index]
	nums[index] = 0
	for n > 0 {
		index++
		if index >= len(nums) {
			index=0
		}
		nums[index]++
		n--
	}
	return nums
}