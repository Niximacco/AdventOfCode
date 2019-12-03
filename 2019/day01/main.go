package main

import (
	"bufio"
	"math"
	"os"
	"flag"
	"fmt"
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
	total2 := 0
	for _, line := range(input_file) {
		val, err := strconv.Atoi(line)
		check(err)

		value := calculate(val)
		fmt.Printf("%s: %d\n", line, value)
		total += value

		additional := calculate(value)
		total2 += value + additional
		for additional > 0 {
			additional = calculate(additional)
			fmt.Printf("add: %d\n", additional)

			if additional > 0 {
				total2 += additional
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("total2: %d\n", total2)

}

func calculate(input int) int {
	return int(math.Floor(float64(input) / 3) - 2)
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
