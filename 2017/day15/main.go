package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
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
	genVals := make(map[string]int)
	genFactors := make(map[string]int)

	var gens []string

	genFactors["A"]=16807
	genFactors["B"]=48271

	numIters := 1

	re := regexp.MustCompile("Generator (.) starts with (.+)")
	for _, line := range input_file {
		matched := re.FindStringSubmatch(line)
		num, err := strconv.Atoi(matched[2])
		check(err)
		genVals[matched[1]]=num
		gens = append(gens, matched[1])

	}

	numMatches := 40000000

	for i:=0; i<numIters; i++ {
		for _, gen := range gens {
			genVals[gen] = genVals[gen] * genFactors[gen] % 2147483647
		}
		binValueA := fmt.Sprintf("%032b", genVals["A"])
		binValueB := fmt.Sprintf("%032b", genVals["B"])
		if binValueA[len(binValueA)-16:] == binValueB[len(binValueB)-16:] {
			numMatches++
		}
		fmt.Printf("\r%d", i)
	}

	fmt.Printf("Part 1 Solution: %d\n", numMatches)




	numIters = 5000000
	numMatches = 0
	for _, line := range input_file {
		matched := re.FindStringSubmatch(line)
		num, err := strconv.Atoi(matched[2])
		check(err)
		genVals[matched[1]]=num
	}

	for i:=0; i<numIters; i++ {
		genVals["A"] = getNextValuePart2(genVals["A"], genFactors["A"], 4)
		genVals["B"] = getNextValuePart2(genVals["B"], genFactors["B"], 8)
		binValueA := fmt.Sprintf("%032b", genVals["A"])
		binValueB := fmt.Sprintf("%032b", genVals["B"])
		if binValueA[len(binValueA)-16:] == binValueB[len(binValueB)-16:] {
			numMatches++
		}
		fmt.Printf("\r%d", i)
	}

	fmt.Printf("Part 2 Solution: %d\n", numMatches)
}

func getNextValuePart2(currentValue int, factor int, multiple int) int {
	value := (currentValue * factor) % 2147483647
	if value % multiple == 0 {
		return value
	}
	return getNextValuePart2(value, factor, multiple)
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
