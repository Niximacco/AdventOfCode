package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	part1, part2 := 0, 0
	isFile := true
	fileId := 0
	var contents []string
	for _, line := range input_file {

		for _, char := range line {
			number, _ := strconv.Atoi(string(char))
			for i := 0; i < number; i++ {
				if isFile {
					contents = append(contents, fmt.Sprintf("%d", fileId))
				} else {
					contents = append(contents, ".")
				}
			}

			isFile = !isFile
			if isFile {
				fileId++
			}
		}
	}

	contentsCopy := make([]string, len(contents))

	copy(contentsCopy, contents)

	fmt.Printf("%s\n", strings.Join(contents, ""))
	defrag(contents)
	fmt.Printf("%s\n", strings.Join(contents, ""))

	part1 = checksum(contents)
	println("=========")

	filesystem := make(map[string]*File)
	filesystem = parseFiles(contentsCopy)
	part2 = defragPart2(filesystem, contentsCopy)

	fmt.Printf("%+v\n", filesystem)

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func parseFiles(contents []string) (files map[string]*File) {
	files = make(map[string]*File)
	for i := 0; i < len(contents); i++ {
		if contents[i] != "." {
			fileId, _ := strconv.Atoi(contents[i])

			length := 0
			fileEnd := i
			for j := i; j < len(contents); j++ {
				if contents[j] != contents[i] {
					break
				}
				length++
				fileEnd = j
			}

			file := File{
				ID:         fileId,
				StartIndex: i,
				Length:     length,
			}

			i = fileEnd

			files[contents[i]] = &file
		}
	}
	return files
}

func defrag(contents []string) {
	nextEmpty := getNextEmptyIndex(contents)
	for i := len(contents) - 1; i >= 0; i-- {
		if contents[i] != "." {
			contents[nextEmpty] = contents[i]
			contents[i] = "."
			nextEmpty = getNextEmptyIndex(contents)
			if nextEmpty >= i {
				break
			}
		}
		//fmt.Printf("%s\n", strings.Join(contents, ""))
	}
}

func defragPart2(filesystem map[string]*File, contents []string) (check int) {
	for i := len(contents) - 1; i >= 0; i-- {
		//padding := ""
		//for j := 0; j < i; j++ {
		//	padding += " "
		//}
		//padding += "v"
		//fmt.Printf("%s%d\n%s\n=======\n", padding, i, strings.Join(contents, ""))
		if contents[i] != "." {
			file := filesystem[contents[i]]
			emptyStart := firstEmptyWithLen(contents, file.Length)
			if emptyStart >= 0 && emptyStart < file.StartIndex {
				//fmt.Printf("\tMoving File: %+v to", file)
				fileIdString := strconv.Itoa(file.ID)
				for j := emptyStart; j < emptyStart+file.Length; j++ {
					contents[j] = fileIdString
				}

				for j := file.StartIndex; j < file.StartIndex+file.Length; j++ {
					contents[j] = "."
				}

				file.StartIndex = emptyStart
				//fmt.Printf("%+v\n", file)
				filesystem[contents[i]] = file
			}
		}
	}

	check = checksum(contents)
	return check
}

func firstEmptyWithLen(contents []string, length int) (emptyStart int) {
	emptyStart = -1
	emptyCount := 0
	inEmpty := false
	for i := 1; i < len(contents)-1; i++ {
		if contents[i] == "." {
			emptyCount++
			if !inEmpty {
				emptyStart = i
				inEmpty = true
			}
			if emptyCount == length {
				return emptyStart
			}
		} else {
			inEmpty = false
			emptyCount = 0
		}
	}
	return -1
}

type Empty struct {
	StartIndex int
	Length     int
}

type File struct {
	ID         int
	StartIndex int
	Length     int
}

func getEmptiesBeforeIndex(contents []string, index int) int {
	return 0
}

func checksum(contents []string) (check int) {
	fmt.Printf("Calculating Checksum for %s: ", strings.Join(contents, ""))
	for i, char := range contents {
		value, _ := strconv.Atoi(char)
		check += i * value
	}
	fmt.Printf("%d\n", check)
	return
}

func getNextEmptyIndex(contents []string) int {
	for i := 0; i < len(contents); i++ {
		if contents[i] == "." {
			return i
		}
	}
	return -1
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
