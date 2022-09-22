package Day01

import (
	"fmt"

	"AdventOfCode/tools"
)

func Goooo() {
	fmt.Println("--------- DAY 02 ---------")

	lines := tools.ReadFile("y2020/days/Day01/input.txt")
	ints := tools.ConvertStrArrayIntoInts(lines)

	for _, x := range ints {
		for _, y := range ints {
			if x+y == 2020 {
				fmt.Printf("%d + %d = %d -> %d\n", x, y, x+y, x*y)
			}
		}
	}
}
