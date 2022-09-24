package d06

import (
	"fmt"
	"time"

	"AdventOfCode/tools"
)

func process(fishes []int, days int) int {
	groups := initGroups(fishes)

	var newGroups []int
	newGroups = append(newGroups, groups...)
	for d := 0; d < days; d++ {
		for i := 8; i >= 0; i-- {
			switch i {
			case 0:
				newGroups[6] += groups[0]
				newGroups[8] = groups[0]
			default:
				newGroups[i-1] = groups[i]
			}
		}
		groups = []int{}
		groups = append(groups, newGroups...)
	}
	// fmt.Print(strings.Trim(fmt.Sprint(groups), "[]"))
	// fmt.Println()
	return count(groups)
}

func count(groups []int) int {
	var c int
	for _, g := range groups {
		c += g
	}
	return c
}

func initGroups(fishes []int) []int {
	groups := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, f := range fishes {
		groups[f]++
	}
	return groups
}

func Main() {
	timeStart := time.Now()
	defer func() {
		fmt.Printf("Execution time: %v\n", time.Since(timeStart))
	}()
	fmt.Println("DAY #6 A")
	data := tools.ReadFile("y2021/days/d06/data.txt")
	fishes := tools.ConvertCommaSeparatedStrIntoInts(data[0])
	fmt.Printf("result: %d\n", process(fishes, 80))

	fmt.Println("DAY #6 B")
	fishes = tools.ConvertCommaSeparatedStrIntoInts(data[0])
	fmt.Printf("result: %d\n", process(fishes, 256))

}
