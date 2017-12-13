package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
	"strconv"
	"regexp"
	"strings"
)

var depthMap = make(map[int]int)

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
	for _, line := range input_file {
		depth, size := parseLine(line)
		depthMap[depth] = size
	}

	sev, _ := calcSeverityWithDelay(0, input_file, false)
	fmt.Printf("\nPart 1 solution: %d\n", sev)

	delay := 0
	caught := true

	for severity:=-1; severity!=0 || caught; delay++ {
		//fmt.Printf("\nDelay: %d\n", delay)
		severity, caught = calcSeverityWithDelay(delay, input_file, true)
		//fmt.Printf("\r%d", delay)
	}
	fmt.Printf("\nPart 2 solution: %d\n", delay-1)

}

func calcSeverityWithDelay(delay int, input_file []string, shortCircuit bool) (int, bool) {
	totalDepth, _ := parseLine(input_file[len(input_file)-1])
	severity := 0
	second := 0
	caughtInRun := false
	//fmt.Printf("Total Depth: %d\tSize at Depth: %d\n", totalDepth, sizeAtDepth)
	//fmt.Printf("DepthMap: %v\n\n", depthMap)
	second += delay

	for currentDepth:=0; currentDepth<=totalDepth; currentDepth++ {
		caught := false
		if _, ok := depthMap[currentDepth]; ok {
			if isCaught(currentDepth, second) {
				caught = true
				severity += (currentDepth * depthMap[currentDepth])
			}
		}
		second++
		if caught && shortCircuit {
			return 1, true
		}
	}

	return severity, caughtInRun
}

func isCaught(depth int, second int) bool {
	return second % (2* depthMap[depth] - 2) == 0
}

func parseLine(line string) (int,int) {
	line = strings.TrimSpace(line)
	re := regexp.MustCompile("(.+?): (.+)")
	matched := re.FindStringSubmatch(line)

	depth, err := strconv.Atoi(matched[1])
	check(err)

	sizeAtDepth, err := strconv.Atoi(matched[2])
	check(err)

	return depth, sizeAtDepth
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
