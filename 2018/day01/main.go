package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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
	total := 0
	repeatedTotal := false
	seenTwice := 0
	var seenList []int
	part1Total := 0
	part1Done := false

	for !repeatedTotal {
		for _, line := range input_file {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			total += i

			if !repeatedTotal {
				if contains(seenList, total) {
					repeatedTotal = true
					seenTwice = total
				}
			}
			seenList = append(seenList, total)
			if !part1Done {
				part1Done = true
				part1Total = total
			}
		}
	}

	fmt.Printf("PART 1: total: %d\n", part1Total)
	fmt.Printf("PART2: seenTwice: %d\n", seenTwice)
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

func contains(arr []int, num int) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
}
