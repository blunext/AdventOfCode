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

	wait := part1(lines)
	fmt.Printf("part 1: %d\n", wait)

	val := part2(lines)
	fmt.Printf("part 2: %d\n", val)
	return
}

func part1(lines []string) int {
	arrival, _ := strconv.Atoi(lines[0])
	list := cleanInputLeaveOnlyNumbers(lines)

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
		//fmt.Printf("bud: %d, nextArrival: %d, wait: %d, min: %d, foundBus: %d, departure: %d\n", no, nextArrival, diff, min, busFound, departure)
	}
	wait := (departure - arrival) * busFound
	return wait
}

func cleanInputLeaveOnlyNumbers(lines []string) []int {
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

func part2(lines []string) int {

	buses := cleanInputForPart2(lines)

	timestamp := 1
	for {
		skip := 1
		ok := true
		for offset := 0; offset < len(buses); offset++ {
			if (timestamp+offset)%buses[offset] != 0 {
				ok = false
				break
			}
			skip *= buses[offset]
		}
		if ok {
			return timestamp
		}
		timestamp += skip
	}
}

func cleanInputForPart2(lines []string) []int {
	var buses []int
	for _, id := range strings.Split(lines[1], ",") {
		if id == "x" {
			id = "1"
		}
		val, _ := strconv.Atoi(id)
		buses = append(buses, val)
	}
	return buses
}
