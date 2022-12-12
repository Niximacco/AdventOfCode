package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Id                 int
	OperationX         string
	OperationY         string
	Operation          string
	TestDiv            int
	TrueThrowToMonkey  int
	FalseThrowToMonkey int
}

var itemsByMonkey map[int][]int
var business map[int]int
var superMod int

func (monkey *Monkey) RunRound(part int) {
	//fmt.Printf("Monkey %d:\n", monkey.Id)
	for _, item := range itemsByMonkey[monkey.Id] {
		business[monkey.Id]++
		//fmt.Printf("  Monkey inspects an item with worry level of %d.\n", item)
		newWorry := 0
		switch monkey.Operation {
		case "*":
			//fmt.Printf("    Worry level is multiplied by %s to ", monkey.OperationY)
			multiplier := item
			if monkey.OperationY != "old" {
				multiplier, _ = strconv.Atoi(monkey.OperationY)
			}

			newWorry = item * multiplier
			//fmt.Printf("%d\n", newWorry)
		case "+":
			//fmt.Printf("    Worry level is incrembented by %s to ", monkey.OperationY)
			incrementer := item
			if monkey.OperationY != "old" {
				incrementer, _ = strconv.Atoi(monkey.OperationY)
			}

			newWorry = item + incrementer
			//fmt.Printf("%d\n", newWorry)

		}

		if part == 1 {
			newWorry = newWorry / 3
		} else {
			newWorry = newWorry % superMod
		}
		//fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d\n", newWorry)

		if newWorry%monkey.TestDiv == 0 {
			//fmt.Printf("    Current worry level is divisible by %d.\n", monkey.TestDiv)
			//fmt.Printf("    Item with worry level %d is thrown to monkey %d\n", newWorry, monkey.TrueThrowToMonkey)
			itemsByMonkey[monkey.TrueThrowToMonkey] = append(itemsByMonkey[monkey.TrueThrowToMonkey], newWorry)
		} else {
			//fmt.Printf("    Current worry level is not divisible by %d.\n", monkey.TestDiv)
			//fmt.Printf("    Item with worry level %d is thrown to monkey %d\n", newWorry, monkey.FalseThrowToMonkey)
			itemsByMonkey[monkey.FalseThrowToMonkey] = append(itemsByMonkey[monkey.FalseThrowToMonkey], newWorry)
		}
	}

	itemsByMonkey[monkey.Id] = make([]int, 0)
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
	part1, part2 := 0, 0
	pt1Monkeys := make([]Monkey, 0)
	pt2Monkeys := make([]Monkey, 0)
	itemsByMonkey = make(map[int][]int)

	business = make(map[int]int)
	superMod = 1

	re := regexp.MustCompile(`Monkey (\d):\n. Starting items: (.+)\n. Operation: new = (.+) (.) (.+)\n. Test: divisible by (\d+)\n    If true: throw to monkey (\d)\n.   If false: throw to monkey (\d)`)
	res := re.FindAllStringSubmatch(strings.Join(input_file, "\n"), -1)
	for i := range res {
		fmt.Printf("Monkey %v\n  Starting Items: '%v'\n  Operation: new = '%v'\n  Test: '%v\n    True: -> '%v'\n    False: -> '%v'\n", res[i][1], res[i][2], res[i][3], res[i][4], res[i][5], res[i][6])
		monkeyNum, err := strconv.Atoi(res[i][1])
		check(err)

		monkey := Monkey{
			Id:                 monkeyNum,
			OperationX:         "",
			OperationY:         "",
			Operation:          "",
			TestDiv:            0,
			TrueThrowToMonkey:  0,
			FalseThrowToMonkey: 0,
		}

		itemsByMonkey[monkeyNum] = make([]int, 0)

		startingItems := res[i][2]
		startingItems = strings.ReplaceAll(startingItems, " ", "")
		items := strings.Split(startingItems, ",")
		for _, item := range items {
			itemNum, err := strconv.Atoi(item)
			check(err)
			itemsByMonkey[monkeyNum] = append(itemsByMonkey[monkeyNum], itemNum)
		}

		monkey.OperationX = res[i][3]
		monkey.Operation = res[i][4]
		monkey.OperationY = res[i][5]

		test, err := strconv.Atoi(res[i][6])
		check(err)
		monkey.TestDiv = test
		superMod = superMod * monkey.TestDiv

		trueMonkey, err := strconv.Atoi(res[i][7])
		check(err)
		monkey.TrueThrowToMonkey = trueMonkey

		falseMonkey, err := strconv.Atoi(res[i][8])
		check(err)
		monkey.FalseThrowToMonkey = falseMonkey

		pt1Monkeys = append(pt1Monkeys, monkey)
		pt2Monkeys = append(pt2Monkeys, monkey)
	}

	backup := make(map[int][]int, 0)
	CopyMap(itemsByMonkey, backup)

	for i := 1; i <= 20; i++ {
		fmt.Printf("Round %d\n", i)
		for _, monkey := range pt1Monkeys {
			monkey.RunRound(1)
		}

		PrintMonkeyContents(itemsByMonkey)
		part1 = Top2()
		PrintMonkeyBusiness()
	}

	business = make(map[int]int, 0)

	CopyMap(backup, itemsByMonkey)

	println("===============")

	for i := 1; i <= 10000; i++ {
		for _, monkey := range pt2Monkeys {
			monkey.RunRound(2)
		}

		part2 = Top2()
		PrintMonkeyContents(itemsByMonkey)
		PrintMonkeyBusiness()
	}

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}

func PrintMonkeyContents(items map[int][]int) {
	for i := 0; i < len(items); i++ {
		fmt.Printf("Monkey %d: %v\n", i, items[i])
	}
}

func PrintMonkeyBusiness() {
	for i := 0; i < len(business); i++ {
		fmt.Printf("Monkey %d inspected items %d times\n", i, business[i])
	}
}

func CopyMap(original map[int][]int, new map[int][]int) {
	for k, v := range original {
		new[k] = v
	}
}

func Top2() (solution int) {
	results := make([]int, 0)
	for i := 0; i < len(business); i++ {
		results = append(results, business[i])
	}

	sort.Ints(results)
	solution = results[len(results)-1] * results[len(results)-2]
	fmt.Printf("Top 2: %d * %d = %d\n", results[len(results)-1], results[len(results)-2], solution)
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
