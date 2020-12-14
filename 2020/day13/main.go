package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bus struct {
	Id     int
	Offset int
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
	arrivalTime, busLines := parseInput(input_file)

	resultPart1 := part1(arrivalTime, busLines)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(busLines)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(arrivalTime int, busLines []Bus) (result int) {
	busFound := false
	currentTime := arrivalTime

	for !busFound {
		for _, bus := range busLines {
			if bus.Id == 0 {
				continue
			}
			if currentTime%bus.Id == 0 {
				busFound = true
				result = (currentTime - arrivalTime) * bus.Id
			}
		}
		currentTime++
	}

	return
}

func part2(busLines []Bus) (currentTime int) {
	for {
		timeDelta := 1
		valid := true

		for _, bus := range busLines {
			if (currentTime+bus.Offset)%bus.Id != 0 {
				valid = false
				break
			}

			timeDelta *= bus.Id
		}

		if valid {
			return
		}

		currentTime += timeDelta
	}
}

func parseInput(input []string) (arrivalTime int, buses []Bus) {
	arrivalTime, _ = strconv.Atoi(input[0])

	inputBusLines := strings.Split(input[1], ",")
	for i, inputBusLine := range inputBusLines {
		if inputBusLine != "x" {
			busLine, _ := strconv.Atoi(inputBusLine)
			buses = append(buses, Bus{busLine, i})
		}
	}

	return
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
