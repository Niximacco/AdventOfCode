package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	First  int
	Second int
}

type RuleSet struct {
	RulesForNumber map[int][]Rule
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
	var ruleSet RuleSet
	ruleSet.RulesForNumber = make(map[int][]Rule)
	part1, part2 := 0, 0
	ruleSection := true
	for _, line := range input_file {
		if line == "" {
			ruleSection = false
			continue
		}
		if ruleSection {
			parts := strings.Split(line, "|")
			first, _ := strconv.Atoi(parts[0])
			second, _ := strconv.Atoi(parts[1])
			rule := Rule{first, second}
			ruleSet.AddRuleForNumber(rule)
			continue
		}

		// process pages
		valid, numbers, failedRules := processPages(ruleSet, line)
		fmt.Printf("%v\t%v\t%v\n", valid, numbers, failedRules)
		if valid {
			part1 += numbers[len(numbers)/2]
		} else {
			part2 += fixedNumbersMiddle(numbers, failedRules, ruleSet)
		}
	}

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func (ruleSet *RuleSet) AddRuleForNumber(rule Rule) {
	if _, ok := ruleSet.RulesForNumber[rule.First]; !ok {
		ruleSet.RulesForNumber[rule.First] = []Rule{}
	}

	if _, ok := ruleSet.RulesForNumber[rule.Second]; !ok {
		ruleSet.RulesForNumber[rule.Second] = []Rule{}
	}

	ruleSet.RulesForNumber[rule.First] = append(ruleSet.RulesForNumber[rule.First], rule)
	ruleSet.RulesForNumber[rule.Second] = append(ruleSet.RulesForNumber[rule.Second], rule)
}

func fixedNumbersMiddle(numbers []int, failedRules []Rule, set RuleSet) int {
	fmt.Printf("%d failed rules. Fixing\n", len(failedRules))
	indexForNumber := make(map[int]int)
	for i, number := range numbers {
		indexForNumber[number] = i
	}

	fmt.Printf("%v\t%v\n", numbers, indexForNumber)
	for _, rule := range failedRules {
		fmt.Printf("%v\n", rule)
		// swap positions of numbers
		numbers[indexForNumber[rule.First]], numbers[indexForNumber[rule.Second]] = numbers[indexForNumber[rule.Second]], numbers[indexForNumber[rule.First]]
		indexForNumber[rule.First], indexForNumber[rule.Second] = indexForNumber[rule.Second], indexForNumber[rule.First]
	}
	fmt.Printf("%v\t%v\n", numbers, indexForNumber)

	pageString := make([]string, len(numbers))
	for number, index := range indexForNumber {
		pageString[index] = strconv.Itoa(number)
	}
	fmt.Printf("New Page String: %s\n", pageString)

	valid, newNumbers, newFailedRules := processPages(set, strings.Join(pageString, ","))
	fmt.Printf("%v\t%v\t%v\n", valid, numbers, failedRules)
	if valid {
		return numbers[len(newNumbers)/2]
	} else {
		return fixedNumbersMiddle(newNumbers, newFailedRules, set)
	}
}

func processPages(ruleSet RuleSet, pages string) (valid bool, numbers []int, failedRules []Rule) {
	fmt.Printf("Processing %s\n", pages)
	seenNumbers := map[int]bool{}
	pageNumber := map[int]bool{}
	numberSplit := strings.Split(pages, ",")
	ruleFailures := map[Rule]bool{}
	valid = true
	for _, numberStr := range numberSplit {
		number, _ := strconv.Atoi(numberStr)
		numbers = append(numbers, number)
		pageNumber[number] = true
	}

	for _, number := range numbers {
		// Get rules for number
		rules := ruleSet.RulesForNumber[number]
		for _, rule := range rules {
			if number == rule.First {
				if seenNumbers[rule.Second] {
					// We've seen the second number already, fail.
					fmt.Printf("Rule failed: %d|%d\n", rule.First, rule.Second)
					ruleFailures[rule] = true
					valid = false
				}
			} else if number == rule.Second {
				if pageNumber[rule.First] && !seenNumbers[rule.First] {
					fmt.Printf("Rule failed: %d|%d\n", rule.First, rule.Second)
					ruleFailures[rule] = true
					valid = false
				}
			}
		}
		seenNumbers[number] = true
	}

	for rule, _ := range ruleFailures {
		failedRules = append(failedRules, rule)
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
