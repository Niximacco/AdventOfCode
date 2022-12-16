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

type Status int

const (
	Incomplete Status = iota
	Correct
	Incorrect
)

type Type int

const (
	Packet Type = iota
	Number
	List
)

type Element struct {
	Type     Type
	Value    int
	Contents []*Element
}

func (numElement *Element) NumToList() (listElement *Element) {
	if numElement.Type != Number {
		return
	}

	listElement = &Element{
		Type:     List,
		Value:    0,
		Contents: []*Element{numElement},
	}
	return
}

func (element *Element) ToString() (elementString string) {
	switch element.Type {
	case Packet:
		elementString = fmt.Sprintf("Packet: [")
	case List:
		elementString = fmt.Sprintf("[")
	case Number:
		elementString = fmt.Sprintf("%d", element.Value)
	}

	concatStrings := make([]string, 0)
	for _, child := range element.Contents {
		concatStrings = append(concatStrings, child.ToString())
	}

	elementString = fmt.Sprintf("%s%s", elementString, strings.Join(concatStrings, ","))

	switch element.Type {
	case Packet:
		elementString = fmt.Sprintf("%s]", elementString)
	case List:
		elementString = fmt.Sprintf("%s]", elementString)
	}

	return
}

func (left *Element) CompareTo(right *Element) (status Status) {
	fmt.Printf("Comparing %s to %s\n", left.ToString(), right.ToString())
	leftIndex := 0
	rightIndex := 0
	for leftIndex < len(left.Contents) && rightIndex < len(right.Contents) {

		leftChild := left.Contents[leftIndex]
		rightChild := right.Contents[rightIndex]
		fmt.Printf("\tComparing %s to %s\n", leftChild.ToString(), rightChild.ToString())
		if leftChild.Type == Number && rightChild.Type == Number {
			if leftChild.Value > rightChild.Value {
				status = Incorrect
			}

			if leftChild.Value < rightChild.Value {
				status = Correct
			}
		}

		if leftChild.Type == List && rightChild.Type == List {
			status = leftChild.CompareTo(rightChild)
		}

		if leftChild.Type == Number && rightChild.Type == List {
			status = leftChild.NumToList().CompareTo(rightChild)
		}

		if leftChild.Type == List && rightChild.Type == Number {
			status = leftChild.CompareTo(rightChild.NumToList())
		}

		if status != Incomplete {
			return
		}

		leftIndex++
		rightIndex++
	}

	if leftIndex >= len(left.Contents) && rightIndex < len(right.Contents) {
		status = Correct
	}

	if leftIndex < len(left.Contents) && rightIndex >= len(right.Contents) {
		status = Incorrect
	}

	return
}

func (element *Element) Parse(packet string) {
	element.Value = 0
	element.Contents = make([]*Element, 0)
	packet = strings.TrimPrefix(packet, "[")
	packet = strings.TrimSuffix(packet, "]")
	packet = strings.ReplaceAll(packet, "[", "[,")
	packet = strings.ReplaceAll(packet, "]", ",]")
	elements := strings.Split(packet, ",")
	for i := 0; i < len(elements); i++ {
		testElement := elements[i]
		endElement := i + 1
		startElement := 0
		if testElement == "[" {
			numOpen := 1
			childElement := Element{
				Type:     List,
				Value:    0,
				Contents: make([]*Element, 0),
			}
			startElement = i + 1
			for inner := startElement; inner < len(elements); inner++ {
				if elements[inner] == "[" {
					numOpen++
				}

				if elements[inner] == "]" {
					numOpen--
					if numOpen == 0 {
						endElement = inner
						break
					}
				}
			}
			childElement.Parse(fmt.Sprintf("[%s]", strings.Join(elements[startElement:endElement], ",")))
			element.Contents = append(element.Contents, &childElement)
			i = endElement
		} else {
			childElement := Element{
				Type:     Number,
				Value:    0,
				Contents: nil,
			}
			if testElement != "" {
				childElement.Value, _ = strconv.Atoi(testElement)
				element.Contents = append(element.Contents, &childElement)
			}
		}
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
	pairIndex := 1
	packets := make([]*Element, 0)
	for i := 0; i < len(input_file); i += 3 {
		packet1 := &Element{
			Type:     Packet,
			Value:    0,
			Contents: make([]*Element, 0),
		}
		packet1.Parse(input_file[i])
		packets = append(packets, packet1)

		packet2 := &Element{
			Type:     Packet,
			Value:    0,
			Contents: make([]*Element, 0),
		}
		packet2.Parse(input_file[i+1])
		packets = append(packets, packet2)

		fmt.Printf("%s <?> %s\n", input_file[i], input_file[i+1])

		status := packet1.CompareTo(packet2)
		switch status {
		case Incomplete:
			fmt.Printf("Incomplete information.\n")
		case Incorrect:
			fmt.Printf("Incorrect!\n")
		case Correct:
			part1 += pairIndex
			fmt.Printf("Correct!\n")
		}

		println("-------")
		pairIndex++
	}

	packetMarker1 := &Element{
		Type:     Packet,
		Value:    0,
		Contents: make([]*Element, 0),
	}
	packetMarker1.Parse("[[2]]")

	packetMarker2 := &Element{
		Type:     Packet,
		Value:    0,
		Contents: make([]*Element, 0),
	}
	packetMarker2.Parse("[[6]]")
	packets = append(packets, packetMarker1)
	packets = append(packets, packetMarker2)

	sort.Slice(packets, func(a, b int) bool {
		return packets[a].CompareTo(packets[b]) == Correct
	})

	index1 := 0
	index2 := 0
	for i, element := range packets {
		fmt.Printf("%s\n", element.ToString())
		if element == packetMarker1 {
			index1 = i + 1
		}

		if element == packetMarker2 {
			index2 = i + 1
		}
	}

	part2 = index1 * index2

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
