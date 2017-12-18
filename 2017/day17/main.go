package main

import (
	"bufio"
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
	part1(input_file)
	part2(input_file)
}

func part1(input_file []string) {
	var data []int
	step, err := strconv.Atoi(input_file[0])
	check(err)
	data = append(data, 0)
	maxIndex := 0
	currentPos := 0

	for i:=1; i<=2017; i++ {
		for j:=0; j<step; j++ {
			if currentPos == maxIndex {
				currentPos = 0
			} else {
				currentPos++
			}
		}

		var newData []int
		for index:=0; index<=currentPos; index++ {
			newData = append(newData, data[index])
		}

		newData = append(newData, i)

		for index:=currentPos+1; index<len(data); index++ {
			newData = append(newData, data[index])
		}


		data = newData
		maxIndex++
		currentPos++
		//fmt.Printf("%v\n", data)
	}

	fmt.Printf("Part 1 Solution: %d\n", data[findIndex(2017, data)+1])

}

func part2(input_file []string) {
	var data []int
	step, err := strconv.Atoi(input_file[0])
	check(err)
	data = append(data, 0)
	maxIndex := 0
	currentPos := 0
	lastZero := 0

	for i:=1; i<=50000000; i++ {
		fmt.Printf("\r%d", i)
		if maxIndex > 0 {
			currentPos = (currentPos + step) % maxIndex
		} else {
			currentPos = 0
		}
		//for j:=0; j<step; j++ {
		//	if currentPos == maxIndex {
		//		currentPos = 0
		//	} else {
		//		currentPos++
		//	}
		//}

		if currentPos == 0 {
			lastZero = maxIndex
		}

		maxIndex++
		currentPos++
		//fmt.Printf("%v\n", data)
	}

	fmt.Printf("Part 2 Solution: %d\n", lastZero)

}

func findIndex(match int, data []int) int {
	for i, element := range data {
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
