package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	ContainedBy []string
	Contains    map[string]int
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

	bagInfo := parseRules(input_file)

	resultPart1 := part1(bagInfo)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(bagInfo)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(bagContainsInfo map[string]Bag) (numColorsContainingGold int) {
	bagToBagsContaining := make(map[string][]string)
	for key, value := range bagContainsInfo {
		for bagContainedBy, _ := range value.Contains {
			bagToBagsContaining[bagContainedBy] = append(bagToBagsContaining[bagContainedBy], key)
		}
	}

	numColorsContainingGold = len(recursiveSearchBagsEventuallyContaining(bagToBagsContaining, "shiny gold"))

	return
}

func part2(bagContainsInfo map[string]Bag) (numBags int) {
	numBags = recursiveSearchHowManyBagsContained(bagContainsInfo, "shiny gold")

	return
}

func recursiveSearchHowManyBagsContained(bagContainsInfo map[string]Bag, color string) (count int) {
	fmt.Printf("checking %s\n", color)
	for bagType, num := range bagContainsInfo[color].Contains {
		count += num
		count += num * recursiveSearchHowManyBagsContained(bagContainsInfo, bagType)
		fmt.Printf("\t%s: %d (%d)\n", bagType, num, count)
	}

	return
}

func recursiveSearchBagsEventuallyContaining(bagToBagsContaining map[string][]string, color string) (items map[string]bool) {
	items = make(map[string]bool)
	fmt.Printf("searching %s\n", color)
	if _, present := bagToBagsContaining[color]; present {
		for _, containingColor := range bagToBagsContaining[color] {
			items[containingColor] = true
			for item, _ := range recursiveSearchBagsEventuallyContaining(bagToBagsContaining, containingColor) {
				items[item] = true
			}
		}
	}

	return
}

func parseRules(lines []string) (bags map[string]Bag) {
	bags = make(map[string]Bag)
	for _, line := range lines {
		newBag := Bag{}
		newBag.Contains = make(map[string]int)

		reLineMatch := regexp.MustCompile("^(.+) bags contain (.+)$")
		lineMatches := reLineMatch.FindStringSubmatch(line)
		containsList := strings.Split(lineMatches[2], ",")
		for _, item := range containsList {
			if strings.Contains(item, "no other bags") {
				continue
			}
			reContainsMatch := regexp.MustCompile("(\\d+) (.+) bag")
			containsMatches := reContainsMatch.FindStringSubmatch(item)

			count, _ := strconv.Atoi(containsMatches[1])
			newBag.Contains[containsMatches[2]] = count
		}

		if lineMatches != nil {
			bags[lineMatches[1]] = newBag
		}
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
