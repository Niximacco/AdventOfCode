package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

var replacements map[string][]string
var reverseReplacements map[string]string
var reverseKeys []string
var atoms map[string]bool

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
	replacements = make(map[string][]string)
	reverseReplacements = make(map[string]string)
	lineNum := 0

	for i, line := range input_file {
		lineNum = i
		lineParts := strings.Fields(line)
		if len(lineParts) == 0 {
			break
		}

		if _, ok := replacements[lineParts[0]]; !ok {
			replacements[lineParts[0]] = make([]string, 0)
		}
		replacements[lineParts[0]] = append(replacements[lineParts[0]], lineParts[2])

		reverseReplacements[lineParts[2]] = lineParts[0]

	}

	input := input_file[lineNum+1]

	results := GetResults(input, replacements)
	fmt.Printf("input: %s\n", input)

	part1 = len(results)

	reverseKeys = make([]string, 0, len(reverseReplacements))
	for k := range reverseReplacements {
		reverseKeys = append(reverseKeys, k)
	}
	sort.Slice(reverseKeys, func(i, j int) bool {
		return len(reverseKeys[i]) < len(reverseKeys[j])
	})

	atoms = make(map[string]bool)
	atoms["C"] = true
	atoms["Ca"] = true
	atoms["Si"] = true
	atoms["B"] = true
	atoms["P"] = true
	atoms["Ti"] = true
	atoms["F"] = true
	atoms["Th"] = true
	atoms["Al"] = true
	atoms["Mg"] = true
	atoms["H"] = true
	atoms["N"] = true
	atoms["O"] = true
	atoms["e"] = true
	atoms["Rn"] = false
	atoms["Ar"] = false
	atoms["Y"] = false

	part2 = BuildMoleculeStepsNeeded(input)

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func GetResults(input string, replaceMap map[string][]string) (resultMap map[string]bool) {
	resultMap = make(map[string]bool)
	for replace, _ := range replaceMap {
		indices := IndexAll(input, replace)
		for _, index := range indices {
			before := input[:index]

			after := ""
			if index < len(input)-len(replace) {
				after = input[index+len(replace):]
			}

			for _, replacement := range replaceMap[replace] {
				result := fmt.Sprintf("%s%s%s", before, replacement, after)
				resultMap[result] = true
			}
		}
	}

	return
}

func BuildMoleculeStepsNeeded(input string) (steps int) {
	fmt.Printf("old: %s\n", input)
	input = strings.ReplaceAll(input, "Rn", "(")
	input = strings.ReplaceAll(input, "Ar", ")")
	input = strings.ReplaceAll(input, "Y", ",")
	fmt.Printf("new: %s\n", input)

	for input != "e" {
		for _, key := range reverseKeys {
			createdKey := key
			createdKey = strings.ReplaceAll(createdKey, "Rn", "(")
			createdKey = strings.ReplaceAll(createdKey, "Ar", ")")
			createdKey = strings.ReplaceAll(createdKey, "Y", ",")
			if strings.Contains(input, createdKey) {
				replacement := reverseReplacements[key]
				input = strings.Replace(input, createdKey, replacement, 1)
				steps++
			}
		}
		fmt.Printf("steps: %d\ninput: %s\n", steps, input)
		time.Sleep(100 * time.Millisecond)
	}

	return
}

func GetTokenized(input string) (tokens []string) {
	tokens = make([]string, 0)
	for len(input) > 0 {
		for key, _ := range atoms {
			if strings.HasPrefix(input, key) {
				tokens = append(tokens, key)
				input = strings.TrimPrefix(input, key)
			}
		}
	}
	return
}

func remove[T comparable](l []T, item T) []T {
	out := make([]T, 0)
	for _, element := range l {
		if element != item {
			out = append(out, element)
		}
	}
	return out
}

func IndexAll(s, substr string) (indices []int) {
	// special case: if substr is empty, return every
	// index in s.
	// the default path would cause an infinite loop
	if len(substr) == 0 {
		for i := range s {
			indices = append(indices, i)
		}
		return
	}

	offset := 0
	for {
		i := strings.Index(s[offset:], substr)
		if i == -1 {
			return
		}
		offset += i
		indices = append(indices, offset)
		offset += len(substr)
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
