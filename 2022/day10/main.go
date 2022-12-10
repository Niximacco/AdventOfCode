package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	Cycle      int
	CompleteOn int
	Operation  string
	Value      int
	X          int
	NextCheck  int
	CheckDelta int
	LineNum    int
	PixelPos   int
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
	cycles := make(map[string]int)
	cycles["noop"] = 0
	cycles["addx"] = 1

	screen := make([]string, 6)

	part1State := State{
		Cycle:      0,
		CompleteOn: 0,
		Operation:  "",
		Value:      0,
		X:          1,
		NextCheck:  20,
		CheckDelta: 40,
		LineNum:    0,
		PixelPos:   0,
	}

	cursor := make([]string, 40)
	crtRow := make([]string, 0)

	for cycle := 1; cycle <= 240; cycle++ {
		// Start
		part1State.Cycle = cycle
		for i := 0; i < len(cursor); i++ {
			if part1State.X+1 == i {
				cursor[i] = "#"
			} else if part1State.X == i {
				cursor[i] = "#"
			} else if part1State.X-1 == i {
				cursor[i] = "#"
			} else {
				cursor[i] = "."
			}
		}
		fmt.Printf("Sprite Position: %s\n\n", strings.Join(cursor, ""))

		if part1State.Operation == "" {
			commandParts := strings.Fields(input_file[part1State.LineNum])
			command := commandParts[0]
			amount := 0
			if len(commandParts) > 1 {
				amount, err = strconv.Atoi(commandParts[1])
				check(err)
			}
			part1State.Value = amount
			part1State.LineNum++

			part1State.Operation = command
			part1State.CompleteOn = cycle + cycles[command]
		}
		fmt.Printf("Start cycle\t%d: Begin Executing %s %d\n", part1State.Cycle, part1State.Operation, part1State.Value)

		//During
		if part1State.Cycle == part1State.NextCheck {
			signalStrength := part1State.Cycle * part1State.X
			part1 += signalStrength
			part1State.NextCheck += part1State.CheckDelta
		}

		fmt.Printf("During cycle\t%d: CRT draws pixel in position %d\n", part1State.Cycle, part1State.PixelPos)
		crtRow = append(crtRow, cursor[part1State.PixelPos])
		fmt.Printf("Current CRT row: %s\n", strings.Join(crtRow, ""))

		part1State.PixelPos++
		if part1State.PixelPos%40 == 0 {
			part1State.PixelPos = 0
			screen = append(screen, strings.Join(crtRow, ""))
			crtRow = make([]string, 0)
		}

		if part1State.Cycle == part1State.CompleteOn {
			if part1State.Operation == "addx" {
				part1State.X += part1State.Value
				fmt.Printf("End of cycle %d: finish executing %s %d (Register X is now %d)\n", part1State.Cycle, part1State.Operation, part1State.Value, part1State.X)
			} else {
				fmt.Printf("End of cycle %d: finish executing %s\n", part1State.Cycle, part1State.Operation)

			}
			part1State.Operation = ""
		}
	}

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: \n")
	PrintScreen(screen)
}

func PrintScreen(screen []string) {
	for _, line := range screen {
		fmt.Printf("%s\n", line)
	}
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
