package main

import (
	"bufio"
	"os"
	"flag"
	"fmt"
)

func main() {
	boolPtr := flag.Bool("test", false, "test mode")
	flag.Parse()

	var filename string
	var values []int
	valueSize := 256
	if *boolPtr {
		filename = "input-test.txt"
		//valueSize = 5
	} else {
		filename = "input.txt"
	}
	input_file, err := readLines(filename)
	check(err)
	//input := strings.Split(input_file[0], ",")

	for i:=0; i < valueSize; i++ {
		values = append(values, i)
	}

	// Your Code goes below!
	fmt.Println(values)
	fmt.Println(input_file)

	currentIndex := 0
	jumpSize := 0


	//for _, num := range input {
	//	num, err := strconv.Atoi(num)
	//	check(err)
	//
	//	fmt.Println("======")
	//	fmt.Printf("Num: %d\tCurrentIndex: %d\t%v\n", num, currentIndex, values)
	//
	//	//Select the 'num' amount of numbers from current index
	//	var reverseArr []int
	//	for i := 0; i < num; i++ {
	//		index := currentIndex + i
	//		if index >= len(values) {
	//			index -= len(values)
	//		}
	//		reverseArr = append(reverseArr, values[index])
	//	}
	//	fmt.Printf("\t%v\n", reverseArr)
	//	//reverse them in the values array
	//	for i := 0; i < num; i++ {
	//		index := currentIndex + i
	//		if index >= len(values) {
	//			index -= len(values)
	//		}
	//		values[index] = reverseArr[num-i-1]
	//		fmt.Printf("\t%d\t%v\n", index, values)
	//
	//	}
	//
	//	//increment currentIndex 'num' + jumpSize
	//	currentIndex = (currentIndex + num + jumpSize)
	//	if currentIndex >= len(values) {
	//		currentIndex -= len(values)
	//	}
	//
	//	//increment jumpSize
	//	jumpSize++
	//}
	//
	//fmt.Println(values)
	//fmt.Printf("Part 1 solution: %d\n", values[0] * values[1])


	// PART 2 BELOW
	unicodeCodePoints := []byte(input_file[0])
	unicodeCodePoints = append(unicodeCodePoints, 17, 31, 73, 47, 23)
	fmt.Printf("%v\n", unicodeCodePoints)

	currentIndex = 0
	jumpSize = 0
	for i:=0; i<64; i++ {
		for _, num := range unicodeCodePoints {
			//num, err := strconv.Atoi(num)
			//check(err)

			fmt.Println("======")
			fmt.Printf("Num: %d\tCurrentIndex: %d\n", num, currentIndex)

			//Select the 'num' amount of numbers from current index
			var reverseArr []int
			for i := 0; i < int(num); i++ {
				index := currentIndex + i
				for index >= len(values) {
					index -= len(values)
				}
				reverseArr = append(reverseArr, values[index])
			}
			//fmt.Printf("\t%v\n", reverseArr)
			//reverse them in the values array
			for j := 0; j< int(num); j++ {
				index := currentIndex + j
				for index >= len(values) {
					index -= len(values)
				}
				values[index] = reverseArr[int(num)-j-1]
				//fmt.Printf("\t%d\t%v\n", index, values)

			}

			//increment currentIndex 'num' + jumpSize
			currentIndex = (currentIndex + int(num) + jumpSize)
			if currentIndex >= len(values) {
				currentIndex -= len(values)
			}

			//increment jumpSize
			jumpSize++
		}
	}
	var finalString string
	for i:=0; i<16; i++ {
		character := values[i*16]
		for j:=1; j<16; j++ {
			character = character ^ values[(i*16)+j]
		}
		hexChars := fmt.Sprintf("%x", character)
		if len(hexChars) < 2 {
			hexChars = "0" + hexChars
		}
		finalString = fmt.Sprintf("%s%s", finalString, hexChars)
		fmt.Printf("Char %d: int: %d\thex: %s\n", i, character, hexChars)
	}
	str := fmt.Sprintf("%s", finalString)
	fmt.Printf("Part 2 solution: %s\n", str)

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
