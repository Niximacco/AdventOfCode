package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Location struct {
	X        int
	Y        int
	IsSeat   bool
	Occupied bool
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
	locations := parseLocations(input_file)

	resultPart1 := part1(locations)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(locations)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)
}

func part1(locations [][]Location) (numOccupied int) {
	changes := true
	for changes {
		changes, locations = processPart1(locations)
	}

	for _, row := range locations {
		for _, location := range row {
			if location.Occupied {
				numOccupied++
			}
		}
	}
	return
}

func part2(locations [][]Location) (numOccupied int) {
	changes := true
	iterations := 0
	printLocations(locations)

	for changes {
		iterations++
		fmt.Printf("\n%d iterations\n", iterations)
		changes, locations = processPart2(locations)
		printLocations(locations)
	}

	for _, row := range locations {
		for _, location := range row {
			if location.Occupied {
				numOccupied++
			}
		}
	}
	return
}

func processPart1(locations [][]Location) (changes bool, newLocations [][]Location) {
	for _, row := range locations {
		var newRow []Location
		for _, location := range row {
			currentState := location.Occupied
			location.calculateNextStatePart1(locations)
			newRow = append(newRow, location)
			if currentState != location.Occupied {
				changes = true
			}
		}
		newLocations = append(newLocations, newRow)
	}

	return
}

func processPart2(locations [][]Location) (changes bool, newLocations [][]Location) {
	for _, row := range locations {
		var newRow []Location
		for _, location := range row {
			currentState := location.Occupied
			location.calculateNextStatePart2(locations)
			newRow = append(newRow, location)
			if currentState != location.Occupied {
				changes = true
			}
		}
		newLocations = append(newLocations, newRow)
	}

	return
}

func (location *Location) calculateNextStatePart1(locations [][]Location) {
	if location.IsSeat {
		if location.Occupied {
			if location.numAdjacent(locations) > 3 {
				location.Occupied = false
			}
		} else {
			if location.numAdjacent(locations) == 0 {
				location.Occupied = true
			}
		}
	}
}

func (location *Location) calculateNextStatePart2(locations [][]Location) {
	if location.IsSeat {
		if location.Occupied {
			if location.numLineOfSight(locations) > 4 {
				location.Occupied = false
			}
		} else {
			if location.numLineOfSight(locations) == 0 {
				location.Occupied = true
			}
		}
	}
}

func (location *Location) numAdjacent(locations [][]Location) (numAdjacent int) {
	for yOffset := -1; yOffset < 2; yOffset++ {
		for xOffset := -1; xOffset < 2; xOffset++ {
			if yOffset == 0 && xOffset == 0 {
				continue
			}

			yCheck := location.Y + yOffset
			xCheck := location.X + xOffset

			if yCheck < 0 || yCheck >= len(locations) {
				continue
			}
			if xCheck < 0 || xCheck >= len(locations[location.Y]) {
				continue
			}

			if locations[yCheck][xCheck].Occupied {
				numAdjacent++
			}
		}
	}
	return
}

func (location *Location) numLineOfSight(locations [][]Location) (numLineOfSight int) {
	for yDirection := -1; yDirection < 2; yDirection++ {
		for xDirection := -1; xDirection < 2; xDirection++ {
			if yDirection == 0 && xDirection == 0 {
				continue
			}

			seatFound := false
			y := location.Y + yDirection
			x := location.X + xDirection

			for !seatFound {
				if y < 0 || y >= len(locations) {
					seatFound = true
					break
				}
				if x < 0 || x >= len(locations[location.Y]) {
					seatFound = true
					break
				}

				if locations[y][x].IsSeat {
					if locations[y][x].Occupied {
						numLineOfSight++
					}
					seatFound = true
				}

				y += yDirection
				x += xDirection
			}
		}
	}
	return
}

func printLocations(locations [][]Location) {
	for _, row := range locations {
		for _, location := range row {
			if !location.IsSeat {
				fmt.Printf(".")
			} else if location.IsSeat && location.Occupied {
				fmt.Printf("#")
			} else {
				fmt.Printf("L")
			}
		}
		fmt.Printf("\n")
	}
}

func parseLocations(input []string) (locations [][]Location) {
	for y, line := range input {
		var rowLocations []Location
		for x, char := range line {
			isSeat := string(char) == "L"

			rowLocations = append(rowLocations, Location{
				X:        x,
				Y:        y,
				IsSeat:   isSeat,
				Occupied: false,
			})
		}
		locations = append(locations, rowLocations)
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
