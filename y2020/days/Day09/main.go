package Day09

import (
	"fmt"

	"AdventOfCode/tools"
)

func Goooo() {
	fmt.Println("--------- DAY 09 ---------")

	// preamble := 5
	// lines := tools.ReadLines(("y2020/days/Day09/testInput.txt"))

	preamble := 25
	lines := tools.ReadLines(("y2020/days/Day09/input.txt"))

	intLines := tools.ConvertStrArrayIntoInts(lines)

	firstNumber := findFirstThatNotSUm(preamble, intLines)
	fmt.Printf("%d\n", firstNumber)

	min, max := findRangeThatSum(firstNumber, intLines)
	fmt.Printf("%d, %d = %d\n", min, max, min+max)

}

func findRangeThatSum(number int, lines []int) (int, int) {
	for i := 0; i < len(lines); i++ {
		sum := 0
		for x := i; x < len(lines); x++ {
			sum += lines[x]
			if sum > number {
				break
			}
			if sum == number {
				return findMinAndMax(lines[i:x])
			}

		}
	}
	return 0, 0
}

func findMinAndMax(lines []int) (int, int) {
	min, max := lines[0], lines[0]
	for i := 1; i < len(lines); i++ {
		if min < lines[i] {
			min = lines[i]
		}
		if max > lines[i] {
			max = lines[i]
		}
	}
	return min, max
}

func findFirstThatNotSUm(preamble int, intLines []int) int {
	for index := preamble; index < len(intLines); index++ {
		if !checkOneInRange(intLines[index], intLines[index-preamble:index]) {
			return intLines[index]
		}
	}
	return 0
}

func checkOneInRange(number int, slice []int) bool {
	for x, val1 := range slice {
		for y, val2 := range slice {
			if val1 == val2 {
				continue
			}
			if slice[x]+slice[y] == number {
				return true
			}
		}
	}
	return false
}
