package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"flag"
	"fmt"
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
	input_file, err := ioutil.ReadFile(filename)
	check(err)
	// Your Code goes below!

	string_values := strings.Split(strings.Replace(string(input_file), "\n", "", -1), ",")
	var values []int
	for _, value := range string_values {
		intVal, err := strconv.Atoi(value)
		check(err)

		values = append(values, intVal)
	}

	result := process(12, 2, values)
	fmt.Printf("Result1: %d\n", result)

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			result2 := process(i, j, values)
			//fmt.Printf("%d,%d: %d\n", i, j, result2)
			if result2 == 19690720 {
				fmt.Printf("FOUND IT!\n")
				fmt.Printf("Result 2: %d, %d === %d\n", i, j, 100 * i + j)
				break
			}
		}
	}
}

func process(replace1 int, replace2 int, valuesOriginal []int) int {
	position := 0
	stop := false

	var values []int
	for _, val := range valuesOriginal {
		values = append(values, val)
	}

	values[1] = replace1
	values[2] = replace2

	//fmt.Printf("%d, %d\n", replace1, replace2)

	for !stop {
		if values[position] == 1 {
			values, stop = add(position, values)
		} else if values[position] == 2 {
			values, stop = multiply(position, values)
		} else if values[position] == 99 {
			stop = true
		}
		position += 4
	}
	return values[0]
}

func add(start_index int, values []int) ([]int, bool) {
	if start_index < len(values)-3 {

		addParamIndex1 := values[start_index+1]
		addParamIndex2 := values[start_index+2]
		storageIndex := values[start_index+3]

		if addParamIndex1 >= len(values) {
			return values, true
		}

		if addParamIndex2 >= len(values) {
			return values, true
		}

		addParam1 := values[addParamIndex1]
		addParam2 := values[addParamIndex2]

		if storageIndex < len(values) {
			values[storageIndex] = addParam1 + addParam2
		}

	} else {
		return values, true
	}
	return values, false
}

func multiply(start_index int, values []int) ([]int, bool) {
	if start_index < len(values) - 3 {
		multParamIndex1 := values[start_index + 1]
		multParamIndex2 := values[start_index + 2]
		storageIndex := values[start_index + 3]

		multParam1 := 0
		multParam2 := 0

		if multParamIndex1 < len(values) {
			multParam1 = values[multParamIndex1]
		} else {
			return values, true
		}

		if multParamIndex2 < len(values) {
			multParam2 = values[multParamIndex2]
		} else {
			return values, true
		}

		if storageIndex < len(values) {
			values[storageIndex] = multParam1 * multParam2
		} else {
			return values, true
		}
	} else {
		return values, true
	}

	return values, false
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
