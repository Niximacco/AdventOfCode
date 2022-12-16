package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Type int

const (
	Air Type = iota
	Rock
	Sand
	StartSand
	Floor
)

type State int

const (
	Rest State = iota
	Falling
	Abyss
)

type Point struct {
	Type       Type
	Coordinate Coordinate
	State      State
}

type Board struct {
	MaxX           int
	MaxY           int
	MinX           int
	MinY           int
	SandSpawn      Coordinate
	Points         map[int]map[int]*Point
	NumRestingSand int
	FloorHeight    int
}

func (point *Point) ToString() (char string) {
	switch point.Type {
	case Air:
		char = "."
	case Rock:
		char = "#"
	case Sand:
		char = "o"
	case StartSand:
		char = "+"
	case Floor:
		char = "="
	default:
		char = " "
	}
	return
}

func (board *Board) Init(lines [][]Coordinate, startSand Coordinate) {
	board.MinX, board.MinY, board.MaxX, board.MaxY = 500, 0, 500, 0
	board.Points = make(map[int]map[int]*Point, 0)
	board.Points[startSand.X] = make(map[int]*Point)
	board.Points[startSand.X][startSand.Y] = &Point{
		Type:       StartSand,
		Coordinate: startSand,
	}

	board.SandSpawn = startSand
	board.NumRestingSand = 0

	for _, line := range lines {
		drawPoint := line[0]
		fmt.Printf("Start Point: %s\n", drawPoint.ToString())
		for i, coordinate := range line {
			if coordinate.X < board.MinX {
				board.MinX = coordinate.X
			}

			if coordinate.Y < board.MinY {
				board.MinY = coordinate.Y
			}

			if coordinate.X > board.MaxX {
				board.MaxX = coordinate.X
			}

			if coordinate.Y > board.MaxY {
				board.MaxY = coordinate.Y
			}

			if i == 0 {
				continue
			}

			for {
				if _, ok := board.Points[drawPoint.X]; !ok {
					board.Points[drawPoint.X] = make(map[int]*Point)
				}

				board.Points[drawPoint.X][drawPoint.Y] = &Point{
					Type:       Rock,
					Coordinate: drawPoint,
				}

				if drawPoint.Equals(&coordinate) {
					break
				}

				if drawPoint.X != coordinate.X {
					if drawPoint.X > coordinate.X {
						drawPoint.X--
					} else {
						drawPoint.X++
					}
				} else if drawPoint.Y != coordinate.Y {
					if drawPoint.Y > coordinate.Y {
						drawPoint.Y--
					} else {
						drawPoint.Y++
					}
				}
			}
		}
	}

	for x := board.MinX - 200; x <= board.MaxX+200; x++ {
		for y := board.MinY - 5; y <= board.MaxY+5; y++ {
			if _, ok := board.Points[x]; !ok {
				board.Points[x] = make(map[int]*Point)
			}
			if _, ok := board.Points[x][y]; !ok {
				board.Points[x][y] = &Point{
					Type: Air,
					Coordinate: Coordinate{
						X: x,
						Y: y,
					},
				}
			}
		}
	}

	board.FloorHeight = board.MaxY + 2

	for x := board.MinX - 200; x <= board.MaxX+200; x++ {
		board.Points[x][board.FloorHeight].Type = Floor
	}

}

func (board *Board) SpawnSand(part int) {
	sand := &Point{
		Type:       Sand,
		Coordinate: board.SandSpawn,
		State:      Falling,
	}
	if part == 1 {
		for sand.State != Abyss {
			sand.Tick(board)
			if sand.State == Rest {
				sand.Coordinate = board.SandSpawn
				sand.State = Falling
			}

			if sand.Coordinate.Y > board.MaxY {
				sand.State = Abyss
			}
		}
	}

	if part == 2 {
		fmt.Printf("Sand (%d) at: %s (at spawn? %v)\n", sand.State, sand.Coordinate.ToString(), sand.Coordinate.Equals(&board.SandSpawn))
		for !(sand.Coordinate.Equals(&board.SandSpawn) && sand.State == Rest) {
			if sand.State == Rest {
				sand.Coordinate = board.SandSpawn
				sand.State = Falling
			}
			fmt.Printf("  Sand at %s\n", sand.Coordinate.ToString())
			sand.Tick(board)
		}
	}
}

func (sand *Point) Tick(board *Board) {
	if board.Points[sand.Coordinate.X][sand.Coordinate.Y+1].Type == Air {
		sand.Coordinate.Y++
	} else if board.Points[sand.Coordinate.X-1][sand.Coordinate.Y+1].Type == Air {
		sand.Coordinate.X--
		sand.Coordinate.Y++
	} else if board.Points[sand.Coordinate.X+1][sand.Coordinate.Y+1].Type == Air {
		sand.Coordinate.X++
		sand.Coordinate.Y++
	} else {
		sand.State = Rest
		board.Points[sand.Coordinate.X][sand.Coordinate.Y].Type = Sand
		board.NumRestingSand++
	}

	fmt.Printf("  New Sand: %s (%d)\n", sand.Coordinate.ToString(), sand.State)
}

func (board *Board) PrintBoard() {
	for y := board.MinY; y <= board.MaxY; y++ {
		for x := board.MinX; x <= board.MaxX; x++ {
			fmt.Printf("%s", board.Points[x][y].ToString())
		}
		fmt.Printf("\n")
	}
}

type Coordinate struct {
	X int
	Y int
}

func (coordinate *Coordinate) ToString() string {
	return fmt.Sprintf("[%d, %d]", coordinate.X, coordinate.Y)
}

func (a *Coordinate) Equals(b *Coordinate) bool {
	return a.X == b.X && a.Y == b.Y
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
	part1, part2 := 0, 0
	lines := make([][]Coordinate, 0)
	for _, line := range input_file {
		coordinates := make([]Coordinate, 0)
		fmt.Printf("%s\n", line)
		points := strings.Split(line, " -> ")
		for _, point := range points {
			fmt.Printf("  -%s\n", point)
			pointParts := strings.Split(point, ",")
			x, _ := strconv.Atoi(pointParts[0])
			y, _ := strconv.Atoi(pointParts[1])
			coordinates = append(coordinates, Coordinate{
				X: x,
				Y: y,
			})
		}
		lines = append(lines, coordinates)
	}

	sandStart := Coordinate{
		X: 500,
		Y: 0,
	}

	board := Board{}
	board.Init(lines, sandStart)
	board.PrintBoard()
	board.SpawnSand(1)
	board.PrintBoard()
	part1 = board.NumRestingSand

	fmt.Printf("part 1 done\n")
	time.Sleep(1 * time.Second)

	board = Board{}
	board.Init(lines, sandStart)
	board.PrintBoard()
	board.SpawnSand(2)
	board.PrintBoard()
	part2 = board.NumRestingSand

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
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
