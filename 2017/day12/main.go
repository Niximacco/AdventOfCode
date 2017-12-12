package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
	"strings"
	"regexp"
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

	var seenNumbers []int
	numGroups := 0
	group := calculateGroupFromNum(0, input_file)
	fmt.Printf("Part 1 Solution: %d\t%v\n", len(group), group)

	for _, num := range group {
		seenNumbers = append(seenNumbers, num)
	}

	numGroups++

	for i:=0; i<len(input_file); i++ {
		if !numInGroup(i, seenNumbers) {
			numGroups++
			group := calculateGroupFromNum(i, input_file)
			for _, num := range group {
				seenNumbers = append(seenNumbers, num)
			}
		}
	}
	fmt.Printf("Part 2 Solution: %d\n", numGroups)

}

func calculateGroupFromNum(matchNum int, input_file []string) []int{
	previousResult := -1
	currentResult := 0
	noChange := false
	var group []int
	group = append(group, matchNum)
	for !noChange {
		for _, line := range input_file {
			line = strings.TrimSpace(line)
			re := regexp.MustCompile("(.+?) <-> (.+)")
			matched := re.FindStringSubmatch(line)
			num, err := strconv.Atoi(matched[1])

			matches := strings.Split(matched[2], ",")
			var pairs []int

			for _, match := range matches {
				match, err := strconv.Atoi(strings.TrimSpace(match))
				check(err)

				pairs = append(pairs, match)
			}

			check(err)

			//fmt.Printf("Num: %d\tPairs: %v\n", num, pairs)
			if numInGroup(num, group) {
				for _, groupNum := range pairs {
					if !numInGroup(groupNum, group) {
						group = append(group, groupNum)
					}
				}
			}

			//fmt.Printf("Num: %d\tPairs: %v\t\tgroup: %v\n", num, pairs, group)
		}
		currentResult = len(group)
		if currentResult == previousResult {
			noChange = true
		}
		previousResult = currentResult
	}
	return group
}

func numInGroup(num int, group []int) bool {
	for _, groupNum := range group {
		if num == groupNum {
			return true
		}
	}

	return false
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
