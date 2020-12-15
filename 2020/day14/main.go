package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	MemAddress   int
	Value        int
	MaskedValue  int
	MaskedString string
	BinaryString string
	Mask         string
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
	instructions := parseInput(input_file)

	resultPart1 := part1(instructions)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(instructions)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)

}

func part1(instructions []Instruction) (sum int) {
	values := make(map[int]int)

	for _, instruction := range instructions {
		instruction.applyMask()
		values[instruction.MemAddress] = instruction.MaskedValue
	}

	for _, v := range values {
		sum += v
	}

	return
}

func part2(instructions []Instruction) (sum int) {
	values := make(map[int]int)

	for _, instruction := range instructions {
		addresses := calcPossibilities(applyMaskPart2(decimalToBinary(instruction.MemAddress), instruction.Mask))
		for _, address := range addresses {
			fmt.Printf("setting [%d] to %d\n", address, instruction.Value)
			values[address] = instruction.Value
		}
	}

	for _, v := range values {
		sum += v
	}

	return
}

func (instruction *Instruction) getPossibleAddressses() (addresses []int) {
	addressBin := applyMaskPart2(decimalToBinary(instruction.MemAddress), instruction.Mask)
	addresses = calcPossibilities(addressBin)

	return
}

func calcPossibilities(binary string) (addresses []int) {
	var possibilities []string
	possibilities = append(possibilities, "")
	for i := 0; i < len(binary); i++ {
		switch string(binary[i]) {
		case "0", "1":
			var newPossibilities []string
			for _, possibility := range possibilities {
				newPossibilities = append(newPossibilities, fmt.Sprintf("%s%s", possibility, string(binary[i])))
			}

			possibilities = newPossibilities
		case "X":
			var newPossibilities []string
			for _, possibility := range possibilities {
				newPossibilities = append(newPossibilities, fmt.Sprintf("%s%s", possibility, "0"))
				newPossibilities = append(newPossibilities, fmt.Sprintf("%s%s", possibility, "1"))
			}

			possibilities = newPossibilities
		}
	}

	for _, possibility := range possibilities {
		addresses = append(addresses, binaryToDecimal(possibility))
	}

	return
}

func (instruction *Instruction) applyMask() {
	result := ""
	for i := 0; i < len(instruction.BinaryString); i++ {
		switch string(instruction.Mask[i]) {
		case "0":
			result = fmt.Sprintf("%s0", result)
		case "1":
			result = fmt.Sprintf("%s1", result)
		case "X":
			result = fmt.Sprintf("%s%s", result, string(instruction.BinaryString[i]))
		}
	}

	instruction.MaskedString = result
	instruction.calcMaskedDecimalValue()
}

func applyMaskPart2(binary string, mask string) (result string) {
	for i := 0; i < len(binary); i++ {
		switch string(mask[i]) {
		case "0":
			result = fmt.Sprintf("%s%s", result, string(binary[i]))
		case "1":
			result = fmt.Sprintf("%s1", result)
		case "X":
			result = fmt.Sprintf("%sX", result)
		}
	}

	return
}

func (instruction *Instruction) generateBinaryString() (binary string) {
	remainder := float64(instruction.Value)
	for i := 35; i >= 0; i-- {
		if remainder >= math.Exp2(float64(i)) {
			binary = fmt.Sprintf("%s1", binary)
			remainder -= math.Exp2(float64(i))
		} else {
			binary = fmt.Sprintf("%s0", binary)
		}
	}

	instruction.BinaryString = binary

	return
}

func (instruction *Instruction) calcMaskedDecimalValue() {
	var value float64
	for i := 0; i < 36; i++ {
		switch string(instruction.MaskedString[i]) {
		case "1":
			value += math.Exp2(float64(35 - i))
		}
	}

	instruction.MaskedValue = int(value)
}

func binaryToDecimal(binary string) (decimal int) {
	var value float64
	for i := 0; i < 36; i++ {
		switch string(binary[i]) {
		case "1":
			value += math.Exp2(float64(35 - i))
		}
	}

	decimal = int(value)
	return
}

func decimalToBinary(decimal int) (binary string) {
	remainder := float64(decimal)
	for i := 35; i >= 0; i-- {
		if remainder >= math.Exp2(float64(i)) {
			binary = fmt.Sprintf("%s1", binary)
			remainder -= math.Exp2(float64(i))
		} else {
			binary = fmt.Sprintf("%s0", binary)
		}
	}

	return
}

func parseInput(input []string) (instructions []Instruction) {
	mask := ""
	for _, line := range input {
		if strings.Contains(line, "mask") {
			mask = regexp.MustCompile(`mask = (.+)$`).FindStringSubmatch(line)[1]
			continue
		}

		lineMatches := regexp.MustCompile(`mem\[(\d+)] = (\d+)$`).FindStringSubmatch(line)

		memAddress, _ := strconv.Atoi(lineMatches[1])
		value, _ := strconv.Atoi(lineMatches[2])

		newInstruction := Instruction{
			MemAddress: memAddress,
			Value:      value,
			Mask:       mask,
		}

		newInstruction.generateBinaryString()

		instructions = append(instructions, newInstruction)
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
