package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	Height   int
	Letter   string
	X        int
	Y        int
	Checked  bool
	Previous *Point
	Step     int
}

type Result struct {
	Stops int
}

var points [][]*Point
var maxStops int
var start *Point
var end *Point

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
	part1, part2 := 0, 0
	points = make([][]*Point, 0)

	for r, line := range input_file {
		points = append(points, make([]*Point, 0))
		for c, char := range line {
			height := 0
			switch string(char) {
			case "S":
				height = 1
			case "E":
				height = 26
			default:
				height = int(char) - 96
			}
			point := Point{
				Height:   height,
				Letter:   string(char),
				X:        c,
				Y:        r,
				Checked:  false,
				Step:     0,
				Previous: nil,
			}

			if point.Letter == "S" {
				start = &point
			}

			if point.Letter == "E" {
				end = &point
			}

			points[r] = append(points[r], &point)
		}
	}

	fmt.Printf("Start: [%d, %d]\n", start.X, start.Y)
	fmt.Printf("End: [%d, %d]\n", end.X, end.Y)
	maxStops = len(points) * len(points[0])

	nextChecks := make(map[*Point]bool, 0)
	for neighbor, _ := range start.GetValidNeighbors() {
		neighbor.Previous = start
		nextChecks[neighbor] = true
	}

	start.Checked = true
	part1 = ExpandSearch(nextChecks, 1)
	PrintPath()

	part2 = Part2()

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func Part2() int {
	results := make([]int, 0)
	for r, _ := range points {
		for c, _ := range points[r] {
			fmt.Printf("[%d, %d]\n", r, c)
			if points[r][c].Height == 1 {
				ClearChecks()
				nextChecks := make(map[*Point]bool, 0)
				for neighbor, _ := range points[r][c].GetValidNeighbors() {
					neighbor.Previous = points[r][c]
					nextChecks[neighbor] = true
				}

				points[r][c].Checked = true
				result := ExpandSearch(nextChecks, 1)
				if result != -1 {
					results = append(results, result)
				}
			}
		}
	}

	sort.Ints(results)
	return results[0]
}

func ClearChecks() {
	for r, _ := range points {
		for c, _ := range points[r] {
			points[r][c].Checked = false
		}
	}
}

func ExpandSearch(edges map[*Point]bool, currentSteps int) (step int) {
	step = currentSteps
	nextChecks := make(map[*Point]bool, 0)
	for point, _ := range edges {
		point.Step = step
		if point == end {
			return
		}

		point.Checked = true
		for neighbor, _ := range point.GetValidNeighbors() {
			neighbor.Previous = point
			nextChecks[neighbor] = true
		}
	}
	//PrintCheckedPoints()

	if currentSteps > maxStops {
		return -1
	}

	return ExpandSearch(nextChecks, currentSteps+1)
}

func (currentPoint *Point) GetValidNeighbors() (neighbors map[*Point]bool) {
	neighbors = make(map[*Point]bool, 0)
	if currentPoint.Y > 0 {
		checkPoint := points[currentPoint.Y-1][currentPoint.X]
		if checkPoint.Height <= currentPoint.Height+1 {
			if !checkPoint.Checked {
				neighbors[checkPoint] = true
			}
		}
	}

	if currentPoint.Y < len(points)-1 {
		checkPoint := points[currentPoint.Y+1][currentPoint.X]
		if checkPoint.Height <= currentPoint.Height+1 {
			if !checkPoint.Checked {
				neighbors[checkPoint] = true
			}
		}
	}

	if currentPoint.X > 0 {
		checkPoint := points[currentPoint.Y][currentPoint.X-1]
		if checkPoint.Height <= currentPoint.Height+1 {
			if !checkPoint.Checked {
				neighbors[checkPoint] = true
			}
		}
	}

	if currentPoint.X < len(points[currentPoint.Y])-1 {
		checkPoint := points[currentPoint.Y][currentPoint.X+1]
		if checkPoint.Height <= currentPoint.Height+1 {
			if !checkPoint.Checked {
				neighbors[checkPoint] = true
			}
		}
	}

	return
}

func PrintPath() {
	fmt.Printf("Path:\n")
	highlight := make(map[*Point]bool)
	printPoint := end
	for printPoint.Previous != nil {
		highlight[printPoint] = true
		printPoint = printPoint.Previous
	}

	for r, _ := range points {
		for c, _ := range points[r] {
			if highlight[points[r][c]] {
				//fmt.Printf("%s", strings.ToUpper(points[r][c].Letter))
				fmt.Printf(".")
			} else {
				fmt.Printf("%s", points[r][c].Letter)
			}
		}
		println()
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
