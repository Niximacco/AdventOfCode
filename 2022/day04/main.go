package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	Start int
	End   int
}

func (thisElf *Elf) Contains(thatElf Elf) (contains bool) {
	fmt.Printf("thisElf: %v thatElf: %v\n", thisElf, thatElf)
	contains = thisElf.Start <= thatElf.Start && thisElf.End >= thatElf.End
	return
}

func (thisElf *Elf) Overlaps(thatElf Elf) (overlap bool) {
	overlap = thisElf.Start <= thatElf.Start && thisElf.End >= thatElf.Start
	return
}

func Create(segment string) (thisElf Elf) {
	parts := strings.Split(segment, "-")
	thisElf.Start, _ = strconv.Atoi(parts[0])
	thisElf.End, _ = strconv.Atoi(parts[1])
	return
}

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
	part1 := 0
	part2 := 0
	for _, line := range input_file {
		fmt.Printf("%s\n", line)
		elfs := strings.Split(line, ",")
		elf1 := Create(elfs[0])
		elf2 := Create(elfs[1])

		if elf1.Contains(elf2) || elf2.Contains(elf1) {
			fmt.Printf("fully contained.\n")
			part1++
		}

		if elf1.Overlaps(elf2) || elf2.Overlaps(elf1) {
			fmt.Printf("overlaps\n")
			part2++
		}

		println()
	}

	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
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
