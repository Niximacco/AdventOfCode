package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
	"github.com/pmcxs/hexgrid"
	"strings"
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
	x := 0
	y := 0
	max := 0

	hexagonA := hexgrid.NewHex(0,0)
	directions := strings.Split(input_file[0], ",")
	for _, direction := range directions {
		switch direction {
		case "n":
			y--
		case "ne":
			x++
			y--
		case "se":
			x++
		case "s":
			y++
		case "sw":
			x--
			y++
		case "nw":
			x--
		}
		currentHex := hexgrid.NewHex(y, x)
		currentDist := hexgrid.HexDistance(hexagonA, currentHex)
		if currentDist > max {
			max = currentDist
		}
	}
	hexagonB := hexgrid.NewHex(y, x)
	distance := hexgrid.HexDistance(hexagonA, hexagonB)
	fmt.Printf("Part 1 solution: %d\n", distance)
	fmt.Printf("Part 2 solution: %d\n", max)



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
