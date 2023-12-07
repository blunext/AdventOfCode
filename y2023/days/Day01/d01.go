package Day01

import (
	"AdventOfCode/tools"
	"errors"
	"fmt"
	"strconv"
)

var digits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func search4Digits(line string) []int {
	var found []int
	for i := range line {
		for digit, value := range digits {
			dl := i + len(digit)
			if dl > len(line) {
				continue
			}
			if line[i:dl] == digit {
				found = append(found, value)
			}
		}
		//fmt.Println(i)
	}
	return found
}

func fromRight(line string) (int, error) {
	return fromLeft(tools.ReverseString(line))
}

func fromLeft(line string) (int, error) {
	for _, char := range line {
		if i, err := strconv.Atoi(string(char)); err == nil {
			return i, nil
		}
	}
	fmt.Println()
	return -1, errors.New("can't find a number")
}
func processA() int {
	sum := 0
	lines := tools.ReadLines("y2023/days/Day01/data.txt")
	for _, line := range lines {
		left, err := fromLeft(line)
		if err != nil {
			panic("no left number found in: " + line)
		}
		right, err := fromRight(line)
		if err != nil {
			panic("no left number found in: " + line)
		}
		//fmt.Printf("%d, %d\n", left, right)
		number := left*10 + right
		sum += number
	}
	return sum
}

func processB() int {
	sum := 0
	lines := tools.ReadLines("y2023/days/Day01/data.txt")
	for _, line := range lines {
		found := search4Digits(line)
		//fmt.Printf("%v\n", found)
		number := found[0]*10 + found[len(found)-1]
		sum += number
	}
	return sum
}
func Main() {
	fmt.Println("DAY #1 A")
	fmt.Printf("result: %v\n", processA())
	fmt.Println("DAY #1 B")
	fmt.Printf("result: %v\n", processB())
}
