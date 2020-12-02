package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Password struct {
	Value1    int
	Value2    int
	Character string
	Password  string
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
	var passwords []Password
	for _, line := range input_file {
		newPassword := Password{}

		lineParts := strings.Split(line, ":")
		newPassword.Password = strings.TrimSpace(lineParts[1])

		lineParts = strings.Split(lineParts[0], " ")
		newPassword.Character = lineParts[1]

		lineParts = strings.Split(lineParts[0], "-")
		value1, _ := strconv.Atoi(lineParts[0])
		value2, _ := strconv.Atoi(lineParts[1])

		newPassword.Value1 = value1
		newPassword.Value2 = value2

		passwords = append(passwords, newPassword)
	}

	resultPart1 := part1(passwords)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(passwords)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)
}

func part1(passwords []Password) (numValid int) {
	for _, password := range passwords {
		count := strings.Count(password.Password, password.Character)
		if (password.Value1 <= count) && (count <= password.Value2) {
			numValid++
		}
	}

	return numValid
}

func part2(passwords []Password) (numValid int) {
	for _, password := range passwords {
		charPos1 := password.Password[password.Value1-1]
		charPos2 := password.Password[password.Value2-1]

		if (string(charPos1) == password.Character) != (string(charPos2) == password.Character) {
			numValid++
		}
	}

	return numValid
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
