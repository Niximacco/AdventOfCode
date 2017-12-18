package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
	"strings"
	"strconv"
)

var registers map[string]int
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
	registers = make(map[string]int)
	registers["lastsnd"] = 0
	done := false

	for i:=0; i<len(input_file) && !done; {
		instruction := input_file[i]
		instructionParts := strings.Fields(instruction)
		arg1 := instructionParts[1]
		arg2 := ""
		if len(instructionParts) > 2 {
			arg2 = instructionParts[2]
		}

		fmt.Printf("%s\n", instruction)
		switch instructionParts[0] {
		case "set":
			set(arg1, arg2)
			i++
		case "add":
			add(arg1, arg2)
			i++
		case "mul":
			mul(arg1, arg2)
			i++
		case "mod":
			mod(arg1, arg2)
			i++
		case "snd":
			snd(arg1)
			i++
		case "rcv":
			result := rcv(arg1)
			if result >= 0 {
				fmt.Printf("Part 1 Answer: %d\n", result)
				done = true
			}
			i++
		case "jgz":
			i += jgz(arg1, arg2)
		}
	}
}

func jgz(arg1 string, arg2 string) int {
	if registers[arg1] > 0 {
		return getArg(arg2)
	} else {
		return 1
	}
}

func rcv(arg1 string) int {
	if registers[arg1] != 0 {
		return registers["lastsnd"]
	} else {
		return -1
	}
}

func mod(arg1 string, arg2 string) {
	registers[arg1] = registers[arg1] % getArg(arg2)
}

func mul(arg1 string, arg2 string) {
	registers[arg1] = registers[arg1] * getArg(arg2)
}

func add(arg1 string, arg2 string) {
	registers[arg1] += getArg(arg2)
}

func snd(arg1 string) {
	registers["lastsnd"]  = getArg(arg1)
}

func set(arg1 string, arg2 string) {
	registers[arg1] = getArg(arg2)
}

func getArg(arg string) int {
	num := 0
	err := error(nil)
	if _, ok := registers[arg]; ok {
		num = registers[arg]
		//do something here
	} else {
		num, err = strconv.Atoi(arg)
		check(err)
	}

	return num
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
