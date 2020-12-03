package Day2

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strconv"
	"strings"
)

func check1(test string) bool {
	// "2-9 c: ccccccccc"

	tokens := strings.Split(test, " ")

	limits := strings.Split(tokens[0], "-")
	min, _ := strconv.Atoi(limits[0])
	max, _ := strconv.Atoi(limits[1])

	letter := strings.Trim(tokens[1], ":")

	occurance := strings.Count(tokens[2], letter)

	if occurance < min || occurance > max {
		return false
	}
	return true
}

func check2(test string) bool {
	// "2-9 c: ccccccccc"

	tokens := strings.Split(test, " ")

	positions := strings.Split(tokens[0], "-")
	pos1, _ := strconv.Atoi(positions[0])
	pos2, _ := strconv.Atoi(positions[1])

	letter := []byte(tokens[1])[0]

	ok1 := []byte(tokens[2])[pos1-1] == letter
	ok2 := []byte(tokens[2])[pos2-1] == letter

	if (ok1 && !ok2) || (!ok1 && ok2) {
		return true
	}
	return false
}

func Goooo() {
	fmt.Println("--------- DAY 03 ---------")
	lines := tools.ReadFile(("days/day03/input.txt"))

	i := 0
	for _, line := range lines {
		if check1(line) {
			i++
		}
	}
	fmt.Printf("password ok: %d\n", i)

	i = 0
	for _, line := range lines {
		if check2(line) {
			i++
		}
	}
	fmt.Printf("password ok: %d\n", i)
}
