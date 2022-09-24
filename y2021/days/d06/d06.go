package d06

import (
	"fmt"
	"time"

	"AdventOfCode/tools"
)

func process(fishes []int8, days int) int {
	// fmt.Print(strings.Trim(fmt.Sprint(fishes), "[]"))
	// fmt.Println()

	var fishesOfNewDay []int8

	for d := 0; d < days; d++ {
		fishesOfNewDay = []int8{}
		fishesOfNewDay = append(fishesOfNewDay, fishes...)

		for i := range fishes {
			fishesOfNewDay[i]--
			if fishesOfNewDay[i] == -1 {
				fishesOfNewDay[i] = 6
				fishesOfNewDay = append(fishesOfNewDay, 8)
			}
		}
		fishes = []int8{}
		fishes = append(fishes, fishesOfNewDay...)

		// fmt.Print(strings.Trim(fmt.Sprint(fishes), "[]"))
		// fmt.Println()
	}
	return len(fishesOfNewDay)
}

func Main() {
	timeStart := time.Now()
	defer func() {
		fmt.Printf("Execution time: %v\n", time.Since(timeStart))
	}()
	fmt.Println("DAY #6 A")
	data := tools.ReadFile("y2021/days/d06/data.txt")
	fishes := tools.ConvertCommaSeparatedStrIntoInts8(data[0])
	fmt.Printf("result: %d\n", process(fishes, 80))

	fmt.Println("DAY #6 B")
	fishes = tools.ConvertCommaSeparatedStrIntoInts8(data[0])
	fmt.Printf("result: %d\n", process(fishes, 256))

}
