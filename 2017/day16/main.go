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
	var testMode bool
	if *boolPtr {
		filename = "input-test.txt"
		testMode = true
	} else {
		filename = "input.txt"
		testMode = false
	}
	input_file, err := readLines(filename)
	check(err)

	cacheLookup := make(map [string]string)
	var seenPos []string
	// Your Code goes below!
	var positions []string
	if testMode {
		positions = append(positions, "a", "b", "c", "d", "e")
	} else {
		positions = append(positions, "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p")
	}

	startingPositions := positions
	instructions := strings.Split(input_file[0], ",")


	positions = doDance(instructions, positions)
	fmt.Printf("%v\n", positions)
	fmt.Printf("Part 1 Solution: %s", makeString(positions))

	fmt.Printf("\n\n")

	fmt.Printf("Starting: %v\n", startingPositions)
	var indexChanges []int
	for i:=0; i<len(positions); i++ {
		indexChanges = append(indexChanges, findIndex(startingPositions[i], positions))
	}

	fmt.Printf("Index Changes: %v\n", indexChanges)
	cacheHit := false
	loopSize := 1
	i:=0
	seenPos = append(seenPos, makeString(startingPositions))
	positions = startingPositions
	positions = doDance(instructions, positions)

	for !cacheHit {
		//fmt.Printf("%v\n", cacheLookup)
		if findIndex(makeString(positions), seenPos) >= 0 {
			fmt.Printf("HIT\t")
			cacheHit = true
			break
		} else {
			seenPos = append(seenPos, makeString(positions))
			prevPos := positions
			loopSize++
			fmt.Printf("NEW\t")
			positions = doDance(instructions, positions)
			cacheLookup[makeString(prevPos)] = makeString(positions)
			fmt.Printf("%d\t%s\tLoopSize: %d\n", i, makeString(positions), loopSize)
		}
		i++
	}
	fmt.Printf("%d\t%s\tLoopSize: %d\n", i, makeString(positions), loopSize)

	startIndex := 0
	for i:=1; i<1000000000; i+=loopSize {
		startIndex=i
	}

	positions = startingPositions
	for i:=startIndex; i<=1000000000; i++ {
		prevPos := positions
		if val, ok := cacheLookup[makeString(positions)]; ok {
			fmt.Printf("HIT\t")
			positions = strings.Split(val, "")
		} else {
			fmt.Printf("NEW\t")
			positions = doDance(instructions, positions)
			cacheLookup[makeString(prevPos)] = makeString(positions)
		}
		fmt.Printf("%d\t%s\n", i, makeString(positions), loopSize)
	}

	fmt.Printf("Part 2 Solution: %s", makeString(positions))

	fmt.Printf("\n\n")

}

func makeString(positions []string) string {
	return strings.Join(positions, "")
}

//func applyIndexChanges(positions []string, indexChanges []int) []string {
//	endPositions := make([]string, 16, 16)
//	for start, end := range indexChanges {
//		endPositions[end] = positions[start]
//	}
//
//	return endPositions
//}

func doDance(instructions []string, positions []string) []string {
	for _, instruction := range instructions {
		switch string(instruction[0]) {
		case "s":
			positions = spin(instruction, positions)
		case "x":
			positions = exchange(instruction, positions)
		case "p":
			positions = partner(instruction, positions)
		}
	}
	return positions
}

func spin(instruction string, positions []string) []string {
	re := regexp.MustCompile("s(.+)")
	matched := re.FindStringSubmatch(instruction)
	num, err := strconv.Atoi(matched[1])
	check(err)

	return append(positions[len(positions)-num:], positions[:len(positions)-num]...)
}

func exchange(instruction string, positions []string) []string {
	re := regexp.MustCompile("x(.+)/(.+)")
	matched := re.FindStringSubmatch(instruction)
	swap1, err := strconv.Atoi(matched[1])
	check(err)

	swap2, err := strconv.Atoi(matched[2])
	check(err)

	temp := positions[swap1]
	positions[swap1] = positions[swap2]
	positions[swap2] = temp

	return positions
}

func partner(instruction string, positions []string) []string {
	re := regexp.MustCompile("p(.+)/(.+)")
	matched := re.FindStringSubmatch(instruction)

	index1 := findIndex(matched[1], positions)
	index2 := findIndex(matched[2], positions)
	instruction = fmt.Sprintf("x%d/%d", index1, index2)

	return exchange(instruction, positions)
}

func findIndex(match string, positions []string) int {
	for i, element := range positions {
		if element == match {
			return i
		}
	}
	return -1
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
