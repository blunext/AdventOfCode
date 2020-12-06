package Day05

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strconv"
	"strings"
)

func bitReplacer(s string) string {
	s = strings.ReplaceAll(s, "F", "0")
	s = strings.ReplaceAll(s, "B", "1")
	s = strings.ReplaceAll(s, "L", "0")
	s = strings.ReplaceAll(s, "R", "1")
	return s
}

func getNumber(s string) int {
	bits := bitReplacer(s)
	number, _ := strconv.ParseInt(bits, 2, 16)
	return int(number)
}

func getSeatId(s string) int {
	row := getNumber(s[:7])
	col := getNumber(s[7:])
	return row*8 + col
}

func Goooo() {
	fmt.Println("--------- DAY 05 ---------")
	lines := tools.ReadFile(("days/Day05/input.txt"))

	seats := make(map[int]bool)
	max := 0
	for _, line := range lines {
		seatId := getSeatId(line)
		seats[seatId] = true
		if seatId > max {
			max = seatId
		}
	}
	fmt.Printf("max %d\n", max)

	for i := 0; i < 128*8; i++ {
		if _, ok := seats[i]; !ok {
			_, ok1 := seats[i-1]
			_, ok2 := seats[i+1]
			if ok1 && ok2 {
				fmt.Printf("seat: %d\n", i)
			}
		}

	}
}
