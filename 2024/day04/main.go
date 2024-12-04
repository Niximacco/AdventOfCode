package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Board struct {
	Cells      [][]Cell
	Rows, Cols int
	Part1      int
	Part2      int
}

type Cell struct {
	Value  byte
	Active bool
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
	board := Board{
		Cells: make([][]Cell, len(input_file[0])),
	}
	for r, row := range input_file {
		board.Cells = append(board.Cells, make([]Cell, len(input_file)))
		board.Rows++
		board.Cols = len(row)
		for _, col := range row {
			board.Cells[r] = append(board.Cells[r], Cell{
				Value:  byte(col),
				Active: false,
			})
		}
	}

	board.PrintBoard()

	for r := 0; r < board.Rows; r++ {
		for c := 0; c < board.Cols; c++ {
			if board.Cells[r][c].Value == 'X' {
				board.checkPositionPart1(r, c)
			}
		}
	}

	board.PrintActive()
	part1 = board.Part1

	board.ResetActive()

	for r := 0; r < board.Rows; r++ {
		for c := 0; c < board.Cols; c++ {
			if board.Cells[r][c].Value == 'A' {
				board.checkPositionPart2(r, c)
			}
		}
	}

	board.PrintActive()
	part2 = board.Part2

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func (board *Board) PrintBoard() {
	fmt.Println("Board:")
	for r := 0; r < board.Rows; r++ {
		for c := 0; c < board.Cols; c++ {
			fmt.Print(string(board.Cells[r][c].Value))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (board *Board) PrintActive() {
	fmt.Println("Board Matches:")
	for r := 0; r < board.Rows; r++ {
		for c := 0; c < board.Cols; c++ {
			if board.Cells[r][c].Active {
				fmt.Print(string(board.Cells[r][c].Value))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (board *Board) checkPositionPart1(r, c int) {
	// if horizontal
	board.checkHorizontalPart1(r, c)
	// if vertical
	board.checkVerticalPart1(r, c)
	// if diagonal
	board.checkDiagonalPart1(r, c)

}

func (board *Board) checkHorizontalPart1(r, c int) {
	// check horizontal forward
	if board.GetCellValue(r, c+1) == 'M' && board.GetCellValue(r, c+2) == 'A' && board.GetCellValue(r, c+3) == 'S' {
		board.Part1++
		board.Cells[r][c].Active = true
		board.Cells[r][c+1].Active = true
		board.Cells[r][c+2].Active = true
		board.Cells[r][c+3].Active = true
	}

	// check horizontal backwards
	if board.GetCellValue(r, c-1) == 'M' && board.GetCellValue(r, c-2) == 'A' && board.GetCellValue(r, c-3) == 'S' {
		board.Part1++
		board.Cells[r][c].Active = true
		board.Cells[r][c-1].Active = true
		board.Cells[r][c-2].Active = true
		board.Cells[r][c-3].Active = true
	}
}

func (board *Board) checkVerticalPart1(r, c int) {
	// check vertical down
	if board.GetCellValue(r+1, c) == 'M' && board.GetCellValue(r+2, c) == 'A' && board.GetCellValue(r+3, c) == 'S' {
		board.Part1++
		board.Cells[r][c].Active = true
		board.Cells[r+1][c].Active = true
		board.Cells[r+2][c].Active = true
		board.Cells[r+3][c].Active = true
	}

	// check vertical up
	if board.GetCellValue(r-1, c) == 'M' && board.GetCellValue(r-2, c) == 'A' && board.GetCellValue(r-3, c) == 'S' {
		board.Part1++
		board.Cells[r][c].Active = true
		board.Cells[r-1][c].Active = true
		board.Cells[r-2][c].Active = true
		board.Cells[r-3][c].Active = true
	}
}

func (board *Board) checkDiagonalPart1(r, c int) {
	// check down right
	if board.GetCellValue(r+1, c+1) == 'M' && board.GetCellValue(r+2, c+2) == 'A' && board.GetCellValue(r+3, c+3) == 'S' {
		board.Part1++
		board.Cells[r][c].Active = true
		board.Cells[r+1][c+1].Active = true
		board.Cells[r+2][c+2].Active = true
		board.Cells[r+3][c+3].Active = true
	}

	// check down left
	if board.GetCellValue(r+1, c-1) == 'M' && board.GetCellValue(r+2, c-2) == 'A' && board.GetCellValue(r+3, c-3) == 'S' {
		board.Part1++
		board.Cells[r][c].Active = true
		board.Cells[r+1][c-1].Active = true
		board.Cells[r+2][c-2].Active = true
		board.Cells[r+3][c-3].Active = true
	}

	// check up right
	if board.GetCellValue(r-1, c+1) == 'M' && board.GetCellValue(r-2, c+2) == 'A' && board.GetCellValue(r-3, c+3) == 'S' {
		board.Part1++
		board.Cells[r][c].Active = true
		board.Cells[r-1][c+1].Active = true
		board.Cells[r-2][c+2].Active = true
		board.Cells[r-3][c+3].Active = true
	}

	// check up left
	if board.GetCellValue(r-1, c-1) == 'M' && board.GetCellValue(r-2, c-2) == 'A' && board.GetCellValue(r-3, c-3) == 'S' {
		board.Part1++
		board.Cells[r][c].Active = true
		board.Cells[r-1][c-1].Active = true
		board.Cells[r-2][c-2].Active = true
		board.Cells[r-3][c-3].Active = true
	}
}

func (board *Board) checkPositionPart2(r, c int) {
	dmatchDR, dmatchUR := false, false

	downRight := board.getSlice([]int{r - 1, r, r + 1}, []int{c - 1, c, c + 1})
	fmt.Printf("[%d,%d] downRight=%s\n", r, c, string(downRight))
	if string(downRight) == "SAM" || string(downRight) == "MAS" {
		dmatchDR = true
	}

	upRight := board.getSlice([]int{r + 1, r, r - 1}, []int{c - 1, c, c + 1})
	fmt.Printf("[%d,%d] upRight=%s\n", r, c, string(upRight))
	if string(upRight) == "SAM" || string(upRight) == "MAS" {
		dmatchUR = true
	}

	if dmatchDR && dmatchUR {
		board.Part2++

		board.activateSlice([]int{r - 1, r, r + 1}, []int{c - 1, c, c + 1})
		board.activateSlice([]int{r + 1, r, r - 1}, []int{c - 1, c, c + 1})
	}

}

func (board *Board) activateSlice(rows []int, cols []int) {
	for i := 0; i < len(rows); i++ {
		board.Cells[rows[i]][cols[i]].Active = true
	}
}

func (board *Board) getSlice(rows []int, cols []int) (bytes []byte) {
	for i := 0; i < len(rows); i++ {
		bytes = append(bytes, byte(board.GetCellValue(rows[i], cols[i])))
	}
	return
}

func (board *Board) GetCellValue(r, c int) rune {
	if r >= 0 && r < board.Rows && c >= 0 && c < board.Cols {
		return rune(board.Cells[r][c].Value)
	}

	return ' '
}

func (board *Board) ResetActive() {
	for r := 0; r < board.Rows; r++ {
		for c := 0; c < board.Cols; c++ {
			board.Cells[r][c].Active = false
		}
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
