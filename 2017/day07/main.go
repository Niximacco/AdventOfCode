package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
	"regexp"
	"strings"
	//"strconv" "time"
	"time"
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
	part1(input_file)
}

func part1(input_file []string) {
	lookup := make(map[string][]string)

	for _, line := range input_file {
		line = strings.TrimSpace(line)
		re := regexp.MustCompile("(.+?) \\((.+?)\\)( -> (.+))?")
		matched := re.FindStringSubmatch(line)

		var supporting []string
		name := matched[1]
		//size, err := strconv.Atoi(matched[2])
		supporting = strings.Split(matched[4], ",")

		if len(supporting) > 1 {
			for _, supportName := range supporting {
				supportName = strings.TrimSpace(supportName)
				lookup[name] = append(lookup[name], supportName)
			}
		}

		if len(lookup[name]) == 0 {
			delete(lookup, name)
		}



		//fmt.Printf("Name: %s\tsize: %d\tSupporting: %s\tNumSupport: %d\n", name, size, supporting, len(supporting))
		//fmt.Println("=====")

	}
	for len(lookup) > 1 {
		fmt.Println(lookup)
		fmt.Println(len(lookup))
		time.Sleep(1*time.Second)
		for k := range lookup {
			//fmt.Println(k)
			for _, child := range lookup[k] {
				//fmt.Println("\t'" + child + "'")
				reduce(lookup, child)
				delete(lookup, child)
			}
		}
	}

	fmt.Println(lookup)
	for k := range lookup {
		part2(input_file, k)
	}
}

func part2(input_file []string, root string) {
	fmt.Println("======\npart2\n=====")
	weights := make(map[string]int)
	parents := make(map[string]string)
	children := make(map[string][]string)
	for _, line := range input_file {
		line = strings.TrimSpace(line)
		re := regexp.MustCompile("(.+?) \\((.+?)\\)( -> (.+))?")
		matched := re.FindStringSubmatch(line)

		name := matched[1]
		size, err := strconv.Atoi(matched[2])
		check(err)

		supporting := strings.Split(matched[4], ",")
		if len(supporting) > 1 {
			for _, supportName := range supporting {
				supportName = strings.TrimSpace(supportName)
				parents[supportName] = name
				children[name] = append(children[name], supportName)
			}
		}

		weights[name] = size
	}
	var childrenWeights []int
	for _, k := range children[root] {
		childrenWeights = append(childrenWeights, sumWeightOfNode(k, children, weights))
	}
	fmt.Printf("%s: (%d) %s = %d\n", root, weights[root], children[root], sumWeightOfNode(root, children, weights))
	fmt.Println(childrenWeights)
	problem := findProblem(root, children, weights)
	shouldEqual := sumWeightOfNode(children[root][findNonOutlierIndex(childrenWeights)], children, weights)

	fmt.Printf("Problem in: %s. ShouldEqual %d\n", problem, shouldEqual)
	fmt.Println(findProblemInNode(problem, shouldEqual, children, weights))
}

func sumWeightOfNode(name string, children map[string][]string, weights map[string]int) int {
	weight := 0
	weight += weights[name]

	for _, key := range children[name] {
		weight += sumWeightOfNode(key, children, weights)
	}

	return weight
}

func findProblem(name string, children map[string][]string, weights map[string]int) string {
	var childrenWeights []int
	for _, k := range children[name] {
		childrenWeights = append(childrenWeights, sumWeightOfNode(k, children, weights))
	}
	problemIndex := findOutlierIndex(childrenWeights)

	return children[name][problemIndex]
}

func findProblemInNode(name string, shouldEqual int,  children map[string][]string, weights map[string]int) int {
	var childrenWeights []int
	childrenWeight := 0
	for _, k := range children[name] {
		childrenWeights = append(childrenWeights, sumWeightOfNode(k, children, weights))
		childrenWeight += sumWeightOfNode(k, children, weights)
	}
	problemIndex := findOutlierIndex(childrenWeights)
	if problemIndex == -1 {
		return shouldEqual - childrenWeight
	} else {
		return findProblemInNode(children[name][problemIndex], sumWeightOfNode(children[name][findNonOutlierIndex(childrenWeights)], children, weights), children, weights)
	}
}



func findOutlierIndex(mySlice []int) int {
	valueSeen := make(map[int]int)
	index := -1
	outlierValue := 0

	for _, value := range mySlice{
		valueSeen[value]++
	}

	for k := range valueSeen {
		if valueSeen[k] == 1 {
			outlierValue = k
		}
	}

	for i, value := range mySlice {
		if value == outlierValue {
			index = i
		}
	}

	return index
}

func findNonOutlierIndex(mySlice []int) int {
	valueSeen := make(map[int]int)
	index := -1
	outlierValue := 0

	for _, value := range mySlice{
		valueSeen[value]++
	}

	for k := range valueSeen {
		if valueSeen[k] == 1 {
			outlierValue = k
		}
	}

	for i, value := range mySlice {
		if value != outlierValue {
			index = i
		}
	}

	return index
}

func reduce(lookup map[string][]string, name string) {
	for _, child := range lookup[name] {
		//fmt.Println("\t'" + child + "'")
		reduce(lookup, child)
		delete(lookup, child)
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
