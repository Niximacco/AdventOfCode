package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ShipPart1 struct {
	X      int
	Y      int
	Facing Direction
}

type ShipPart2 struct {
	X         int
	Y         int
	WaypointX int
	WaypointY int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
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

	resultPart1 := part1(input_file)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(input_file)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(lines []string) (manhattanDistance int) {
	ship := ShipPart1{
		0,
		0,
		East,
	}

	for _, line := range lines {
		instruction, amount := parseLine(line)

		switch instruction {
		case "N":
			ship.north(amount)
		case "S":
			ship.south(amount)
		case "E":
			ship.east(amount)
		case "W":
			ship.west(amount)
		case "F":
			ship.forward(amount)
		case "L", "R":
			ship.rotate(instruction, amount)
		}
	}

	manhattanDistance = ship.manhattanDistance()
	return
}

func part2(lines []string) (manhattanDistance int) {
	ship := ShipPart2{
		0,
		0,
		10,
		1,
	}
	fmt.Printf("Ship @ [%d, %d], Waypoint [%d, %d]\n", ship.X, ship.Y, ship.WaypointX, ship.WaypointY)

	for _, line := range lines {
		fmt.Printf("%s:\t", line)
		instruction, amount := parseLine(line)

		switch instruction {
		case "N":
			ship.north(amount)
		case "S":
			ship.south(amount)
		case "E":
			ship.east(amount)
		case "W":
			ship.west(amount)
		case "F":
			ship.forward(amount)
		case "L", "R":
			ship.rotate(instruction, amount)
		}

		fmt.Printf("Ship @ [%d, %d], Waypoint [%d, %d]\n", ship.X, ship.Y, ship.WaypointX, ship.WaypointY)
	}

	manhattanDistance = ship.manhattanDistance()
	return
}

func (ship *ShipPart1) north(amount int) {
	ship.Y += amount
}

func (ship *ShipPart1) south(amount int) {
	ship.Y -= amount
}

func (ship *ShipPart1) east(amount int) {
	ship.X += amount
}

func (ship *ShipPart1) west(amount int) {
	ship.X -= amount
}

func (ship *ShipPart1) rotate(direction string, amount int) {
	numChange := amount / 90
	currentDir := ship.Facing
	for numChange > 0 {
		switch direction {
		case "L":
			currentDir--
		case "R":
			currentDir++
		}

		numChange--
	}

	if currentDir < 0 {
		currentDir += 4
	}
	ship.Facing = currentDir % 4
}

func (ship *ShipPart1) forward(amount int) {
	switch ship.Facing {
	case North:
		ship.north(amount)
	case South:
		ship.south(amount)
	case East:
		ship.east(amount)
	case West:
		ship.west(amount)
	}
}

func (ship *ShipPart1) manhattanDistance() (distance int) {
	distance = Abs(ship.Y) + Abs(ship.X)
	return
}

func (ship *ShipPart2) north(amount int) {
	ship.WaypointY += amount
}

func (ship *ShipPart2) east(amount int) {
	ship.WaypointX += amount
}

func (ship *ShipPart2) south(amount int) {
	ship.WaypointY -= amount
}

func (ship *ShipPart2) west(amount int) {
	ship.WaypointX -= amount
}

func (ship *ShipPart2) forward(amount int) {
	ship.X += ship.WaypointX * amount
	ship.Y += ship.WaypointY * amount
}

func (ship *ShipPart2) rotate(direction string, amount int) {
	numChange := amount / 90
	for numChange > 0 {
		switch direction {
		case "L":
			// [x, y] -> [-y, x]
			newY := ship.WaypointX
			ship.WaypointX = -ship.WaypointY
			ship.WaypointY = newY
		case "R":
			// [x, y] -> [y, -x]
			newX := ship.WaypointY
			ship.WaypointY = -ship.WaypointX
			ship.WaypointX = newX
		}

		numChange--
	}
}

func (ship *ShipPart2) manhattanDistance() (distance int) {
	distance = Abs(ship.Y) + Abs(ship.X)
	return
}

func parseLine(line string) (instruction string, amount int) {
	lineMatches := regexp.MustCompile("^([A-Z])(\\d+)$").FindStringSubmatch(line)

	if len(lineMatches) > 2 {
		instruction = lineMatches[1]
		amount, _ = strconv.Atoi(lineMatches[2])
	}

	return
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
