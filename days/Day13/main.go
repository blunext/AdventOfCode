package Day13

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strconv"
	"strings"
)

func Goooo() {
	fmt.Println("--------- DAY 13 ---------")
	//lines := tools.ReadFile(("days/Day13/testInput.txt"))
	lines := tools.ReadFile(("days/Day13/input.txt"))
	arrival, _ := strconv.Atoi(lines[0])

	list := cleanInputList(lines)

	min := 999999999999999999
	busFound := 0
	departure := 0
	for _, no := range list {
		nextArrival := (arrival/no + 1) * no
		diff := nextArrival - arrival
		if diff < min {
			min = diff
			busFound = no
			departure = nextArrival
		}
		fmt.Printf("bud: %d, nextArrival: %d, wait: %d, min: %d, foundBus: %d, departure: %d\n", no, nextArrival, diff, min, busFound, departure)
	}

	fmt.Printf("%d", (departure-arrival)*busFound)
}

func cleanInputList(lines []string) []int {
	var list []int
	inputLines := strings.Split(lines[1], ",")
	for _, v := range inputLines {
		if v == "x" {
			continue
		}
		i, _ := strconv.Atoi(v)
		list = append(list, i)
	}
	return list
}
