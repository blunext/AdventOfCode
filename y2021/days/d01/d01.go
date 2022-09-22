package d01

import (
	"fmt"

	"AdventOfCode/tools"
)

func countDepths(depths []int) int {
	n, last := 0, 0
	if len(depths) > 0 {
		last = depths[0]
	}
	for _, depth := range depths {
		if depth > last {
			n++
		}
		last = depth
	}
	return n
}

func transformMeasurements(measurements []int) []int {
	var transformed []int
	for i := 0; i < len(measurements); i++ {
		if i+3 > len(measurements) {
			break
		}
		sum := 0
		for t := 0; t < 3; t++ {
			sum += measurements[i+t]
		}
		transformed = append(transformed, sum)
	}
	return transformed
}

func Main() {
	depths := tools.GetSliceOfInts("y2021/days/d01/data.txt")
	result := countDepths(depths)
	fmt.Println("DAY #1 A")
	fmt.Printf("result: %d\n", result)

	fmt.Println("DAY #1 B")
	result = countDepths(transformMeasurements(depths))
	fmt.Printf("result: %d\n", result)
}
