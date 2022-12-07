package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Element struct {
	Type     string
	Name     string
	Size     int
	Parent   *Element
	Contents Directory
}

type Directory map[string]*Element

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

	root := Element{
		Type:     "D",
		Name:     "/",
		Size:     0,
		Parent:   nil,
		Contents: make(Directory),
	}
	fsPt1 := make(Directory)
	fsPt1["/"] = &root

	current := fsPt1["/"]

	lsOutput := false
	for _, line := range input_file {

		fmt.Printf("%s [Current: %s]\n", line, current.Name)
		if line[0] == '$' {
			lsOutput = false
			commandParts := strings.Fields(line)
			cmd := commandParts[1]

			switch cmd {
			case "cd":
				arg := commandParts[2]
				if arg == ".." {
					fmt.Printf("\tMoving Up one Dir. From %s to ", current.Name)
					current = current.GetParent()
					fmt.Printf("%s\n", current.Name)
				} else {
					fmt.Printf("\tMoving to dir %s from %s\n", arg, current.Name)
					if _, ok := current.Contents[arg]; ok {
						current = current.Contents[arg]
					}
				}
			case "ls":
				fmt.Printf("\tlist found. Contents next\n")
				lsOutput = true
			}
			continue
		}

		if lsOutput {
			lsParts := strings.Fields(line)
			name := lsParts[1]
			fmt.Printf("\tCreating in: %+v\n", current)
			if lsParts[0] == "dir" {
				fmt.Printf("\tNew Directory: '%s'\n", name)
				newDir := &Element{
					Type:     "D",
					Size:     0,
					Name:     name,
					Parent:   nil,
					Contents: make(Directory),
				}

				newDir.Parent = current
				current.Contents[name] = newDir
			} else {
				size, _ := strconv.Atoi(lsParts[0])

				fmt.Printf("\tNew File: '%s' size %d\n", name, size)
				file := &Element{
					Type:     "F",
					Name:     name,
					Size:     size,
					Parent:   nil,
					Contents: nil,
				}
				file.Parent = current
				current.Contents[name] = file
			}
			fmt.Printf("\tdone: %+v\n", current)
		}
	}

	root.Tree(0)
	sizes := root.GetSizes()

	totalDiskSize := 70000000
	spaceNeeded := 30000000
	unused := totalDiskSize - root.RecursiveGetSize()
	toFree := spaceNeeded - unused
	fmt.Printf("unused: %d\nto free: %d\n", unused, toFree)
	sort.Ints(sizes)

	pt1Max := 100000
	part2 = totalDiskSize
	for _, size := range sizes {
		fmt.Printf("%d\n", size)
		if size <= pt1Max {
			part1 += size
		}

		if size >= toFree && size < part2 {
			part2 = size
		}
	}

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func (element Element) GetParent() *Element {
	fmt.Printf("Parent of %s is: %+v", element.Name, element.Parent)
	return element.Parent
}

func (element Element) Tree(depth int) {
	fmt.Printf("%s%s:\n", PadToDepth(depth), element.Name)
	for name, child := range element.Contents {
		switch child.Type {
		case "D":
			child.Tree(depth + 1)
		case "F":
			fmt.Printf("%s%s: %d\n", PadToDepth(depth+1), name, child.Size)
		}
	}
}

func (element Element) GetSizes() (sizes []int) {
	sizes = make([]int, 0)
	fmt.Printf("%s: %d\n", element.Name, element.RecursiveGetSize())
	sizes = append(sizes, element.RecursiveGetSize())
	for _, child := range element.Contents {
		if child.Type == "D" {
			sizes = append(sizes, child.GetSizes()...)
		}
	}
	return
}

func (element Element) RecursiveGetSize() (size int) {
	if element.Type == "F" {
		size = element.Size
	} else {
		for _, child := range element.Contents {
			size += child.RecursiveGetSize()
		}
	}

	return
}

func PadToDepth(depth int) (padding string) {
	for i := 0; i < depth; i++ {
		padding = fmt.Sprintf("  %s", padding)
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
