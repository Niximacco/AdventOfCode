package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	Height            int
	Top               bool
	Bottom            bool
	Right             bool
	Left              bool
	ViewingDistTop    int
	viewingDistBottom int
	viewingDistRight  int
	viewingDistLeft   int
	ScenicScore       int
	Row               int
	Col               int
	Forest            *Forest
}

func NewTree(height int, row int, col int) (newTree Tree) {
	newTree = Tree{
		Height:            height,
		Top:               false,
		Bottom:            false,
		Left:              false,
		Right:             false,
		ViewingDistTop:    0,
		viewingDistBottom: 0,
		viewingDistLeft:   0,
		viewingDistRight:  0,
		Row:               row,
		Col:               col,
	}

	return
}

type Forest [][]Tree

func (forest Forest) CalculateVisibilities() {
	topMax := make(map[int]int)
	bottomMax := make(map[int]int)

	for i, _ := range forest[0] {
		topMax[i] = -1
		bottomMax[i] = -1
	}

	for r := 0; r < len(forest); r++ {
		leftMax := -1
		rightMax := -1

		row := forest[r]
		for c := 0; c < len(forest[r]); c++ {
			if row[c].Height > leftMax {
				row[c].Left = true
				leftMax = row[c].Height
			}

			fromRight := len(row) - c - 1
			if row[fromRight].Height > rightMax {
				row[fromRight].Right = true
				rightMax = row[fromRight].Height
			}

			if row[c].Height > topMax[c] {
				row[c].Top = true
				topMax[c] = row[c].Height
			}

			fromBottom := len(forest) - r - 1
			if forest[fromBottom][c].Height > bottomMax[c] {
				forest[fromBottom][c].Bottom = true
				bottomMax[c] = forest[fromBottom][c].Height
			}
		}
	}
}

func (forest Forest) GetRow(rowNum int) (row []Tree) {
	row = forest[rowNum]
	return
}

func (forest Forest) GetCol(colNum int) (col []Tree) {
	col = make([]Tree, 0)
	for row, _ := range forest {
		col = append(col, forest[row][colNum])
	}

	return
}

func (forest Forest) GetBestScenicScore() (best int) {
	best = 0
	for row, _ := range forest {
		for col, _ := range forest[row] {
			forest[row][col].CalculateScenicScore()
			if forest[row][col].ScenicScore > best {
				best = forest[row][col].ScenicScore
			}
		}
	}
	return
}

func (tree *Tree) CalculateScenicScore() {
	treeRow := tree.Forest.GetRow(tree.Row)
	treeCol := tree.Forest.GetCol(tree.Col)

	// Calc Left
	for c := tree.Col - 1; c >= 0; c-- {
		tree.viewingDistLeft++
		if tree.Height <= treeRow[c].Height {
			break
		}
	}

	// Calc Right
	for c := tree.Col + 1; c < len(treeRow); c++ {
		tree.viewingDistRight++
		if tree.Height <= treeRow[c].Height {
			break
		}
	}

	// Calc Top
	for r := tree.Row - 1; r >= 0; r-- {
		tree.ViewingDistTop++
		if tree.Height <= treeCol[r].Height {
			break
		}
	}

	// Calc Bottom
	for r := tree.Row + 1; r < len(treeCol); r++ {
		tree.viewingDistBottom++
		if tree.Height <= treeCol[r].Height {
			break
		}
	}

	tree.ScenicScore = tree.viewingDistLeft * tree.viewingDistRight * tree.ViewingDistTop * tree.viewingDistBottom
}

func (forest Forest) PrintHeights() {
	for row, _ := range forest {
		for col := range forest[row] {
			fmt.Printf("%d", forest[row][col].Height)
		}
		println()
	}
}

func (forest Forest) PrintVisibleTrees() {
	for row, _ := range forest {
		for col := range forest[row] {
			tree := forest[row][col]
			if tree.Top || tree.Bottom || tree.Left || tree.Right {
				fmt.Printf("%d", forest[row][col].Height)
			} else {
				fmt.Printf(" ")
			}
		}
		println()
	}
}

func (forest Forest) SumVisibleTrees() (sum int) {
	for row, _ := range forest {
		for col, _ := range forest[row] {
			tree := forest[row][col]
			if tree.Top || tree.Bottom || tree.Left || tree.Right {
				sum++
			}
		}
	}
	return
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
	forest := make(Forest, 0)
	for row, line := range input_file {
		forest = append(forest, make([]Tree, 0))
		for col, digit := range line {
			height, _ := strconv.Atoi(string(digit))
			newTree := NewTree(height, row, col)
			newTree.Forest = &forest
			forest[row] = append(forest[row], newTree)
		}
	}

	forest.CalculateVisibilities()

	forest.PrintHeights()
	println()
	forest.PrintVisibleTrees()

	part1 = forest.SumVisibleTrees()

	println()
	part2 = forest.GetBestScenicScore()

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
