package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Passport map[string]string

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

	passports := parsePassports(fileContents)
	check(err)

	resultPart1 := part1(passports)
	fmt.Printf("Part 1 Result: %d\n", resultPart1)

	resultPart2 := part2(passports)
	fmt.Printf("Part 2 Result: %d\n", resultPart2)
}

func parsePassports(data string) (passports []Passport) {
	for _, passport := range strings.Split(data, "\n\n") {
		newPassport := Passport{}
		for _, field := range strings.Fields(passport) {
			keyValue := strings.Split(field, ":")
			newPassport[keyValue[0]] = keyValue[1]
		}
		passports = append(passports, newPassport)
	}
	return
}

func part1(passports []Passport) (valid int) {
	for _, passport := range passports {
		if passport.isPassportValidPart1() {
			valid++
		}
	}
	return
}

func part2(passports []Passport) (valid int) {
	for _, passport := range passports {
		if passport.isPassportValidPart2() {
			valid++
		}
	}
	return
}

func (passport Passport) isPassportValidPart1() (valid bool) {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range requiredFields {
		if _, present := passport[field]; !present {
			return false
		}
	}
	return true
}

func (passport Passport) isPassportValidPart2() (valid bool) {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	type minMax struct {
		min int
		max int
	}

	// Values specified from the puzzle
	minMaxFields := map[string]minMax{
		"byr": {1920, 2002},
		"iyr": {2010, 2020},
		"eyr": {2020, 2030},
		"in":  {59, 76},
		"cm":  {150, 193},
	}

	for _, field := range requiredFields {
		if value, present := passport[field]; !present {
			return false
		} else {
			fmt.Printf("passport %v\n", passport)
			switch field {
			case "byr", "iyr", "eyr":
				number, _ := strconv.Atoi(value)
				if number < minMaxFields[field].min || number > minMaxFields[field].max {
					fmt.Printf("failed %s\n", field)
					return false
				}
			case "hgt":
				var reHeight = regexp.MustCompile("([0-9]+)([a-z]+)")
				matches := reHeight.FindStringSubmatch(value)
				if len(matches) < 2 {

					fmt.Printf("failed 1 %s\n", field)
					return false
				}

				number, _ := strconv.Atoi(matches[1])
				if number < minMaxFields[matches[2]].min || number > minMaxFields[matches[2]].max {
					fmt.Printf("failed 2 %s\n", field)
					return false
				}
			case "hcl":
				matches, _ := regexp.MatchString("^#\\w{6}$", value)
				if !matches {
					fmt.Printf("failed %s\n", field)
					return false
				}
			case "ecl":
				switch value {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":

				default:
					fmt.Printf("failed %s\n", field)
					return false
				}
			case "pid":
				matches, _ := regexp.MatchString("^[0-9]{9}$", value)
				if !matches {
					fmt.Printf("failed %s\n", field)
					return false
				}
			}
		}
	}

	fmt.Printf("VALID: %v\n", passport)
	return true
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
