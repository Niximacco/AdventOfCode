package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"
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

	wire1 := strings.Split(input_file[0], ",")
	wire2 := strings.Split(input_file[1], ",")

	wire1Pos := getPositions(wire1)
	wire2Pos := getPositions(wire2)

	result := HashIntersection(wire1Pos, wire2Pos)
	fmt.Printf("%v\n", result)

	closest := findClosestDistance(result)

	fmt.Printf("Part 1: %d\n", closest)
}

func findClosestDistance(distances []interface{}) int {
	min := calculateDistance(distances[0].(string))
	for _, distance := range distances {
		dist := calculateDistance(distance.(string))
		if dist < min {
			min = dist
		}
	}

	return min
}

func calculateDistance(distance string) int {
	numbers := strings.Split(distance, ":")
	num1, err := strconv.Atoi(numbers[0])
	check(err)

	num2, err := strconv.Atoi(numbers[1])
	check(err)

	if num1 < 0 {
		num1 = -num1
	}

	if num2 < 0 {
		num2 = -num2
	}

	return num1 + num2
}

func HashIntersection(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	hash := make(map[interface{}]bool)
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		hash[el] = true
	}

	for i := 0; i < bv.Len(); i++ {
		el := bv.Index(i).Interface()
		if _, found := hash[el]; found {
			set = append(set, el)
		}
	}

	return set
}

func getPositions(commands []string) []string {
	var positions []string
	x := 0
	y := 0
	for _, command := range commands {
		direction := command[0]
		amount, err := strconv.Atoi(command[1:])
		check(err)

		fmt.Printf("moving %c by %d\n", direction, amount)

		for i := 0; i < amount; i++ {
			switch direction {
			case 'L':
				x--
				break
			case 'U':
				y++
				break
			case 'R':
				x++
				break
			case 'D':
				y--
				break
			}
			positions = append(positions, fmt.Sprintf("%d:%d", x, y))
		}
	}

	return positions
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
