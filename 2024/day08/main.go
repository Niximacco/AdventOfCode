package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Space struct {
	Tower          rune
	Antinodes      []rune
	AntinodesPart2 []rune
	X              int
	Y              int
}

type Board struct {
	Spaces [][]Space
	Height int
	Width  int
	Towers []Space
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

	board.Height = len(input_file)
	board.Width = len(input_file[0])
	board.Spaces = make([][]Space, board.Width)
	board.Towers = []Space{}
	for i, _ := range board.Spaces {
		board.Spaces[i] = make([]Space, board.Height)
		for j, _ := range board.Spaces[i] {
			board.Spaces[i][j] = Space{}
		}
	}

	for y, line := range input_file {
		for x, char := range line {
			if char == '.' {
				board.Spaces[x][y] = Space{Tower: ' ', Antinodes: []rune{}, X: x, Y: y}
			} else {
				fmt.Printf("^ [%d, %d]: %c\n", x, y, char)
				board.Spaces[x][y] = Space{
					Tower:     char,
					Antinodes: []rune{},
					X:         x,
					Y:         y,
				}

				board.Towers = append(board.Towers, board.Spaces[x][y])
			}
		}
	}

	board.Print()
	fmt.Printf("Towers: %+v\n", board.Towers)

	for _, tower := range board.Towers {
		// Get all matching towers
		matches := board.GetMatchingTowers(tower)
		fmt.Printf("%c [%d, %d] matches: %+v\n", tower.Tower, tower.X, tower.Y, matches)

		for _, match := range matches {
			deltaX, deltaY := tower.GetDeltasFromSpace(match)
			antiNodeX := tower.X + deltaX
			antiNodeY := tower.Y + deltaY
			board.Spaces[tower.X][tower.Y].AntinodesPart2 = append(board.Spaces[tower.X][tower.Y].AntinodesPart2, tower.Tower)

			if board.InBounds(antiNodeX, antiNodeY) {
				fmt.Printf("\tAntinode at [%d, %d]\n", antiNodeX, antiNodeY)
				board.Spaces[antiNodeX][antiNodeY].Antinodes = append(board.Spaces[antiNodeX][antiNodeY].Antinodes, match.Tower)
			}

			for board.InBounds(antiNodeX, antiNodeY) {
				board.Spaces[antiNodeX][antiNodeY].AntinodesPart2 = append(board.Spaces[antiNodeX][antiNodeY].AntinodesPart2, match.Tower)
				antiNodeX += deltaX
				antiNodeY += deltaY
			}
		}
	}

	println()
	board.Print()
	part1 = board.CountAntinodes()
	part2 = board.CountAntinodesPart2()

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func (board *Board) InBounds(x int, y int) bool {
	if x >= board.Width || x < 0 {
		return false
	}

	if y >= board.Height || y < 0 {
		return false
	}

	return true
}

func (space1 *Space) GetDeltasFromSpace(space2 Space) (x, y int) {
	x = space1.X - space2.X
	y = space1.Y - space2.Y
	return
}

func (board *Board) CountAntinodes() (result int) {
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			if len(board.Spaces[x][y].Antinodes) > 0 {
				result++
			}
		}
	}
	return
}

func (board *Board) CountAntinodesPart2() (result int) {
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			if len(board.Spaces[x][y].AntinodesPart2) > 0 {
				result++
			}
		}
	}
	return
}

func (board *Board) GetMatchingTowers(match Space) (matching []Space) {
	for _, space := range board.Towers {
		if space.Tower == match.Tower {
			if space.X != match.X && space.Y != match.Y {
				matching = append(matching, space)
			}
		}
	}

	return
}

func (board *Board) Print() {
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			space := board.Spaces[x][y]
			if len(space.Antinodes) != 0 || len(space.AntinodesPart2) != 0 {
				fmt.Printf("#")
			} else if space.Tower == ' ' {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", space.Tower)
			}
		}
		fmt.Println()
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
