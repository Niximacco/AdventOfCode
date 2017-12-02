package main

import (
    "os"
    "bufio"
    "strings"
	"strconv"
	"fmt"
)

func main() {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	inFile, _ := os.Open("input.txt")
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines)

    checksum1 := 0
    checksum2 := 0

    for scanner.Scan() {
		min := MaxInt
		max := MinInt
        words := strings.Fields(scanner.Text())
        for _, word1 := range words {
        	num1, err := strconv.Atoi(word1)
        	if err != nil {
        		panic(err)
			}

			if num1 < min {
				min = num1
			}

			if num1 > max {
				max = num1
			}
			for _, word2 := range words {
				num2, err := strconv.Atoi(word2)
				if err != nil {
					panic(err)
				}

				if num1 != num2 {
					if num1 % num2 == 0 {
						checksum2 += (num1 / num2)
					}
				}
			}
		}
		checksum1 += (max-min)
    }
    fmt.Printf("Part 1 Answer: %d", checksum1)
	fmt.Printf("Part 2 Answer: %d", checksum2)
}