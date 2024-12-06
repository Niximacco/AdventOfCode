package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Seen struct {
	up    bool
	right bool
	down  bool
	left  bool
}
type Space struct {
	Occupied   bool
	Visited    bool
	Horizontal bool
	Vertical   bool
	Option     bool
	Seen       Seen
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

func (direction direction) Char() rune {
	switch direction {
	case up:
		return '^'
	case right:
		return '>'
	case down:
		return 'v'
	case left:
		return '<'
	}

	return 'G'
}

func (direction direction) GetRotationDirection() direction {
	switch direction {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}

	return up
}

type Guard struct {
	Facing direction
	X      int
	Y      int
}

type Board struct {
	Guard      Guard
	StartGuard Guard
	Spaces     [][]Space
	Height     int
	Width      int
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
	board := Board{}
	board.Guard = Guard{
		Facing: up,
		X:      0,
		Y:      0,
	}
	startGuard := Guard{
		Facing: 0,
		X:      0,
		Y:      0,
	}

	board.Height = len(input_file)
	board.Width = len(input_file[0])
	board.Spaces = make([][]Space, board.Width)
	for i, _ := range board.Spaces {
		board.Spaces[i] = make([]Space, board.Height)
		for j, _ := range board.Spaces[i] {
			board.Spaces[i][j] = Space{}
		}
	}

	fmt.Printf("%+v\n", board)

	for y, line := range input_file {
		for x, char := range line {
			if char == '.' {
				board.Spaces[x][y] = Space{
					Occupied: false,
					Visited:  false,
					Seen:     Seen{},
				}
			} else if char == '#' {
				fmt.Printf("# [%d, %d]\n", x, y)
				board.Spaces[x][y] = Space{
					Occupied: true,
					Visited:  false,
					Seen:     Seen{},
				}
			} else if char == '^' {
				fmt.Printf("^ [%d, %d]\n", x, y)
				board.Spaces[x][y] = Space{
					Occupied:   false,
					Visited:    true,
					Horizontal: false,
					Vertical:   true,
					Seen:       Seen{up: true},
				}
				startGuard.X = x
				startGuard.Y = y
				board.StartGuard = startGuard
				board.Guard = startGuard
			}
		}
	}

	board.Print()
	fmt.Println()
	board.Simulate()
	board.Print()
	part1 = board.CountVisited()

	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			fmt.Printf("%d, %d\n", x, y)
			board.Reset()
			if x == board.Guard.X && y == board.Guard.Y {
				continue
			}

			if board.Spaces[x][y].Occupied {
				continue
			}

			board.Spaces[x][y].Occupied = true
			looping := board.Simulate()
			if looping {
				part2++
			}
			board.Spaces[x][y].Occupied = false
		}
	}

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func (board *Board) Print() {
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			space := board.Spaces[x][y]
			if space.Occupied {
				fmt.Printf("#")
			} else if board.Guard.X == x && board.Guard.Y == y {
				fmt.Printf("%c", board.Guard.Facing.Char())
			} else if space.Visited {
				if space.Horizontal && !space.Vertical {
					fmt.Printf("-")
				} else if space.Vertical && !space.Horizontal {
					fmt.Printf("|")
				} else {
					fmt.Printf("+")
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func (board *Board) CountVisited() (count int) {
	for _, row := range board.Spaces {
		for _, space := range row {
			if space.Visited {
				count++
			}
		}
	}
	return
}

func (board *Board) Simulate() (looping bool) {
	for board.NextMoveInBounds() {
		//fmt.Printf("Guard is facing %c at [%d, %d], moving to ", board.Guard.Facing.Char(), board.Guard.X, board.Guard.Y)
		// If space in front is empty, move forward
		nextX, nextY := board.GetNextGuardCoordinates()
		//fmt.Printf("[%d, %d]\n", nextX, nextY)
		if !board.Spaces[nextX][nextY].Occupied {
			board.Guard.X = nextX
			board.Guard.Y = nextY
			board.Spaces[nextX][nextY].Visited = true

			switch board.Guard.Facing {
			case up:
				if board.Spaces[nextX][nextY].Seen.up {
					return true
				}
				board.Spaces[nextX][nextY].Seen.up = true
			case right:
				if board.Spaces[nextX][nextY].Seen.right {
					return true
				}
				board.Spaces[nextX][nextY].Seen.right = true
			case down:
				if board.Spaces[nextX][nextY].Seen.down {
					return true
				}
				board.Spaces[nextX][nextY].Seen.down = true
			case left:
				if board.Spaces[nextX][nextY].Seen.left {
					return true
				}
				board.Spaces[nextX][nextY].Seen.left = true
			}

			if board.Guard.Facing == up || board.Guard.Facing == down {
				board.Spaces[nextX][nextY].Vertical = true
			}

			if board.Guard.Facing == left || board.Guard.Facing == right {
				board.Spaces[nextX][nextY].Horizontal = true
			}
		} else {
			// rotate 90 degrees
			board.Guard.Facing = board.Guard.Facing.GetRotationDirection()
			board.Spaces[board.Guard.X][board.Guard.Y].Vertical = true
			board.Spaces[board.Guard.X][board.Guard.Y].Horizontal = true
		}
	}

	return false
}

func (board *Board) GetNextGuardCoordinates() (x int, y int) {
	switch board.Guard.Facing {
	case up:
		return board.Guard.X, board.Guard.Y - 1
	case right:
		return board.Guard.X + 1, board.Guard.Y
	case down:
		return board.Guard.X, board.Guard.Y + 1
	case left:
		return board.Guard.X - 1, board.Guard.Y
	}

	return board.Guard.X, board.Guard.Y
}

func (board *Board) NextMoveInBounds() bool {
	switch board.Guard.Facing {
	case up:
		return !(board.Guard.Y == 0)
	case right:
		return !(board.Guard.X == board.Width-1)
	case down:
		return !(board.Guard.Y == board.Height-1)
	case left:
		return !(board.Guard.X == 0)
	}

	return true
}

func (board *Board) Reset() {
	board.Guard = board.StartGuard
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			board.Spaces[x][y].Seen = Seen{}
			board.Spaces[x][y].Visited = false
			board.Spaces[x][y].Vertical = false
			board.Spaces[x][y].Horizontal = false
		}
	}

	board.Spaces[board.Guard.X][board.Guard.Y].Visited = true
	board.Spaces[board.Guard.X][board.Guard.Y].Seen.up = true
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
