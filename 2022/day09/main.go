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

type Knot struct {
	X       int
	Y       int
	Child   *Knot
	Parent  *Knot
	Visited map[string]int
	Marker  string
}

func (knot *Knot) GetX() int {
	return knot.X
}

func (knot *Knot) GetY() int {
	return knot.Y
}

func (knot *Knot) SetX(x int) {
	knot.X = x
}

func (knot *Knot) SetY(y int) {
	knot.Y = y
}

func (knot *Knot) GetChild() *Knot {
	return knot.Child
}

func (knot *Knot) SetChild(child *Knot) {
	knot.Child = child
}

func (knot *Knot) GetParent() *Knot {
	return knot.Parent
}

func (knot *Knot) SetParent(child *Knot) {
	knot.Parent = child
}

func (knot *Knot) GetMarker() string {
	return knot.Marker
}

func (knot *Knot) SetMarker(marker string) {
	knot.Marker = marker
}

func (knot *Knot) AddVisit() {
	visitString := fmt.Sprintf("%d,%d", knot.X, knot.Y)
	knot.Visited[visitString] += 1
}

func (knot *Knot) TouchingHead() (touching bool) {
	touching = Abs(knot.Parent.X-knot.X) <= 1 && Abs(knot.Parent.Y-knot.Y) <= 1
	return
}

func (knot *Knot) Move(direction string) {
	switch direction {
	case "U":
		knot.Y++
	case "R":
		knot.X++
	case "D":
		knot.Y--
	case "L":
		knot.X--
	}

	if knot.Child != nil {
		knot.Child.MoveChild()
	}
}

func (knot *Knot) MoveChild() {
	if knot.TouchingHead() {
		return
	}

	xDif := knot.Parent.X - knot.X
	yDif := knot.Parent.Y - knot.Y

	switch {
	case xDif > 0:
		knot.X++
	case xDif < 0:
		knot.X--
	}

	switch {
	case yDif > 0:
		knot.Y++
	case yDif < 0:
		knot.Y--
	}

	knot.AddVisit()
	if knot.Child != nil {
		knot.Child.MoveChild()
	}
}

type Board struct {
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	OffsetX   int
	OffsetY   int
	Positions [][]*Knot
	Matrix    [][]string
}

func NewTailKnot(marker string, parent *Knot) (knot Knot) {
	knot = Knot{
		X:       0,
		Y:       0,
		Child:   nil,
		Parent:  parent,
		Visited: make(map[string]int, 0),
		Marker:  marker,
	}

	parent.SetChild(&knot)

	knot.Visited["0,0"]++
	return
}

func (knot *Knot) PrintRope() {
	fmt.Printf("%s: [%d,%d]\n", knot.Marker, knot.X, knot.Y)
	if knot.Child != nil {
		knot.Child.PrintChild(1)
	}
}

