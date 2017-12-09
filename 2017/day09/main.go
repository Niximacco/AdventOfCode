package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
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
	input := input_file[0]
	check(err)
	// Your Code goes below!

	inGarbage := false
	nextIgnored := false
	numGarbage := 0
	currentScore := 1
	scoreSum := 0
	for i := 0; i<len(input); i++ {
		if string(input[i]) == "{" && !nextIgnored {
			if inGarbage {
				numGarbage++
			} else {
				scoreSum += currentScore
				currentScore++
				nextIgnored = false
			}
		} else if string(input[i]) == "}" && !nextIgnored && !inGarbage {
			if inGarbage {
				numGarbage++
			} else {
				currentScore--
				nextIgnored = false
			}
		} else if string(input[i]) == "<" && !nextIgnored {
			if inGarbage {
				numGarbage++
			}
			inGarbage = true
			nextIgnored = false
		} else if string(input[i]) == ">" && !nextIgnored {
			inGarbage = false
			nextIgnored = false
		} else if string(input[i]) == "!" && !nextIgnored {
			nextIgnored = !nextIgnored
		} else {
			if inGarbage && !nextIgnored {
				numGarbage++
			}
			nextIgnored = false
		}

	}

	fmt.Printf("Part 1 Score: %d\n", scoreSum)
	fmt.Printf("Part 2 Score: %d\n", numGarbage)

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
