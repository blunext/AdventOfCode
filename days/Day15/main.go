package Day15

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strings"
)

type number struct {
	pos, prevPos, count int
}

func Goooo() {
	fmt.Println("--------- DAY 15 ---------")
	input := "2,1,10,11,0,6"
	//input := "0,3,6"

	lastSpoken := findNumber(input, 2020)
	fmt.Printf("Part 1: %d\n", lastSpoken)

	//lastSpoken = findNumber(input, 30000000)
	fmt.Printf("Part 2: %d <- want to find real finding, uncomment findNumber function\n", lastSpoken)

}

func findNumber(input string, matchOccurange int) int {
	memory := make(map[int]*number)
	lastSpoken := 0
	i := 1
	for _, lastSpoken = range tools.ConvertStrArrayIntoInts(strings.Split(input, ",")) {
		memory[lastSpoken] = &number{pos: i, count: 1}
		i++
	}

	for ; i <= matchOccurange; i++ {
		//fmt.Printf("%d: %d\n", i-1, lastSpoken)
		lastNumber := memory[lastSpoken]
		if lastNumber.count == 1 {
			if zero, ok := memory[0]; ok {
				zero.prevPos = zero.pos
				zero.pos = i
				zero.count++
			} else {
				memory[0] = &number{pos: i, count: 1}
			}
			lastSpoken = 0
			continue
		}

		newNumber := lastNumber.pos - lastNumber.prevPos
		if num, ok := memory[newNumber]; ok {
			num.prevPos = num.pos
			num.pos = i
			num.count++
			lastSpoken = newNumber
			continue
		}
		memory[newNumber] = &number{pos: i, count: 1}
		lastSpoken = newNumber
	}
	return lastSpoken
}