func (knot *Knot) PrintChild(depth int) {
	padding := ">"
	for i := 0; i < depth; i++ {
		padding = fmt.Sprintf("%s%s", "-", padding)
	}
	fmt.Printf("%s%s: [%d,%d]\n", padding, knot.Marker, knot.X, knot.Y)
	if knot.Child != nil {
		knot.Child.PrintChild(depth + 1)
	}
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
	pt1Head := Knot{
		X:      0,
		Y:      0,
		Marker: "H",
	}
	pt1Tail := NewTailKnot("T", &pt1Head)
	pt1Rope := []*Knot{&pt1Head, &pt1Tail}
	boardPt1 := Board{
		MinX:      0,
		MaxX:      0,
		MinY:      0,
		MaxY:      0,
		Positions: make([][]*Knot, 0),
	}
	boardPt1.AddPositions(pt1Rope)

	pt2Rope := make([]*Knot, 0)
	pt2Head := Knot{
		X:      0,
		Y:      0,
		Marker: "H",
	}
	pt2Rope = append(pt2Rope, &pt2Head)

	for i := 1; i <= 9; i++ {
		tail := NewTailKnot(fmt.Sprintf("%d", i), pt2Rope[i-1])
		pt2Rope = append(pt2Rope, &tail)
		pt2Rope[i-1].SetChild(&tail)
	}

	boardPt2 := Board{
		MinX:      0,
		MaxX:      0,
		MinY:      0,
		MaxY:      0,
		Positions: make([][]*Knot, 0),
	}
	boardPt1.AddPositions(pt2Rope)

	for _, line := range input_file {
		lineParts := strings.Fields(line)
		dir := lineParts[0]
		amt, _ := strconv.Atoi(lineParts[1])
		fmt.Printf("Moving %s by %d\n", dir, amt)
		for count := 0; count < amt; count++ {
			pt1Head.Move(dir)
			pt2Head.Move(dir)
			pt2Head.PrintRope()

			boardPt1.AddPositions(pt1Rope)
			boardPt2.AddPositions(pt2Rope)
		}
	}

	boardPt1.InitMatrix()
	//board.Animate(1)
	boardPt1.AddVisitedTails(pt1Rope)
	boardPt1.Print()
	part1 = boardPt1.CountVisited()

	boardPt2.InitMatrix()
	boardPt2.AddVisitedTails(pt2Rope)
	boardPt2.Print()
	part2 = boardPt2.CountVisited()

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func (board *Board) AddPositions(knots []*Knot) {
	board.Positions = append(board.Positions, knots)
	for _, knot := range knots {
		board.UpdateMinMaxVals(*knot)
	}
}

func (board *Board) UpdateMinMaxVals(knot Knot) {
	if knot.X > board.MaxX {
		board.MaxX = knot.X
	}

	if knot.X < board.MinX {
		board.MinX = knot.X
	}

	if knot.Y > board.MaxY {
		board.MaxY = knot.Y
	}

	if knot.Y < board.MinY {
		board.MinY = knot.Y
	}
}

func (board *Board) InitMatrix() {
	fmt.Printf("Making Board. X: [(%d)-(%d)]\tY: [(%d)-(%d)]\n", board.MinX, board.MaxX, board.MinY, board.MaxY)
	width := board.MaxX - board.MinX + 1
	height := board.MaxY - board.MinY + 1
	fmt.Printf("Board Dimensions: %dx%d\n", width, height)
	board.Matrix = make([][]string, 0)
	board.OffsetX = board.MinX
	board.OffsetY = board.MinY
	for row := 0; row < height; row++ {
		board.Matrix = append(board.Matrix, make([]string, 0))
		for col := 0; col < width; col++ {
			board.Matrix[row] = append(board.Matrix[row], ".")
		}
	}

	board.SetMatrixPoint(0, 0, "S")
	board.Print()
}

func (board *Board) SetMatrixPoint(x int, y int, Value string) {
	board.Matrix[y-board.OffsetY][x-board.OffsetX] = Value
}

func (board *Board) Print() {
	fmt.Printf("%dx%d\n", len(board.Matrix[0]), len(board.Matrix))
	for row := len(board.Matrix) - 1; row >= 0; row-- {
		for col := 0; col < len(board.Matrix[row]); col++ {
			fmt.Printf("%s", board.Matrix[row][col])
		}
		println()
	}
	println()
}

func (board *Board) AddVisitedTails(knots []*Knot) {
	knot := knots[len(knots)-1]
	for position, _ := range knot.Visited {
		parts := strings.Split(position, ",")
		X, _ := strconv.Atoi(parts[0])
		Y, _ := strconv.Atoi(parts[1])
		board.DrawVisit(X, Y)
	}

	board.SetMatrixPoint(0, 0, "S")
}

func (board *Board) DrawVisit(x int, y int) {
	board.SetMatrixPoint(x, y, "#")
}

func (board *Board) Clear(knots []*Knot) {
	for _, knot := range knots {
		board.SetMatrixPoint(knot.X, knot.Y, knot.Marker)
	}
	board.SetMatrixPoint(0, 0, "S")
}

func (board *Board) Draw(knots []*Knot) {
	board.SetMatrixPoint(0, 0, "S")
	for i := len(knots) - 1; i >= 0; i++ {
		board.SetMatrixPoint(knots[i].X, knots[i].Y, knots[i].Marker)
	}
}

func (board *Board) Animate(sleepDur int) {
	board.InitMatrix()
	previousPosition := make([]*Knot, 0)

	for _, pos := range board.Positions {
		board.Clear(previousPosition)
		board.Draw(pos)
		previousPosition = pos
		board.Print()
		time.Sleep(time.Duration(sleepDur) * time.Second)
	}
}

func (board *Board) CountVisited() (visits int) {
	for row := 0; row < len(board.Matrix); row++ {
		for col := 0; col < len(board.Matrix[row]); col++ {
			if board.Matrix[row][col] == "#" || board.Matrix[row][col] == "S" {
				visits++
			}
		}
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
