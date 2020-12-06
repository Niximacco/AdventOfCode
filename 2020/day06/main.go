package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Group struct {
	Members          []string
	NumUniqueAnswers int
	NumCommonAnswers int
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
	fileContents, err := readContent(filename)
	check(err)

	// Your Code goes below!
	groups := parseGroups(fileContents)
	resultPart1 := part1(groups)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(groups)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(groups []Group) (sumUniqueAnswers int) {
	for _, group := range groups {
		group.sumUniqueMemberAnswers()
		sumUniqueAnswers += group.NumUniqueAnswers
	}
	return
}

func part2(groups []Group) (sumCommonAnswers int) {
	for _, group := range groups {
		group.sumCommonAnswers()
		sumCommonAnswers += group.NumCommonAnswers
	}
	return
}

func (group *Group) sumUniqueMemberAnswers() {
	uniqueAnswers := make(map[string]bool)

	for _, member := range group.Members {
		for _, char := range member {
			uniqueAnswers[string(char)] = true
		}
	}
	group.NumUniqueAnswers = len(uniqueAnswers)
}

func (group *Group) sumCommonAnswers() {
	answersCount := make(map[string]int)
	commonAnswers := 0

	for _, member := range group.Members {
		fmt.Printf("%s\n", member)
		for _, rune := range member {
			char := string(rune)
			if _, present := answersCount[char]; present {
				answersCount[char]++
			} else {
				answersCount[char] = 1
			}
		}
	}

	for key, value := range answersCount {
		fmt.Printf("\tChecking %s  (%d == %d)", key, value, len(group.Members))
		if value == len(group.Members) {
			fmt.Printf(": true\n")
			commonAnswers++
		} else {
			fmt.Printf(": false\n")
		}
	}
	fmt.Printf("\t\tcommon: %d\n", commonAnswers)
	group.NumCommonAnswers = commonAnswers
}

func parseGroups(input string) (groups []Group) {
	for _, groupString := range strings.Split(input, "\n\n") {
		newGroup := Group{}
		newGroup.Members = strings.Split(groupString, "\n")
		groups = append(groups, newGroup)
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readContent(path string) (content string, err error) {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	byteContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	content = string(byteContent)
	return
}
