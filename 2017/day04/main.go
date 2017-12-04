package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"sort"
)

func main() {
	input_file, err := readLines("input.txt")
	check(err)

	valid1 := true
	valid2 := true
	numValid1 := 0
	numValid2 := 0
	for _, line := range input_file {
		words := strings.Fields(line)
		for i, word := range words {
			for j, checkWord := range words {
				if i != j {
					if areWordsAnagrams(word, checkWord) {
						valid2 = false
					}
					if word == checkWord {
						valid1 = false
					}
				}
			}
		}
		if valid1 {
			numValid1++
		}
		if valid2 {
			numValid2++
		}
		valid1 = true
		valid2 = true
	}

	fmt.Printf("Part 1: %d\n", numValid1)
	fmt.Printf("Part 2: %d", numValid2)
}

func areWordsAnagrams(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	sorted1 := SortString(word1)
	sorted2 := SortString(word2)
	return sorted1 == sorted2
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

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}