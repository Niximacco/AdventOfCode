package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
	"strings"
	"strconv"
)

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
	numUsed := 0
	hashStringBase := strings.TrimSpace(input_file[0])

	lines := make([][]string, 128)
	for i := range lines {
		lines[i] = make([]string, 128)
	}

	for i:=0; i<128; i++ {
		hashString := fmt.Sprintf("%s-%d", hashStringBase, i)
		hashed := knotHash(hashString)
		binaryString := ""
		for _, character := range []byte(hashed) {
			value := getValueFromHex(string(character))
			binValue := fmt.Sprintf("%04b", value)
			binaryString = fmt.Sprintf("%s%s", binaryString, binValue)
		}
		binaryArr := []byte(binaryString)


		for j, num := range binaryArr {
			if num == '1' {
				numUsed++
				fmt.Printf("#")
				lines[i][j]="#"
			} else {
				fmt.Printf(".")
				lines[i][j]="."
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Part 1 Solution: %d\n\n", numUsed)

	numGroups := calculateGroups(lines)

	fmt.Printf("Part 2 Solution: %d\n",numGroups)

}

func calculateGroups(lines [][]string) int {
	currentGroupNum := 1
	changed := false
	for row:=0; row<128; row++ {
		for column:=0; column<128; column++ {
			currentGroup := fmt.Sprintf("%d", currentGroupNum)
			lines, changed = calculatePoint(row, column, currentGroup, lines)
			if changed {
				currentGroupNum++
			}
		}
	}
	for row:=0; row<128; row++ {
		for col:=0; col<128; col++ {
			fmt.Printf("%s", lines[row][col])
		}
		fmt.Printf("\n")
	}
	return currentGroupNum - 1
}

func calculatePoint(row int, column int, currentGroup string, lines [][]string) ([][]string, bool) {
	changed := false
	if lines[row][column] == "#" {
		//Not in a group yet.
		lines[row][column] = currentGroup
		changed = true
	}

	if changed {
		if row > 0 {
			lines, _ = calculatePoint(row-1, column, currentGroup, lines)
		}

		if row < 127 {
			lines, _ = calculatePoint(row+1, column, currentGroup, lines)
		}

		if column > 0 {
			lines, _ = calculatePoint(row, column-1, currentGroup, lines)
		}

		if column < 127 {
			lines, _ = calculatePoint(row, column+1, currentGroup, lines)
		}
	}


	return lines, changed
}

func getValueFromHex(input string) int {
	switch input {
	case "a":
		return 10
	case "b":
		return 11
	case "c":
		return 12
	case "d":
		return 13
	case "e":
		return 14
	case "f":
		return 15
	default:
		val, err := strconv.Atoi(input)
		check(err)

		return val
	}
}

func knotHash(input string) string {
	unicodeCodePoints := []byte(input)
	unicodeCodePoints = append(unicodeCodePoints, 17, 31, 73, 47, 23)


	currentIndex := 0
	jumpSize := 0
	var values []int
	for i := 0; i < 256; i++ {
		values = append(values, i)
	}


	for i:=0; i<64; i++ {
		for _, num := range unicodeCodePoints {
			//Select the 'num' amount of numbers from current index
			var reverseArr []int
			for i := 0; i < int(num); i++ {
				index := currentIndex + i
				for index >= len(values) {
					index -= len(values)
				}
				reverseArr = append(reverseArr, values[index])
			}
			//fmt.Printf("\t%v\n", reverseArr)
			//reverse them in the values array
			for j := 0; j < int(num); j++ {
				index := currentIndex + j
				for index >= len(values) {
					index -= len(values)
				}
				values[index] = reverseArr[int(num)-j-1]
				//fmt.Printf("\t%d\t%v\n", index, values)

			}

			//increment currentIndex 'num' + jumpSize
			currentIndex = (currentIndex + int(num) + jumpSize)
			if currentIndex >= len(values) {
				currentIndex -= len(values)
			}

			//increment jumpSize
			jumpSize++
		}
	}
	var finalString string
	for i:=0; i<16; i++ {
		character := values[i*16]
		for j:=1; j<16; j++ {
			character = character ^ values[(i*16)+j]
		}
		hexChars := fmt.Sprintf("%x", character)
		if len(hexChars) < 2 {
			hexChars = "0" + hexChars
		}
		finalString = fmt.Sprintf("%s%s", finalString, hexChars)
		//fmt.Printf("Char %d: int: %d\thex: %s\n", i, character, hexChars)
	}
	return fmt.Sprintf("%s", finalString)
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
